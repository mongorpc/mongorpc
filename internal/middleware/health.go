// Package middleware provides health check support for MongoRPC.
package middleware

import (
	"context"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

// HealthChecker manages service health status.
type HealthChecker struct {
	mu       sync.RWMutex
	services map[string]grpc_health_v1.HealthCheckResponse_ServingStatus
}

// NewHealthChecker creates a new health checker.
func NewHealthChecker() *HealthChecker {
	return &HealthChecker{
		services: make(map[string]grpc_health_v1.HealthCheckResponse_ServingStatus),
	}
}

// SetServingStatus sets the serving status for a service.
func (h *HealthChecker) SetServingStatus(service string, status grpc_health_v1.HealthCheckResponse_ServingStatus) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.services[service] = status
}

// GetServingStatus gets the serving status for a service.
func (h *HealthChecker) GetServingStatus(service string) grpc_health_v1.HealthCheckResponse_ServingStatus {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if status, ok := h.services[service]; ok {
		return status
	}
	return grpc_health_v1.HealthCheckResponse_SERVICE_UNKNOWN
}

// Check implements the gRPC health check service.
func (h *HealthChecker) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	status := h.GetServingStatus(req.Service)
	return &grpc_health_v1.HealthCheckResponse{Status: status}, nil
}

// Watch implements the streaming health check.
func (h *HealthChecker) Watch(req *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {
	// Simple implementation - send current status
	status := h.GetServingStatus(req.Service)
	return stream.Send(&grpc_health_v1.HealthCheckResponse{Status: status})
}

// HealthInterceptor creates a health check interceptor that rejects requests when unhealthy.
func HealthInterceptor(h *HealthChecker, serviceName string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		st := h.GetServingStatus(serviceName)
		if st != grpc_health_v1.HealthCheckResponse_SERVING {
			return nil, status.Error(codes.Unavailable, "service is not serving")
		}
		return handler(ctx, req)
	}
}
