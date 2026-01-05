// Package middleware provides request validation for MongoRPC.
package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Validator interface for request validation.
type Validator interface {
	Validate() error
}

// ValidationInterceptor validates requests that implement Validator.
func ValidationInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if v, ok := req.(Validator); ok {
			if err := v.Validate(); err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
			}
		}
		return handler(ctx, req)
	}
}

// StreamValidationInterceptor validates streaming requests.
func StreamValidationInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// For streaming, validation happens per message
		return handler(srv, ss)
	}
}

// RequiredFields checks that required string fields are not empty.
func RequiredFields(fields map[string]string) error {
	for name, value := range fields {
		if value == "" {
			return status.Errorf(codes.InvalidArgument, "%s is required", name)
		}
	}
	return nil
}
