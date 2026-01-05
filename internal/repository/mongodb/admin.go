// Package mongodb provides admin operations.
package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// DatabaseInfo represents database information.
type DatabaseInfo struct {
	Name       string
	SizeOnDisk int64
	Empty      bool
}

// ListDatabases lists all databases.
func (c *Client) ListDatabases(ctx context.Context) ([]DatabaseInfo, error) {
	result, err := c.client.ListDatabases(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	databases := make([]DatabaseInfo, len(result.Databases))
	for i, db := range result.Databases {
		databases[i] = DatabaseInfo{
			Name:       db.Name,
			SizeOnDisk: db.SizeOnDisk,
			Empty:      db.Empty,
		}
	}
	return databases, nil
}

// DropDatabase drops a database.
func (c *Client) DropDatabase(ctx context.Context, database string) error {
	return c.client.Database(database).Drop(ctx)
}

// ServerStatus returns server status.
func (c *Client) ServerStatus(ctx context.Context) (bson.M, error) {
	var result bson.M
	err := c.client.Database("admin").RunCommand(ctx, bson.D{{Key: "serverStatus", Value: 1}}).Decode(&result)
	return result, err
}

// GetBuildInfo returns MongoDB build information.
func (c *Client) GetBuildInfo(ctx context.Context) (bson.M, error) {
	var result bson.M
	err := c.client.Database("admin").RunCommand(ctx, bson.D{{Key: "buildInfo", Value: 1}}).Decode(&result)
	return result, err
}

// RunCommand runs an arbitrary command on a database.
func (c *Client) RunCommand(ctx context.Context, database string, cmd bson.D) (bson.M, error) {
	var result bson.M
	err := c.client.Database(database).RunCommand(ctx, cmd).Decode(&result)
	return result, err
}

// CollStats returns collection statistics.
func (c *Client) CollStats(ctx context.Context, database, collection string) (bson.M, error) {
	cmd := bson.D{{Key: "collStats", Value: collection}}
	var result bson.M
	err := c.client.Database(database).RunCommand(ctx, cmd).Decode(&result)
	return result, err
}

// DBStats returns database statistics.
func (c *Client) DBStats(ctx context.Context, database string) (bson.M, error) {
	cmd := bson.D{{Key: "dbStats", Value: 1}}
	var result bson.M
	err := c.client.Database(database).RunCommand(ctx, cmd).Decode(&result)
	return result, err
}
