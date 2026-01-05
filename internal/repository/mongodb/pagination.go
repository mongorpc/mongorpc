// Package mongodb provides cursor-based pagination.
package mongodb

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// PageToken represents an opaque pagination token.
type PageToken struct {
	LastID    bson.ObjectID `json:"id,omitempty"`
	LastValue interface{}   `json:"v,omitempty"`
	SortField string        `json:"f,omitempty"`
	Timestamp time.Time     `json:"ts"`
}

// Encode encodes the page token to a string.
func (p *PageToken) Encode() string {
	data, _ := json.Marshal(p)
	return base64.URLEncoding.EncodeToString(data)
}

// DecodePageToken decodes a page token string.
func DecodePageToken(token string) (*PageToken, error) {
	if token == "" {
		return nil, nil
	}

	data, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	var pt PageToken
	if err := json.Unmarshal(data, &pt); err != nil {
		return nil, err
	}

	return &pt, nil
}

// PaginationOptions configures pagination.
type PaginationOptions struct {
	PageSize  int32
	PageToken string
	SortField string
	SortDesc  bool
}

// PaginatedResult contains paginated results.
type PaginatedResult struct {
	Documents     []bson.M
	NextPageToken string
	HasMore       bool
}

// Paginate performs a cursor-based paginated query.
func (c *Client) Paginate(ctx context.Context, database, collection string, filter bson.D, opts PaginationOptions) (*PaginatedResult, error) {
	coll := c.client.Database(database).Collection(collection)

	if opts.PageSize <= 0 {
		opts.PageSize = 20
	}
	if opts.PageSize > 1000 {
		opts.PageSize = 1000
	}

	// Parse page token if provided
	var pageToken *PageToken
	if opts.PageToken != "" {
		var err error
		pageToken, err = DecodePageToken(opts.PageToken)
		if err != nil {
			return nil, err
		}
	}

	// Build filter with cursor position
	queryFilter := filter
	if pageToken != nil {
		sortField := opts.SortField
		if sortField == "" {
			sortField = "_id"
		}

		var cursorFilter bson.D
		if opts.SortDesc {
			cursorFilter = bson.D{{Key: sortField, Value: bson.D{{Key: "$lt", Value: pageToken.LastValue}}}}
		} else {
			cursorFilter = bson.D{{Key: sortField, Value: bson.D{{Key: "$gt", Value: pageToken.LastValue}}}}
		}

		if len(queryFilter) > 0 {
			queryFilter = bson.D{{Key: "$and", Value: bson.A{filter, cursorFilter}}}
		} else {
			queryFilter = cursorFilter
		}
	}

	// Build sort
	sortDir := 1
	if opts.SortDesc {
		sortDir = -1
	}
	sortField := opts.SortField
	if sortField == "" {
		sortField = "_id"
	}
	sort := bson.D{{Key: sortField, Value: sortDir}}

	// Query one extra to detect more pages
	findOpts := options.Find().
		SetLimit(int64(opts.PageSize + 1)).
		SetSort(sort)

	cursor, err := coll.Find(ctx, queryFilter, findOpts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var docs []bson.M
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	result := &PaginatedResult{
		Documents: docs,
		HasMore:   len(docs) > int(opts.PageSize),
	}

	// Trim to page size and generate next token
	if result.HasMore {
		result.Documents = docs[:opts.PageSize]
		lastDoc := docs[opts.PageSize-1]

		var lastValue interface{}
		if sortField == "_id" {
			lastValue = lastDoc["_id"]
		} else {
			lastValue = lastDoc[sortField]
		}

		nextToken := &PageToken{
			LastValue: lastValue,
			SortField: sortField,
			Timestamp: time.Now(),
		}
		if oid, ok := lastDoc["_id"].(bson.ObjectID); ok {
			nextToken.LastID = oid
		}

		result.NextPageToken = nextToken.Encode()
	}

	return result, nil
}

// StreamingPaginator provides streaming pagination for large result sets.
type StreamingPaginator struct {
	coll      *mongo.Collection
	filter    bson.D
	sort      bson.D
	batchSize int32
	cursor    *mongo.Cursor
}

// NewStreamingPaginator creates a streaming paginator.
func (c *Client) NewStreamingPaginator(database, collection string, filter, sort bson.D, batchSize int32) *StreamingPaginator {
	if batchSize <= 0 {
		batchSize = 100
	}

	return &StreamingPaginator{
		coll:      c.client.Database(database).Collection(collection),
		filter:    filter,
		sort:      sort,
		batchSize: batchSize,
	}
}

// Start initializes the cursor.
func (p *StreamingPaginator) Start(ctx context.Context) error {
	opts := options.Find().
		SetSort(p.sort).
		SetBatchSize(p.batchSize)

	var err error
	p.cursor, err = p.coll.Find(ctx, p.filter, opts)
	return err
}

// Next returns the next batch of documents.
func (p *StreamingPaginator) Next(ctx context.Context) ([]bson.M, error) {
	if p.cursor == nil {
		return nil, nil
	}

	var batch []bson.M
	for i := int32(0); i < p.batchSize && p.cursor.Next(ctx); i++ {
		var doc bson.M
		if err := p.cursor.Decode(&doc); err != nil {
			return batch, err
		}
		batch = append(batch, doc)
	}

	if err := p.cursor.Err(); err != nil {
		return batch, err
	}

	return batch, nil
}

// Close closes the paginator.
func (p *StreamingPaginator) Close(ctx context.Context) error {
	if p.cursor != nil {
		return p.cursor.Close(ctx)
	}
	return nil
}
