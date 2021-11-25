package client

import (
	"context"

	"github.com/mongorpc/mongorpc/lib/decoder"
	"github.com/mongorpc/mongorpc/lib/encoder"
	"github.com/mongorpc/mongorpc/lib/mongorpc"
)

// Document is a wrapper for a document in a collection.
type Document struct {
	client     *MongoRPCClient
	documentID string
	parent     *Collection
}

// Initilize the document with the document id.
func (coll *Collection) Document(documentID string) *Document {
	//	return document with the given id.
	return &Document{
		documentID: documentID,
		client:     coll.client,
		parent:     coll,
	}
}

// Get returns the document with the given id.
func (doc *Document) Get(ctx context.Context) (interface{}, error) {

	database := doc.parent.parent
	collection := doc.parent

	// crate mongorpc get document request
	resp, err := doc.client.mongorpc.GetDocument(ctx, &mongorpc.GetDocumentRequest{
		Database:   database.name,
		Collection: collection.name,
		DocumentId: &mongorpc.ObjectId{
			Id: doc.documentID,
		},
	})
	if err != nil {
		return nil, err
	}

	// decode mongorpc mongorpc result to interface
	result := decoder.Decode(resp)
	return result, nil
}

// Update updates the document with the given id.
func (doc *Document) Update(ctx context.Context, data interface{}) (interface{}, error) {

	database := doc.parent.parent
	collection := doc.parent

	// crate mongorpc get document request
	resp, err := doc.client.mongorpc.UpdateDocument(ctx, &mongorpc.UpdateDocumentRequest{
		Database:   database.name,
		Collection: collection.name,
		DocumentId: &mongorpc.ObjectId{
			Id: doc.documentID,
		},
		Document: encoder.Encode(data),
		Replace:  false,
	})
	if err != nil {
		return nil, err
	}

	// decode mongorpc mongorpc result to interface
	result := decoder.Decode(resp)
	return result, nil
}

// Replace replaces the document with the given id.
func (doc *Document) Replace(ctx context.Context, data interface{}) (interface{}, error) {

	database := doc.parent.parent
	collection := doc.parent

	// crate mongorpc get document request
	resp, err := doc.client.mongorpc.UpdateDocument(ctx, &mongorpc.UpdateDocumentRequest{
		Database:   database.name,
		Collection: collection.name,
		DocumentId: &mongorpc.ObjectId{
			Id: doc.documentID,
		},
		Document: encoder.Encode(data),
		Replace:  true,
	})
	if err != nil {
		return nil, err
	}

	// decode mongorpc mongorpc result to interface
	result := decoder.Decode(resp)
	return result, nil
}

// Delete deletes the document with the given id.
func (doc *Document) Delete(ctx context.Context) (interface{}, error) {

	database := doc.parent.parent
	collection := doc.parent

	// crate mongorpc get document request
	resp, err := doc.client.mongorpc.DeleteDocument(ctx, &mongorpc.DeleteDocumentRequest{
		Database:   database.name,
		Collection: collection.name,
		DocumentId: &mongorpc.ObjectId{
			Id: doc.documentID,
		},
	})
	if err != nil {
		return nil, err
	}

	// decode mongorpc mongorpc result to interface
	result := decoder.Decode(resp)
	return result, nil
}
