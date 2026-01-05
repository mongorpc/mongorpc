// Package service provides query operations.
package service

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/mongorpc/mongorpc/internal/repository/mongodb"
)

// QueryService handles query operations.
type QueryService struct {
	db *mongodb.Client
}

// NewQueryService creates a new query service.
func NewQueryService(db *mongodb.Client) *QueryService {
	return &QueryService{db: db}
}

// QueryOptions contains options for queries.
type QueryOptions struct {
	Filter bson.D
	Sort   bson.D
	Limit  int64
	Skip   int64
}

// Find executes a find query.
func (s *QueryService) Find(ctx context.Context, database, collection string, opts QueryOptions) ([]bson.M, error) {
	coll := s.db.Client().Database(database).Collection(collection)

	findOpts := options.Find()
	if opts.Limit > 0 {
		findOpts.SetLimit(opts.Limit)
	}
	if opts.Skip > 0 {
		findOpts.SetSkip(opts.Skip)
	}
	if len(opts.Sort) > 0 {
		findOpts.SetSort(opts.Sort)
	}

	cursor, err := coll.Find(ctx, opts.Filter, findOpts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// Count counts documents matching the filter.
func (s *QueryService) Count(ctx context.Context, database, collection string, filter bson.D) (int64, error) {
	coll := s.db.Client().Database(database).Collection(collection)
	return coll.CountDocuments(ctx, filter)
}

// Distinct returns distinct values for a field.
func (s *QueryService) Distinct(ctx context.Context, database, collection, field string, filter bson.D) ([]interface{}, error) {
	coll := s.db.Client().Database(database).Collection(collection)

	result := coll.Distinct(ctx, field, filter)
	var values []interface{}
	if err := result.Decode(&values); err != nil {
		return nil, err
	}
	return values, nil
}

// AggregateOptions contains options for aggregation.
type AggregateOptions struct {
	Pipeline     bson.A
	AllowDiskUse bool
	BatchSize    int32
}

// Aggregate executes an aggregation pipeline.
func (s *QueryService) Aggregate(ctx context.Context, database, collection string, opts AggregateOptions) ([]bson.M, error) {
	coll := s.db.Client().Database(database).Collection(collection)

	aggOpts := options.Aggregate()
	if opts.AllowDiskUse {
		aggOpts.SetAllowDiskUse(true)
	}
	if opts.BatchSize > 0 {
		aggOpts.SetBatchSize(opts.BatchSize)
	}

	cursor, err := coll.Aggregate(ctx, opts.Pipeline, aggOpts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
