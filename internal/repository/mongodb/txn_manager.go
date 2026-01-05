// Package mongodb provides transaction management.
package mongodb

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readconcern"
	"go.mongodb.org/mongo-driver/v2/mongo/writeconcern"
)

// TransactionManager manages MongoDB transactions.
type TransactionManager struct {
	client   *mongo.Client
	mu       sync.RWMutex
	sessions map[string]*ManagedSession
	timeout  time.Duration
}

// ManagedSession wraps a MongoDB session with metadata.
type ManagedSession struct {
	Session   *mongo.Session
	StartTime time.Time
	Database  string
	ReadOnly  bool
}

// NewTransactionManager creates a new transaction manager.
func NewTransactionManager(client *mongo.Client, timeout time.Duration) *TransactionManager {
	if timeout <= 0 {
		timeout = 5 * time.Minute
	}

	tm := &TransactionManager{
		client:   client,
		sessions: make(map[string]*ManagedSession),
		timeout:  timeout,
	}

	// Start cleanup routine
	go tm.cleanup()

	return tm
}

// Begin starts a new transaction.
func (tm *TransactionManager) Begin(ctx context.Context, database string, readOnly bool) (string, error) {
	session, err := tm.client.StartSession()
	if err != nil {
		return "", err
	}

	txnOpts := options.Transaction()
	if readOnly {
		txnOpts.SetReadConcern(readconcern.Snapshot())
	} else {
		txnOpts.SetWriteConcern(writeconcern.Majority())
		txnOpts.SetReadConcern(readconcern.Snapshot())
	}

	if err := session.StartTransaction(txnOpts); err != nil {
		session.EndSession(ctx)
		return "", err
	}

	// Generate transaction ID
	txnID := bson.NewObjectID().Hex()

	tm.mu.Lock()
	tm.sessions[txnID] = &ManagedSession{
		Session:   session,
		StartTime: time.Now(),
		Database:  database,
		ReadOnly:  readOnly,
	}
	tm.mu.Unlock()

	return txnID, nil
}

// Commit commits a transaction.
func (tm *TransactionManager) Commit(ctx context.Context, txnID string) error {
	tm.mu.Lock()
	ms, exists := tm.sessions[txnID]
	if exists {
		delete(tm.sessions, txnID)
	}
	tm.mu.Unlock()

	if !exists {
		return mongo.ErrNoDocuments
	}

	defer ms.Session.EndSession(ctx)
	return ms.Session.CommitTransaction(ctx)
}

// Abort aborts a transaction.
func (tm *TransactionManager) Abort(ctx context.Context, txnID string) error {
	tm.mu.Lock()
	ms, exists := tm.sessions[txnID]
	if exists {
		delete(tm.sessions, txnID)
	}
	tm.mu.Unlock()

	if !exists {
		return mongo.ErrNoDocuments
	}

	defer ms.Session.EndSession(ctx)
	return ms.Session.AbortTransaction(ctx)
}

// GetSession returns the session for a transaction.
func (tm *TransactionManager) GetSession(txnID string) (*mongo.Session, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	ms, exists := tm.sessions[txnID]
	if !exists {
		return nil, false
	}
	return ms.Session, true
}

// WithTransaction executes a function within a transaction.
func (tm *TransactionManager) WithTransaction(ctx context.Context, database string, fn func(ctx context.Context) error) error {
	session, err := tm.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	txnOpts := options.Transaction().
		SetWriteConcern(writeconcern.Majority()).
		SetReadConcern(readconcern.Snapshot())

	_, err = session.WithTransaction(ctx, func(ctx context.Context) (interface{}, error) {
		return nil, fn(ctx)
	}, txnOpts)

	return err
}

// ActiveTransactions returns the number of active transactions.
func (tm *TransactionManager) ActiveTransactions() int {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return len(tm.sessions)
}

func (tm *TransactionManager) cleanup() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		tm.mu.Lock()
		now := time.Now()
		for id, ms := range tm.sessions {
			if now.Sub(ms.StartTime) > tm.timeout {
				// Abort timed-out transactions
				_ = ms.Session.AbortTransaction(context.Background())
				ms.Session.EndSession(context.Background())
				delete(tm.sessions, id)
			}
		}
		tm.mu.Unlock()
	}
}
