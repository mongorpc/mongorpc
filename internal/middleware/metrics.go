// Package middleware provides gRPC interceptors for MongoRPC.
package middleware

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

// Metrics holds the metrics for the server.
type Metrics struct {
	RequestCount    map[string]int64
	RequestDuration map[string]time.Duration
	ErrorCount      map[string]int64
}

// NewMetrics creates a new Metrics instance.
func NewMetrics() *Metrics {
	return &Metrics{
		RequestCount:    make(map[string]int64),
		RequestDuration: make(map[string]time.Duration),
		ErrorCount:      make(map[string]int64),
	}
}

// MetricsInterceptor collects request metrics.
func MetricsInterceptor(m *Metrics) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		// Call handler
		resp, err := handler(ctx, req)

		// Record metrics
		duration := time.Since(start)
		m.RequestCount[info.FullMethod]++
		m.RequestDuration[info.FullMethod] += duration
		if err != nil {
			m.ErrorCount[info.FullMethod]++
		}

		return resp, err
	}
}

// StreamMetricsInterceptor collects streaming request metrics.
func StreamMetricsInterceptor(m *Metrics) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()

		err := handler(srv, ss)

		// Record metrics
		duration := time.Since(start)
		m.RequestCount[info.FullMethod]++
		m.RequestDuration[info.FullMethod] += duration
		if err != nil {
			m.ErrorCount[info.FullMethod]++
		}

		return err
	}
}

// GetMetrics returns the current metrics snapshot.
func (m *Metrics) GetMetrics() map[string]interface{} {
	result := make(map[string]interface{})

	for method, count := range m.RequestCount {
		result[method] = map[string]interface{}{
			"count":          count,
			"total_duration": m.RequestDuration[method].String(),
			"error_count":    m.ErrorCount[method],
		}
	}

	return result
}
