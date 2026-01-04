// Package middleware provides gRPC interceptors for MongoRPC.
package middleware

import (
	"context"
	"crypto/subtle"
	"log/slog"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthConfig holds authentication configuration.
type AuthConfig struct {
	// APIKeys is a list of valid API keys
	APIKeys []string
	// JWTSecret is the secret for JWT validation
	JWTSecret string
	// SkipMethods are methods that don't require authentication
	SkipMethods []string
}

// AuthInterceptor provides API key and JWT authentication.
func AuthInterceptor(config *AuthConfig) grpc.UnaryServerInterceptor {
	skipMethodsMap := make(map[string]bool)
	for _, method := range config.SkipMethods {
		skipMethodsMap[method] = true
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Skip authentication for allowed methods
		if skipMethodsMap[info.FullMethod] {
			return handler(ctx, req)
		}

		// Extract metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		// Try API key authentication first
		if apiKeys := md.Get("x-api-key"); len(apiKeys) > 0 {
			if validateAPIKey(apiKeys[0], config.APIKeys) {
				slog.Debug("API key authentication successful", "method", info.FullMethod)
				return handler(ctx, req)
			}
			return nil, status.Error(codes.Unauthenticated, "invalid API key")
		}

		// Try Bearer token authentication
		if authHeaders := md.Get("authorization"); len(authHeaders) > 0 {
			token := strings.TrimPrefix(authHeaders[0], "Bearer ")
			if token != authHeaders[0] { // Bearer prefix was found
				if validateToken(token, config.JWTSecret) {
					slog.Debug("JWT authentication successful", "method", info.FullMethod)
					return handler(ctx, req)
				}
				return nil, status.Error(codes.Unauthenticated, "invalid token")
			}
		}

		return nil, status.Error(codes.Unauthenticated, "authentication required")
	}
}

// StreamAuthInterceptor provides authentication for streaming RPCs.
func StreamAuthInterceptor(config *AuthConfig) grpc.StreamServerInterceptor {
	skipMethodsMap := make(map[string]bool)
	for _, method := range config.SkipMethods {
		skipMethodsMap[method] = true
	}

	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Skip authentication for allowed methods
		if skipMethodsMap[info.FullMethod] {
			return handler(srv, ss)
		}

		// Extract metadata
		md, ok := metadata.FromIncomingContext(ss.Context())
		if !ok {
			return status.Error(codes.Unauthenticated, "missing metadata")
		}

		// Try API key authentication first
		if apiKeys := md.Get("x-api-key"); len(apiKeys) > 0 {
			if validateAPIKey(apiKeys[0], config.APIKeys) {
				return handler(srv, ss)
			}
			return status.Error(codes.Unauthenticated, "invalid API key")
		}

		// Try Bearer token authentication
		if authHeaders := md.Get("authorization"); len(authHeaders) > 0 {
			token := strings.TrimPrefix(authHeaders[0], "Bearer ")
			if token != authHeaders[0] {
				if validateToken(token, config.JWTSecret) {
					return handler(srv, ss)
				}
				return status.Error(codes.Unauthenticated, "invalid token")
			}
		}

		return status.Error(codes.Unauthenticated, "authentication required")
	}
}

// validateAPIKey checks if the provided key matches any valid key.
func validateAPIKey(key string, validKeys []string) bool {
	for _, validKey := range validKeys {
		if subtle.ConstantTimeCompare([]byte(key), []byte(validKey)) == 1 {
			return true
		}
	}
	return false
}

// validateToken validates a JWT token.
// TODO: Implement proper JWT validation with claims extraction.
func validateToken(token, secret string) bool {
	// For now, just check if token and secret are non-empty
	// In production, use a proper JWT library
	return len(token) > 0 && len(secret) > 0
}
