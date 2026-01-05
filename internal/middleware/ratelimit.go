// Package middleware provides rate limiting for MongoRPC.
package middleware

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// RateLimiter provides token bucket rate limiting.
type RateLimiter struct {
	mu       sync.Mutex
	tokens   map[string]*bucket
	rate     float64 // tokens per second
	burst    int     // max tokens
	cleanupT time.Duration
}

type bucket struct {
	tokens    float64
	lastCheck time.Time
}

// NewRateLimiter creates a new rate limiter.
func NewRateLimiter(rate float64, burst int) *RateLimiter {
	rl := &RateLimiter{
		tokens:   make(map[string]*bucket),
		rate:     rate,
		burst:    burst,
		cleanupT: time.Minute * 5,
	}
	go rl.cleanup()
	return rl
}

// Allow checks if a request is allowed for the given key.
func (rl *RateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	b, exists := rl.tokens[key]
	if !exists {
		rl.tokens[key] = &bucket{
			tokens:    float64(rl.burst) - 1,
			lastCheck: now,
		}
		return true
	}

	// Add tokens based on elapsed time
	elapsed := now.Sub(b.lastCheck).Seconds()
	b.tokens += elapsed * rl.rate
	if b.tokens > float64(rl.burst) {
		b.tokens = float64(rl.burst)
	}
	b.lastCheck = now

	// Check if we can consume a token
	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}

func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(rl.cleanupT)
	for range ticker.C {
		rl.mu.Lock()
		cutoff := time.Now().Add(-rl.cleanupT)
		for k, b := range rl.tokens {
			if b.lastCheck.Before(cutoff) {
				delete(rl.tokens, k)
			}
		}
		rl.mu.Unlock()
	}
}

// RateLimitInterceptor creates a rate limiting interceptor.
func RateLimitInterceptor(rl *RateLimiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Extract client identifier from metadata
		key := "default"
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if ids := md.Get("x-client-id"); len(ids) > 0 {
				key = ids[0]
			} else if ips := md.Get("x-forwarded-for"); len(ips) > 0 {
				key = ips[0]
			}
		}

		if !rl.Allow(key) {
			return nil, status.Error(codes.ResourceExhausted, "rate limit exceeded")
		}
		return handler(ctx, req)
	}
}

// StreamRateLimitInterceptor creates a streaming rate limiting interceptor.
func StreamRateLimitInterceptor(rl *RateLimiter) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		key := "default"
		if md, ok := metadata.FromIncomingContext(ss.Context()); ok {
			if ids := md.Get("x-client-id"); len(ids) > 0 {
				key = ids[0]
			}
		}

		if !rl.Allow(key) {
			return status.Error(codes.ResourceExhausted, "rate limit exceeded")
		}
		return handler(srv, ss)
	}
}
