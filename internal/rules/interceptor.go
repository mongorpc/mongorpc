// Package rules provides middleware for rules enforcement.
package rules

import (
	"context"
	"log/slog"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// RulesInterceptor creates a gRPC interceptor that enforces security rules.
func RulesInterceptor(engine *Engine) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Extract operation info from method name
		operation := extractOperation(info.FullMethod)
		if operation == "" {
			// Unknown operation, allow by default
			return handler(ctx, req)
		}

		// Extract resource info from request
		database, collection, docID := extractResourceInfo(req)
		if database == "" {
			// No resource info, skip rules check
			return handler(ctx, req)
		}

		// Extract auth info from context
		userID, userEmail, isAdmin := extractAuthInfo(ctx)

		// Build rules request
		rulesReq := &Request{
			UserID:     userID,
			UserEmail:  userEmail,
			IsAdmin:    isAdmin,
			Database:   database,
			Collection: collection,
			DocumentID: docID,
			Operation:  operation,
		}

		// Evaluate rules
		allowed, err := engine.Evaluate(ctx, rulesReq)
		if err != nil {
			slog.Error("Rules evaluation error", "error", err)
			return nil, status.Error(codes.Internal, "rules evaluation failed")
		}

		if !allowed {
			slog.Warn("Access denied by rules",
				"method", info.FullMethod,
				"database", database,
				"collection", collection,
				"user", userID,
			)
			return nil, status.Error(codes.PermissionDenied, "access denied")
		}

		return handler(ctx, req)
	}
}

// StreamRulesInterceptor creates a streaming interceptor for rules enforcement.
func StreamRulesInterceptor(engine *Engine) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// For streaming RPCs, we check at stream start
		operation := extractOperation(info.FullMethod)
		if operation == "" {
			return handler(srv, ss)
		}

		// For now, allow streaming operations
		// Full implementation would wrap the stream
		return handler(srv, ss)
	}
}

// extractOperation maps gRPC method names to operations.
func extractOperation(method string) Operation {
	// Parse method name: /package.Service/Method
	parts := strings.Split(method, "/")
	if len(parts) < 3 {
		return ""
	}
	methodName := parts[2]

	switch {
	case strings.HasPrefix(methodName, "Get"), strings.HasPrefix(methodName, "List"):
		return OpRead
	case strings.HasPrefix(methodName, "Create"), strings.HasPrefix(methodName, "Insert"):
		return OpCreate
	case strings.HasPrefix(methodName, "Update"):
		return OpUpdate
	case strings.HasPrefix(methodName, "Delete"):
		return OpDelete
	case strings.HasPrefix(methodName, "Run"), strings.HasPrefix(methodName, "Aggregate"):
		return OpRead
	case strings.HasPrefix(methodName, "Watch"):
		return OpRead
	default:
		return ""
	}
}

// extractResourceInfo extracts database/collection/docID from request.
func extractResourceInfo(req interface{}) (database, collection, docID string) {
	// Use reflection or type assertion to extract fields
	// This is a simplified version using type assertion
	type hasDatabase interface {
		GetDatabase() string
	}
	type hasCollection interface {
		GetCollection() string
	}
	type hasID interface {
		GetId() interface{ GetHex() string }
	}

	if r, ok := req.(hasDatabase); ok {
		database = r.GetDatabase()
	}
	if r, ok := req.(hasCollection); ok {
		collection = r.GetCollection()
	}
	if r, ok := req.(hasID); ok {
		if id := r.GetId(); id != nil {
			docID = id.GetHex()
		}
	}

	return
}

// extractAuthInfo extracts user info from context metadata.
func extractAuthInfo(ctx context.Context) (userID, userEmail string, isAdmin bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", "", false
	}

	if ids := md.Get("x-user-id"); len(ids) > 0 {
		userID = ids[0]
	}
	if emails := md.Get("x-user-email"); len(emails) > 0 {
		userEmail = emails[0]
	}
	if admins := md.Get("x-is-admin"); len(admins) > 0 {
		isAdmin = admins[0] == "true"
	}

	return
}
