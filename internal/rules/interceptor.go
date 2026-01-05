// Package rules provides middleware for rules enforcement.
package rules

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	mongorpcv1 "github.com/mongorpc/mongorpc/gen/mongorpc/v1"
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
			UserID:       userID,
			UserEmail:    userEmail,
			IsAdmin:      isAdmin,
			Database:     database,
			Collection:   collection,
			DocumentID:   docID,
			Operation:    operation,
			IncomingData: extractIncomingData(req),
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

// extractIncomingData extracts data from the request.
// extractIncomingData extracts data from the request.
func extractIncomingData(req interface{}) map[string]interface{} {
	switch r := req.(type) {
	case *mongorpcv1.CreateDocumentRequest:
		if r.Document != nil {
			return documentToMap(r.Document)
		}
	}
	return nil
}

func documentToMap(doc *mongorpcv1.Document) map[string]interface{} {
	if doc == nil || doc.Fields == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range doc.Fields {
		result[k], _ = valueToInterface(v)
	}
	return result
}

func valueToInterface(val *mongorpcv1.Value) (interface{}, error) {
	if val == nil {
		return nil, nil
	}

	switch v := val.ValueType.(type) {
	case *mongorpcv1.Value_NullValue:
		return nil, nil
	case *mongorpcv1.Value_BooleanValue:
		return v.BooleanValue, nil
	case *mongorpcv1.Value_Int32Value:
		return v.Int32Value, nil
	case *mongorpcv1.Value_Int64Value:
		return v.Int64Value, nil
	case *mongorpcv1.Value_DoubleValue:
		return v.DoubleValue, nil
	case *mongorpcv1.Value_StringValue:
		return v.StringValue, nil
	case *mongorpcv1.Value_BytesValue:
		return v.BytesValue, nil
	case *mongorpcv1.Value_ObjectIdValue:
		if v.ObjectIdValue == nil {
			return nil, nil
		}
		return v.ObjectIdValue.Hex, nil
	case *mongorpcv1.Value_MapValue:
		if v.MapValue == nil {
			return nil, nil
		}
		result := make(map[string]interface{})
		for k, fieldVal := range v.MapValue.Fields {
			val, err := valueToInterface(fieldVal)
			if err != nil {
				return nil, err
			}
			result[k] = val
		}
		return result, nil
	case *mongorpcv1.Value_ArrayValue:
		if v.ArrayValue == nil {
			return nil, nil
		}
		result := make([]interface{}, 0, len(v.ArrayValue.Values))
		for _, itemVal := range v.ArrayValue.Values {
			val, err := valueToInterface(itemVal)
			if err != nil {
				return nil, err
			}
			result = append(result, val)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported value type: %T", val.ValueType)
	}
}
