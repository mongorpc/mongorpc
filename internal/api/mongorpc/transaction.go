// Package mongorpc provides transaction RPC implementations.
package mongorpc

import (
	"context"
	"log/slog"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mongorpcv1 "github.com/mongorpc/mongorpc/gen/mongorpc/v1"
)

// TransactionManager manages active transactions.
type TransactionManager struct {
	mu       sync.RWMutex
	sessions map[string]*mongo.Session
}

// NewTransactionManager creates a new transaction manager.
func NewTransactionManager() *TransactionManager {
	return &TransactionManager{
		sessions: make(map[string]*mongo.Session),
	}
}

// Global transaction manager (in production, this would be injected)
var txnManager = NewTransactionManager()

// BeginTransaction starts a new transaction.
func (s *Server) BeginTransaction(ctx context.Context, req *mongorpcv1.BeginTransactionRequest) (*mongorpcv1.BeginTransactionResponse, error) {
	slog.Debug("BeginTransaction")

	// Start a new session
	session, err := s.db.Client().StartSession()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to start session: %v", err)
	}

	// Build session options
	txnOpts := options.Transaction()
	if req.Options != nil {
		// Apply read concern
		if req.Options.GetReadOnly() != nil {
			// Read-only transaction
		}
		// Apply write concern
		if req.Options.GetReadWrite() != nil {
			// Read-write transaction
		}
	}

	// Start transaction
	if err := session.StartTransaction(txnOpts); err != nil {
		session.EndSession(ctx)
		return nil, status.Errorf(codes.Internal, "failed to start transaction: %v", err)
	}

	// Generate transaction ID (using session ID as reference)
	txnID := session.ID().String()

	// Store session
	txnManager.mu.Lock()
	txnManager.sessions[txnID] = session
	txnManager.mu.Unlock()

	slog.Info("Transaction started", "transaction_id", txnID)

	return &mongorpcv1.BeginTransactionResponse{
		Transaction: []byte(txnID),
	}, nil
}

// CommitTransaction commits an active transaction.
func (s *Server) CommitTransaction(ctx context.Context, req *mongorpcv1.CommitTransactionRequest) (*mongorpcv1.CommitTransactionResponse, error) {
	txnID := string(req.Transaction)
	slog.Debug("CommitTransaction", "transaction_id", txnID)

	// Get session
	txnManager.mu.RLock()
	session, exists := txnManager.sessions[txnID]
	txnManager.mu.RUnlock()

	if !exists {
		return nil, status.Error(codes.NotFound, "transaction not found")
	}

	// Commit transaction
	if err := session.CommitTransaction(ctx); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to commit transaction: %v", err)
	}

	// Cleanup
	session.EndSession(ctx)
	txnManager.mu.Lock()
	delete(txnManager.sessions, txnID)
	txnManager.mu.Unlock()

	slog.Info("Transaction committed", "transaction_id", txnID)

	return &mongorpcv1.CommitTransactionResponse{}, nil
}

// AbortTransaction aborts an active transaction.
func (s *Server) AbortTransaction(ctx context.Context, req *mongorpcv1.AbortTransactionRequest) (*mongorpcv1.AbortTransactionResponse, error) {
	txnID := string(req.Transaction)
	slog.Debug("AbortTransaction", "transaction_id", txnID)

	// Get session
	txnManager.mu.RLock()
	session, exists := txnManager.sessions[txnID]
	txnManager.mu.RUnlock()

	if !exists {
		return nil, status.Error(codes.NotFound, "transaction not found")
	}

	// Abort transaction
	if err := session.AbortTransaction(ctx); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to abort transaction: %v", err)
	}

	// Cleanup
	session.EndSession(ctx)
	txnManager.mu.Lock()
	delete(txnManager.sessions, txnID)
	txnManager.mu.Unlock()

	slog.Info("Transaction aborted", "transaction_id", txnID)

	return &mongorpcv1.AbortTransactionResponse{}, nil
}

// GetSession returns the session for a transaction (for use in other RPC methods).
func (s *Server) GetSession(txnID string) (*mongo.Session, bool) {
	txnManager.mu.RLock()
	defer txnManager.mu.RUnlock()
	session, exists := txnManager.sessions[txnID]
	return session, exists
}
