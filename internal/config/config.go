// Package config provides configuration management for MongoRPC server.
package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds the server configuration.
type Config struct {
	// Server settings
	ServerAddress string

	// MongoDB settings
	MongoDBURI      string
	MongoDBDatabase string

	// Timeouts
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

// Load creates a Config from environment variables with sensible defaults.
func Load() *Config {
	return &Config{
		ServerAddress:   getEnv("SERVER_ADDRESS", "localhost:50051"),
		MongoDBURI:      getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		MongoDBDatabase: getEnv("MONGODB_DATABASE", "mongorpc"),
		ReadTimeout:     getDuration("READ_TIMEOUT", 30*time.Second),
		WriteTimeout:    getDuration("WRITE_TIMEOUT", 30*time.Second),
		ShutdownTimeout: getDuration("SHUTDOWN_TIMEOUT", 10*time.Second),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if seconds, err := strconv.Atoi(value); err == nil {
			return time.Duration(seconds) * time.Second
		}
	}
	return defaultValue
}
