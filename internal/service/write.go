// Package service provides write operations.
package service

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/mongorpc/mongorpc/internal/repository/mongodb"
)

// WriteService handles write operations.
type WriteService struct {
	db *mongodb.Client
}

// NewWriteService creates a new write service.
func NewWriteService(db *mongodb.Client) *WriteService {
	return &WriteService{db: db}
}

// InsertManyResult contains results from InsertMany.
type InsertManyResult struct {
	InsertedIDs []bson.ObjectID
}

// InsertMany inserts multiple documents.
func (s *WriteService) InsertMany(ctx context.Context, database, collection string, docs []interface{}, ordered bool) (*InsertManyResult, error) {
	coll := s.db.Client().Database(database).Collection(collection)

	opts := options.InsertMany()
	if !ordered {
		opts.SetOrdered(false)
	}

	result, err := coll.InsertMany(ctx, docs, opts)
	if err != nil {
		return nil, err
	}

	insertedIDs := make([]bson.ObjectID, len(result.InsertedIDs))
	for i, id := range result.InsertedIDs {
		if oid, ok := id.(bson.ObjectID); ok {
			insertedIDs[i] = oid
		}
	}

	return &InsertManyResult{InsertedIDs: insertedIDs}, nil
}

// UpdateManyResult contains results from UpdateMany.
type UpdateManyResult struct {
	MatchedCount  int64
	ModifiedCount int64
	UpsertedCount int64
	UpsertedID    *bson.ObjectID
}

// UpdateMany updates multiple documents.
func (s *WriteService) UpdateMany(ctx context.Context, database, collection string, filter, update bson.D, upsert bool) (*UpdateManyResult, error) {
	coll := s.db.Client().Database(database).Collection(collection)

	opts := options.UpdateMany()
	if upsert {
		opts.SetUpsert(true)
	}

	result, err := coll.UpdateMany(ctx, filter, update, opts)
	if err != nil {
		return nil, err
	}

	res := &UpdateManyResult{
		MatchedCount:  result.MatchedCount,
		ModifiedCount: result.ModifiedCount,
	}

	if result.UpsertedID != nil {
		if oid, ok := result.UpsertedID.(bson.ObjectID); ok {
			res.UpsertedID = &oid
			res.UpsertedCount = 1
		}
	}

	return res, nil
}

// DeleteMany deletes multiple documents.
func (s *WriteService) DeleteMany(ctx context.Context, database, collection string, filter bson.D) (int64, error) {
	coll := s.db.Client().Database(database).Collection(collection)

	result, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

// FindOneAndUpdate finds and updates a document, returning the result.
func (s *WriteService) FindOneAndUpdate(ctx context.Context, database, collection string, filter, update bson.D, returnNew bool) (bson.M, error) {
	coll := s.db.Client().Database(database).Collection(collection)

	opts := options.FindOneAndUpdate()
	if returnNew {
		opts.SetReturnDocument(options.After)
	}

	var result bson.M
	err := coll.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindOneAndDelete finds and deletes a document, returning the deleted document.
func (s *WriteService) FindOneAndDelete(ctx context.Context, database, collection string, filter bson.D) (bson.M, error) {
	coll := s.db.Client().Database(database).Collection(collection)

	var result bson.M
	err := coll.FindOneAndDelete(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
