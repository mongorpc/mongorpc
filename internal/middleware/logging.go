// Package middleware provides gRPC interceptors for MongoRPC.
package middleware

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// LoggingInterceptor logs all gRPC requests and responses.
func LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		// Extract request ID from metadata if present
		requestID := ""
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if ids := md.Get("x-request-id"); len(ids) > 0 {
				requestID = ids[0]
			}
		}

		// Log request
		slog.Info("gRPC request",
			"method", info.FullMethod,
			"request_id", requestID,
		)

		// Call handler
		resp, err := handler(ctx, req)

		// Log response
		duration := time.Since(start)
		logLevel := slog.LevelInfo
		if err != nil {
			logLevel = slog.LevelError
		}

		slog.Log(ctx, logLevel, "gRPC response",
			"method", info.FullMethod,
			"request_id", requestID,
			"duration_ms", duration.Milliseconds(),
			"error", err,
		)

		return resp, err
	}
}

// StreamLoggingInterceptor logs streaming gRPC requests.
func StreamLoggingInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()

		slog.Info("gRPC stream started",
			"method", info.FullMethod,
			"is_client_stream", info.IsClientStream,
			"is_server_stream", info.IsServerStream,
		)

		err := handler(srv, ss)

		duration := time.Since(start)
		slog.Info("gRPC stream ended",
			"method", info.FullMethod,
			"duration_ms", duration.Milliseconds(),
			"error", err,
		)

		return err
	}
}

// RecoveryInterceptor recovers from panics and returns internal errors.
func RecoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("Panic recovered in gRPC handler",
					"method", info.FullMethod,
					"panic", r,
				)
				err = status.Errorf(codes.Internal, "internal server error")
			}
		}()
		return handler(ctx, req)
	}
}

// StreamRecoveryInterceptor recovers from panics in streaming handlers.
func StreamRecoveryInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("Panic recovered in gRPC stream handler",
					"method", info.FullMethod,
					"panic", r,
				)
				err = status.Errorf(codes.Internal, "internal server error")
			}
		}()
		return handler(srv, ss)
	}
}
