package client

import (
	"context"

	"github.com/mongorpc/mongorpc/lib/mongorpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Collection struct {
	client *MongoRPCClient
	name   string
	parent *Database
}

// Initilize a new collection
func (db *Database) Collection(name string) *Collection {

	// Create a new collection
	return &Collection{
		name:   name,
		client: db.client,
		parent: db,
	}
}

func (c *Collection) Insert(ctx context.Context, doc interface{}) (*primitive.ObjectID, error) {
	database := c.parent

	// crate mongorpc get document request
	resp, err := c.client.mongorpc.InsertDocument(ctx, &mongorpc.InsertDocumentRequest{
		Database:   database.name,
		Collection: c.name,
	})
	if err != nil {
		return nil, err
	}

	// decode mongorpc mongorpc result to interface
	result, err := primitive.ObjectIDFromHex(resp.Id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
