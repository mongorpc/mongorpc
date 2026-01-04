package middleware

import (
	"context"
	"testing"

	"google.golang.org/grpc"
)

func TestMetrics_RequestCounting(t *testing.T) {
	m := NewMetrics()

	// Simulate interceptor calls
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return "response", nil
	}

	interceptor := MetricsInterceptor(m)

	// Call the interceptor multiple times
	info := &grpc.UnaryServerInfo{FullMethod: "/test.Service/Method"}

	for i := 0; i < 5; i++ {
		_, _ = interceptor(context.Background(), nil, info, handler)
	}

	// Check metrics
	if m.RequestCount["/test.Service/Method"] != 5 {
		t.Errorf("expected 5 requests, got %d", m.RequestCount["/test.Service/Method"])
	}

	if m.ErrorCount["/test.Service/Method"] != 0 {
		t.Errorf("expected 0 errors, got %d", m.ErrorCount["/test.Service/Method"])
	}
}

func TestMetrics_ErrorCounting(t *testing.T) {
	m := NewMetrics()

	// Handler that returns an error
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, context.Canceled
	}

	interceptor := MetricsInterceptor(m)
	info := &grpc.UnaryServerInfo{FullMethod: "/test.Service/ErrorMethod"}

	_, _ = interceptor(context.Background(), nil, info, handler)

	if m.ErrorCount["/test.Service/ErrorMethod"] != 1 {
		t.Errorf("expected 1 error, got %d", m.ErrorCount["/test.Service/ErrorMethod"])
	}
}

func TestAuthInterceptor_APIKey(t *testing.T) {
	config := &AuthConfig{
		APIKeys:     []string{"valid-api-key"},
		SkipMethods: []string{"/test.Service/Public"},
	}

	interceptor := AuthInterceptor(config)

	// Handler that should be called
	handlerCalled := false
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		handlerCalled = true
		return "success", nil
	}

	info := &grpc.UnaryServerInfo{FullMethod: "/test.Service/Public"}

	// Test skip method
	_, err := interceptor(context.Background(), nil, info, handler)
	if err != nil {
		t.Errorf("expected no error for skip method, got %v", err)
	}
	if !handlerCalled {
		t.Error("expected handler to be called for skip method")
	}
}

func TestValidateAPIKey(t *testing.T) {
	validKeys := []string{"key1", "key2", "key3"}

	tests := []struct {
		key      string
		expected bool
	}{
		{"key1", true},
		{"key2", true},
		{"key3", true},
		{"invalid", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			result := validateAPIKey(tt.key, validKeys)
			if result != tt.expected {
				t.Errorf("validateAPIKey(%q) = %v, expected %v", tt.key, result, tt.expected)
			}
		})
	}
}

func TestRecoveryInterceptor(t *testing.T) {
	interceptor := RecoveryInterceptor()

	// Handler that panics
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		panic("test panic")
	}

	info := &grpc.UnaryServerInfo{FullMethod: "/test.Service/Panic"}

	// Should recover and return error
	_, err := interceptor(context.Background(), nil, info, handler)
	if err == nil {
		t.Error("expected error after panic recovery")
	}
}
