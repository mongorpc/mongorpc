// Package mongodb provides index operations.
package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// IndexInfo represents index information.
type IndexInfo struct {
	Name   string
	Keys   bson.D
	Unique bool
	Sparse bool
	TTL    *int32
}

// CreateIndex creates an index on a collection.
func (c *Client) CreateIndex(ctx context.Context, database, collection string, index IndexInfo) (string, error) {
	coll := c.client.Database(database).Collection(collection)

	indexModel := mongo.IndexModel{
		Keys: index.Keys,
	}

	indexOpts := options.Index()
	if index.Name != "" {
		indexOpts.SetName(index.Name)
	}
	if index.Unique {
		indexOpts.SetUnique(true)
	}
	if index.Sparse {
		indexOpts.SetSparse(true)
	}
	if index.TTL != nil {
		indexOpts.SetExpireAfterSeconds(*index.TTL)
	}
	indexModel.Options = indexOpts

	return coll.Indexes().CreateOne(ctx, indexModel)
}

// ListIndexes lists indexes on a collection.
func (c *Client) ListIndexes(ctx context.Context, database, collection string) ([]IndexInfo, error) {
	coll := c.client.Database(database).Collection(collection)

	cursor, err := coll.Indexes().List(ctx)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var indexes []IndexInfo
	for cursor.Next(ctx) {
		var idx bson.M
		if err := cursor.Decode(&idx); err != nil {
			continue
		}

		info := IndexInfo{}
		if name, ok := idx["name"].(string); ok {
			info.Name = name
		}
		if keys, ok := idx["key"].(bson.D); ok {
			info.Keys = keys
		}
		if unique, ok := idx["unique"].(bool); ok {
			info.Unique = unique
		}
		if sparse, ok := idx["sparse"].(bool); ok {
			info.Sparse = sparse
		}

		indexes = append(indexes, info)
	}

	return indexes, nil
}

// DropIndex drops an index by name.
func (c *Client) DropIndex(ctx context.Context, database, collection, indexName string) error {
	coll := c.client.Database(database).Collection(collection)
	return coll.Indexes().DropOne(ctx, indexName)
}
