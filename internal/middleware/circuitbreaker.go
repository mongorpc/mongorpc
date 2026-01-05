// Package middleware provides circuit breaker for MongoRPC.
package middleware

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CircuitState represents the state of the circuit breaker.
type CircuitState int

const (
	CircuitClosed CircuitState = iota
	CircuitOpen
	CircuitHalfOpen
)

func (s CircuitState) String() string {
	switch s {
	case CircuitClosed:
		return "closed"
	case CircuitOpen:
		return "open"
	case CircuitHalfOpen:
		return "half-open"
	default:
		return "unknown"
	}
}

// CircuitBreaker implements the circuit breaker pattern.
type CircuitBreaker struct {
	mu              sync.RWMutex
	state           CircuitState
	failures        int
	successes       int
	lastFailureTime time.Time

	// Configuration
	failureThreshold int
	successThreshold int
	timeout          time.Duration
	onStateChange    func(from, to CircuitState)
}

// CircuitBreakerConfig configures the circuit breaker.
type CircuitBreakerConfig struct {
	FailureThreshold int           // Failures before opening
	SuccessThreshold int           // Successes in half-open to close
	Timeout          time.Duration // Time before half-open
	OnStateChange    func(from, to CircuitState)
}

// NewCircuitBreaker creates a new circuit breaker.
func NewCircuitBreaker(config CircuitBreakerConfig) *CircuitBreaker {
	if config.FailureThreshold <= 0 {
		config.FailureThreshold = 5
	}
	if config.SuccessThreshold <= 0 {
		config.SuccessThreshold = 3
	}
	if config.Timeout <= 0 {
		config.Timeout = 30 * time.Second
	}

	return &CircuitBreaker{
		state:            CircuitClosed,
		failureThreshold: config.FailureThreshold,
		successThreshold: config.SuccessThreshold,
		timeout:          config.Timeout,
		onStateChange:    config.OnStateChange,
	}
}

// State returns the current circuit state.
func (cb *CircuitBreaker) State() CircuitState {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}

// Allow checks if the request should be allowed.
func (cb *CircuitBreaker) Allow() bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case CircuitClosed:
		return true
	case CircuitOpen:
		if time.Since(cb.lastFailureTime) > cb.timeout {
			cb.setState(CircuitHalfOpen)
			return true
		}
		return false
	case CircuitHalfOpen:
		return true
	}
	return false
}

// RecordSuccess records a successful call.
func (cb *CircuitBreaker) RecordSuccess() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case CircuitHalfOpen:
		cb.successes++
		if cb.successes >= cb.successThreshold {
			cb.setState(CircuitClosed)
		}
	case CircuitClosed:
		cb.failures = 0
	}
}

// RecordFailure records a failed call.
func (cb *CircuitBreaker) RecordFailure() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.failures++
	cb.lastFailureTime = time.Now()

	switch cb.state {
	case CircuitClosed:
		if cb.failures >= cb.failureThreshold {
			cb.setState(CircuitOpen)
		}
	case CircuitHalfOpen:
		cb.setState(CircuitOpen)
	}
}

func (cb *CircuitBreaker) setState(newState CircuitState) {
	if cb.state == newState {
		return
	}

	oldState := cb.state
	cb.state = newState
	cb.failures = 0
	cb.successes = 0

	slog.Info("Circuit breaker state changed", "from", oldState, "to", newState)

	if cb.onStateChange != nil {
		go cb.onStateChange(oldState, newState)
	}
}

// CircuitBreakerInterceptor creates a circuit breaker interceptor.
func CircuitBreakerInterceptor(cb *CircuitBreaker) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !cb.Allow() {
			return nil, status.Error(codes.Unavailable, "circuit breaker is open")
		}

		resp, err := handler(ctx, req)
		if err != nil {
			if isCircuitBreakerError(err) {
				cb.RecordFailure()
			}
		} else {
			cb.RecordSuccess()
		}

		return resp, err
	}
}

// StreamCircuitBreakerInterceptor creates a streaming circuit breaker interceptor.
func StreamCircuitBreakerInterceptor(cb *CircuitBreaker) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if !cb.Allow() {
			return status.Error(codes.Unavailable, "circuit breaker is open")
		}

		err := handler(srv, ss)
		if err != nil {
			if isCircuitBreakerError(err) {
				cb.RecordFailure()
			}
		} else {
			cb.RecordSuccess()
		}

		return err
	}
}

// isCircuitBreakerError determines if an error should trip the circuit.
func isCircuitBreakerError(err error) bool {
	if err == nil {
		return false
	}

	code := status.Code(err)
	switch code {
	case codes.Unavailable, codes.ResourceExhausted, codes.DeadlineExceeded, codes.Internal:
		return true
	}
	return false
}
