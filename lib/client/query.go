package client

import (
	"context"

	"github.com/mongorpc/mongorpc/lib/decoder"
	"github.com/mongorpc/mongorpc/lib/encoder"
	"github.com/mongorpc/mongorpc/lib/mongorpc"
)

type QueryBuilder struct {
	client     *MongoRPCClient
	collection *Collection
	limit      int32
	skip       int32
	sort       map[string]interface{}
	filter     map[string]map[string]interface{}
}
type Order int32

const (
	ASCENDING  Order = 1
	DESCENDING Order = -1
)

func (q *QueryBuilder) Limit(limit int32) *QueryBuilder {
	q.limit = limit
	return q
}

func (q *QueryBuilder) Skip(skip int32) *QueryBuilder {
	q.skip = skip
	return q
}

func (q *QueryBuilder) Sort(field string, order Order) *QueryBuilder {
	if q.sort == nil {
		q.sort = make(map[string]interface{})
	}
	q.sort[field] = int32(order)
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
		Sort:       encoder.Encode(q.sort),
		Query:      encoder.Encode(q.filter),
	})
	if err != nil {
		return nil, err
	}

	// decode mongorpc mongorpc result to interface
	result := decoder.Decode(resp)
	return result, nil
}
