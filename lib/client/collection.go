package client

import (
	"context"

	"github.com/mongorpc/mongorpc/lib/decoder"
	"github.com/mongorpc/mongorpc/lib/encoder"
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
		Document:   encoder.Encode(doc),
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

func (c *Collection) InsertMany(ctx context.Context, docs []interface{}) (interface{}, error) {
	database := c.parent

	documents := []*mongorpc.Value{}
	for _, doc := range docs {
		documents = append(documents, encoder.Encode(doc))
	}

	// crate mongorpc get document request
	resp, err := c.client.mongorpc.BulkInsertDocuments(ctx, &mongorpc.BulkInsertDocumentsRequest{
		Database:   database.name,
		Collection: c.name,
		Documents:  documents,
	})
	if err != nil {
		return nil, err
	}

	// decode mongorpc mongorpc result to interface
	return decoder.Decode(resp), nil
}
