package client

import (
	"context"

	"github.com/mongorpc/mongorpc/lib/decoder"
	"github.com/mongorpc/mongorpc/lib/mongorpc"
)

type QueryBuilder struct {
	client     *MongoRPCClient
	collection *Collection
	limit      int32
	skip       int32
}

func (q *QueryBuilder) Limit(limit int32) *QueryBuilder {
	q.limit = limit
	return q
}

func (q *QueryBuilder) Skip(skip int32) *QueryBuilder {
	q.skip = skip
	return q
}

// Get returns the document with the given id.
func (q *QueryBuilder) Get(ctx context.Context) (interface{}, error) {

	database := q.collection.parent
	collection := q.collection

	// crate mongorpc get document request
	resp, err := q.client.mongorpc.QueryDocuments(ctx, &mongorpc.QueryDocumentsRequest{
		Database:   database.name,
		Collection: collection.name,
		Limit:      q.limit,
		Skip:       q.skip,
	})
	if err != nil {
		return nil, err
	}

	// decode mongorpc mongorpc result to interface
	result := decoder.Decode(resp)
	return result, nil
}
