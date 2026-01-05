// Package mongodb provides collection operations.
package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// CollectionInfo represents collection information.
type CollectionInfo struct {
	Name    string
	Type    string
	Options bson.M
}

// CreateCollection creates a new collection.
func (c *Client) CreateCollection(ctx context.Context, database, collection string) error {
	return c.client.Database(database).CreateCollection(ctx, collection)
}

// CreateCollectionWithOptions creates a collection with options.
func (c *Client) CreateCollectionWithOptions(ctx context.Context, database, collection string, capped bool, maxSize, maxDocs int64) error {
	opts := options.CreateCollection()
	if capped {
		opts.SetCapped(true)
		if maxSize > 0 {
			opts.SetSizeInBytes(maxSize)
		}
		if maxDocs > 0 {
			opts.SetMaxDocuments(maxDocs)
		}
	}

	return c.client.Database(database).CreateCollection(ctx, collection, opts)
}

// DropCollection drops a collection.
func (c *Client) DropCollection(ctx context.Context, database, collection string) error {
	return c.client.Database(database).Collection(collection).Drop(ctx)
}

// ListCollections lists collections in a database.
func (c *Client) ListCollections(ctx context.Context, database string, includeSystem bool) ([]CollectionInfo, error) {
	filter := bson.D{}
	if !includeSystem {
		filter = bson.D{{Key: "name", Value: bson.D{{Key: "$regex", Value: "^(?!system\\.)"}}}}
	}

	cursor, err := c.client.Database(database).ListCollections(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var collections []CollectionInfo
	for cursor.Next(ctx) {
		var coll bson.M
		if err := cursor.Decode(&coll); err != nil {
			continue
		}

		info := CollectionInfo{}
		if name, ok := coll["name"].(string); ok {
			info.Name = name
		}
		if typeStr, ok := coll["type"].(string); ok {
			info.Type = typeStr
		}
		if opts, ok := coll["options"].(bson.M); ok {
			info.Options = opts
		}

		collections = append(collections, info)
	}

	return collections, nil
}

// RenameCollection renames a collection.
func (c *Client) RenameCollection(ctx context.Context, database, oldName, newName string) error {
	cmd := bson.D{
		{Key: "renameCollection", Value: database + "." + oldName},
		{Key: "to", Value: database + "." + newName},
	}
	return c.client.Database("admin").RunCommand(ctx, cmd).Err()
}
