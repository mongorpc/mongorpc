// Package mongorpc provides additional RPC method implementations.
package mongorpc

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mongorpcv1 "github.com/mongorpc/mongorpc/gen/mongorpc/v1"
)

// UpdateDocument updates an existing document.
func (s *Server) UpdateDocument(ctx context.Context, req *mongorpcv1.UpdateDocumentRequest) (*mongorpcv1.UpdateDocumentResponse, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}
	if req.Id == nil || req.Id.Hex == "" {
		return nil, status.Error(codes.InvalidArgument, "document id is required")
	}
	if req.Update == nil {
		return nil, status.Error(codes.InvalidArgument, "update specification is required")
	}

	slog.Debug("UpdateDocument", "database", req.Database, "collection", req.Collection, "id", req.Id.Hex)

	// Parse ObjectID
	objectID, err := bson.ObjectIDFromHex(req.Id.Hex)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid object id: %v", err)
	}

	// Build filter
	filter := bson.D{{Key: "_id", Value: objectID}}

	// Build update
	update, err := updateSpecToBSON(req.Update)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid update: %v", err)
	}

	// Get collection
	coll := s.db.Client().Database(req.Database).Collection(req.Collection)

	// Build update options
	opts := options.UpdateOne()
	if req.Upsert {
		opts.SetUpsert(true)
	}

	// Execute update
	result, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update document: %v", err)
	}

	response := &mongorpcv1.UpdateDocumentResponse{
		WriteResult: &mongorpcv1.WriteResult{
			MatchedCount:  result.MatchedCount,
			ModifiedCount: result.ModifiedCount,
		},
	}

	// If requested, return the updated document
	if req.ReturnDocument {
		var updatedDoc bson.M
		err = coll.FindOne(ctx, filter).Decode(&updatedDoc)
		if err == nil {
			doc, _ := bsonToDocument(updatedDoc)
			response.Document = doc
		}
	}

	// Check for upserted ID
	if result.UpsertedID != nil {
		if upsertedOID, ok := result.UpsertedID.(bson.ObjectID); ok {
			response.WriteResult.UpsertedId = &mongorpcv1.ObjectId{Hex: upsertedOID.Hex()}
		}
	}

	return response, nil
}

// CountDocuments counts documents matching a filter.
func (s *Server) CountDocuments(ctx context.Context, req *mongorpcv1.CountDocumentsRequest) (*mongorpcv1.CountDocumentsResponse, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}

	slog.Debug("CountDocuments", "database", req.Database, "collection", req.Collection)

	// Get collection
	coll := s.db.Client().Database(req.Database).Collection(req.Collection)

	// Build filter
	filter := bson.D{}
	if req.Filter != nil {
		var err error
		filter, err = filterToBSON(req.Filter)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid filter: %v", err)
		}
	}

	// Build count options
	opts := options.Count()
	if req.Limit > 0 {
		opts.SetLimit(req.Limit)
	}
	if req.Skip > 0 {
		opts.SetSkip(req.Skip)
	}

	// Execute count
	count, err := coll.CountDocuments(ctx, filter, opts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to count documents: %v", err)
	}

	return &mongorpcv1.CountDocumentsResponse{
		Count: count,
	}, nil
}

// InsertMany inserts multiple documents.
func (s *Server) InsertMany(ctx context.Context, req *mongorpcv1.InsertManyRequest) (*mongorpcv1.InsertManyResponse, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}
	if len(req.Documents) == 0 {
		return nil, status.Error(codes.InvalidArgument, "at least one document is required")
	}

	slog.Debug("InsertMany", "database", req.Database, "collection", req.Collection, "count", len(req.Documents))

	// Convert documents
	docs := make([]interface{}, len(req.Documents))
	for i, doc := range req.Documents {
		bsonDoc, err := documentToBSON(doc)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid document at index %d: %v", i, err)
		}
		docs[i] = bsonDoc
	}

	// Get collection
	coll := s.db.Client().Database(req.Database).Collection(req.Collection)

	// Build options
	opts := options.InsertMany()
	if !req.Ordered {
		opts.SetOrdered(false)
	}

	// Execute insert
	result, err := coll.InsertMany(ctx, docs, opts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert documents: %v", err)
	}

	// Convert inserted IDs
	insertedIDs := make([]*mongorpcv1.ObjectId, len(result.InsertedIDs))
	for i, id := range result.InsertedIDs {
		if oid, ok := id.(bson.ObjectID); ok {
			insertedIDs[i] = &mongorpcv1.ObjectId{Hex: oid.Hex()}
		}
	}

	return &mongorpcv1.InsertManyResponse{
		InsertedIds: insertedIDs,
		WriteResult: &mongorpcv1.WriteResult{
			InsertedCount: int64(len(result.InsertedIDs)),
		},
	}, nil
}

// DeleteMany deletes multiple documents.
func (s *Server) DeleteMany(ctx context.Context, req *mongorpcv1.DeleteManyRequest) (*mongorpcv1.DeleteManyResponse, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}

	slog.Debug("DeleteMany", "database", req.Database, "collection", req.Collection)

	// Get collection
	coll := s.db.Client().Database(req.Database).Collection(req.Collection)

	// Build filter
	filter := bson.D{}
	if req.Filter != nil {
		var err error
		filter, err = filterToBSON(req.Filter)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid filter: %v", err)
		}
	}

	// Execute delete
	result, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete documents: %v", err)
	}

	return &mongorpcv1.DeleteManyResponse{
		WriteResult: &mongorpcv1.WriteResult{
			DeletedCount: result.DeletedCount,
		},
	}, nil
}

// updateSpecToBSON converts a proto UpdateSpec to BSON update.
func updateSpecToBSON(spec *mongorpcv1.UpdateSpec) (bson.D, error) {
	if spec == nil {
		return bson.D{}, nil
	}

	switch u := spec.UpdateType.(type) {
	case *mongorpcv1.UpdateSpec_Operators:
		return updateOperatorsToBSON(u.Operators)
	case *mongorpcv1.UpdateSpec_Pipeline:
		// Pipeline updates not yet supported
		return nil, status.Error(codes.Unimplemented, "pipeline updates not yet supported")
	default:
		return bson.D{}, nil
	}
}

// updateOperatorsToBSON converts UpdateOperators to BSON.
func updateOperatorsToBSON(ops *mongorpcv1.UpdateOperators) (bson.D, error) {
	if ops == nil {
		return bson.D{}, nil
	}

	result := bson.D{}

	// $set
	if len(ops.Set) > 0 {
		setDoc := bson.D{}
		for k, v := range ops.Set {
			bsonVal, err := valueToBSON(v)
			if err != nil {
				return nil, err
			}
			setDoc = append(setDoc, bson.E{Key: k, Value: bsonVal})
		}
		result = append(result, bson.E{Key: "$set", Value: setDoc})
	}

	// $unset
	if len(ops.Unset) > 0 {
		unsetDoc := bson.D{}
		for _, field := range ops.Unset {
			unsetDoc = append(unsetDoc, bson.E{Key: field, Value: ""})
		}
		result = append(result, bson.E{Key: "$unset", Value: unsetDoc})
	}

	// $inc
	if len(ops.Inc) > 0 {
		incDoc := bson.D{}
		for k, v := range ops.Inc {
			bsonVal, err := valueToBSON(v)
			if err != nil {
				return nil, err
			}
			incDoc = append(incDoc, bson.E{Key: k, Value: bsonVal})
		}
		result = append(result, bson.E{Key: "$inc", Value: incDoc})
	}

	return result, nil
}
