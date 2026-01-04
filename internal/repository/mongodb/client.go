// Package mongodb provides MongoDB repository implementations.
package mongodb

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Client wraps the MongoDB client with convenience methods.
type Client struct {
	client *mongo.Client
	db     *mongo.Database
}

// New creates a new MongoDB client and connects to the database.
func New(ctx context.Context, uri, database string) (*Client, error) {
	slog.Info("Connecting to MongoDB", "uri", uri, "database", database)

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping to verify connection
	pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := client.Ping(pingCtx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	slog.Info("Connected to MongoDB successfully")

	return &Client{
		client: client,
		db:     client.Database(database),
	}, nil
}

// Close disconnects from MongoDB.
func (c *Client) Close(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}

// Database returns the database handle.
func (c *Client) Database() *mongo.Database {
	return c.db
}

// Collection returns a collection handle.
func (c *Client) Collection(name string) *mongo.Collection {
	return c.db.Collection(name)
}

// Client returns the underlying mongo client for advanced operations.
func (c *Client) Client() *mongo.Client {
	return c.client
}
