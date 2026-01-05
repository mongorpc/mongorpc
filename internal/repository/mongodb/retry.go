// Package mongodb provides retry logic for MongoDB operations.
package mongodb

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// RetryConfig configures retry behavior.
type RetryConfig struct {
	MaxRetries     int
	InitialBackoff time.Duration
	MaxBackoff     time.Duration
	Multiplier     float64
}

// DefaultRetryConfig returns sensible defaults.
func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxRetries:     3,
		InitialBackoff: 100 * time.Millisecond,
		MaxBackoff:     5 * time.Second,
		Multiplier:     2.0,
	}
}

// Retrier provides retry functionality for MongoDB operations.
type Retrier struct {
	config RetryConfig
}

// NewRetrier creates a new retrier.
func NewRetrier(config RetryConfig) *Retrier {
	if config.MaxRetries <= 0 {
		config.MaxRetries = 3
	}
	if config.InitialBackoff <= 0 {
		config.InitialBackoff = 100 * time.Millisecond
	}
	if config.MaxBackoff <= 0 {
		config.MaxBackoff = 5 * time.Second
	}
	if config.Multiplier <= 0 {
		config.Multiplier = 2.0
	}

	return &Retrier{config: config}
}

// Do executes a function with retry logic.
func (r *Retrier) Do(ctx context.Context, fn func() error) error {
	var lastErr error
	backoff := r.config.InitialBackoff

	for attempt := 0; attempt <= r.config.MaxRetries; attempt++ {
		if attempt > 0 {
			slog.Debug("Retrying operation", "attempt", attempt, "backoff", backoff)

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(backoff):
			}

			backoff = time.Duration(float64(backoff) * r.config.Multiplier)
			if backoff > r.config.MaxBackoff {
				backoff = r.config.MaxBackoff
			}
		}

		lastErr = fn()
		if lastErr == nil {
			return nil
		}

		if !isRetryable(lastErr) {
			return lastErr
		}
	}

	return lastErr
}

// DoWithResult executes a function that returns a result with retry logic.
func (r *Retrier) DoWithResult(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
	var result interface{}
	err := r.Do(ctx, func() error {
		var err error
		result, err = fn()
		return err
	})
	return result, err
}

// isRetryable determines if an error is retryable.
func isRetryable(err error) bool {
	if err == nil {
		return false
	}

	// Check for transient MongoDB errors
	if mongo.IsTimeout(err) {
		return true
	}
	if mongo.IsNetworkError(err) {
		return true
	}

	// Check for specific server errors that are retryable
	serverErr, ok := err.(mongo.ServerError)
	if ok {
		// Transient transaction errors
		if serverErr.HasErrorLabel("TransientTransactionError") {
			return true
		}
		// Retryable writes
		if serverErr.HasErrorLabel("RetryableWriteError") {
			return true
		}
	}

	return false
}

// RetryableFunc wraps a function with retry logic.
func RetryableFunc[T any](ctx context.Context, r *Retrier, fn func() (T, error)) (T, error) {
	var result T
	err := r.Do(ctx, func() error {
		var err error
		result, err = fn()
		return err
	})
	return result, err
}
