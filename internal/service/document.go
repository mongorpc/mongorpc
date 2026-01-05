// Package service provides business logic layer between API and repository.
package service

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/mongorpc/mongorpc/internal/repository/mongodb"
)

// DocumentService handles document operations.
type DocumentService struct {
	db *mongodb.Client
}

// NewDocumentService creates a new document service.
func NewDocumentService(db *mongodb.Client) *DocumentService {
	return &DocumentService{db: db}
}

// GetDocument retrieves a document by ID.
func (s *DocumentService) GetDocument(ctx context.Context, database, collection string, id bson.ObjectID) (bson.M, error) {
	coll := s.db.Client().Database(database).Collection(collection)
	filter := bson.D{{Key: "_id", Value: id}}

	var result bson.M
	err := coll.FindOne(ctx, filter).Decode(&result)
	return result, err
}

// CreateDocument inserts a new document.
func (s *DocumentService) CreateDocument(ctx context.Context, database, collection string, doc bson.D) (bson.ObjectID, error) {
	coll := s.db.Client().Database(database).Collection(collection)

	result, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return bson.ObjectID{}, err
	}

	if oid, ok := result.InsertedID.(bson.ObjectID); ok {
		return oid, nil
	}
	return bson.ObjectID{}, nil
}

// UpdateDocument updates a document by ID.
func (s *DocumentService) UpdateDocument(ctx context.Context, database, collection string, id bson.ObjectID, update bson.D, upsert bool) (int64, int64, error) {
	coll := s.db.Client().Database(database).Collection(collection)
	filter := bson.D{{Key: "_id", Value: id}}

	opts := options.UpdateOne()
	if upsert {
		opts.SetUpsert(true)
	}

	result, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return 0, 0, err
	}
	return result.MatchedCount, result.ModifiedCount, nil
}

// DeleteDocument deletes a document by ID.
func (s *DocumentService) DeleteDocument(ctx context.Context, database, collection string, id bson.ObjectID) (int64, error) {
	coll := s.db.Client().Database(database).Collection(collection)
	filter := bson.D{{Key: "_id", Value: id}}

	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

// ListDocuments lists documents with optional filter, limit, and skip.
func (s *DocumentService) ListDocuments(ctx context.Context, database, collection string, filter bson.D, limit, skip int64) ([]bson.M, error) {
	coll := s.db.Client().Database(database).Collection(collection)

	opts := options.Find()
	if limit > 0 {
		opts.SetLimit(limit)
	}
	if skip > 0 {
		opts.SetSkip(skip)
	}

	cursor, err := coll.Find(ctx, filter, opts)
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
