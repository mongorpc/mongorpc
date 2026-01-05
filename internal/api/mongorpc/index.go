// Package mongorpc provides the gRPC service implementations.
package mongorpc

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	mongorpcv1 "github.com/mongorpc/mongorpc/gen/mongorpc/v1"
)

// ListIndexes lists all indexes on a collection.
func (s *Server) ListIndexes(ctx context.Context, req *mongorpcv1.ListIndexesRequest) (*mongorpcv1.ListIndexesResponse, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}

	// Admin-only operation
	if !s.isAdmin(ctx) {
		return nil, status.Error(codes.PermissionDenied, "admin privileges required for index operations")
	}

	slog.Debug("ListIndexes", "database", req.Database, "collection", req.Collection)

	collection := s.db.Client().Database(req.Database).Collection(req.Collection)
	cursor, err := collection.Indexes().List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list indexes: %v", err)
	}
	defer cursor.Close(ctx)

	var indexes []*mongorpcv1.IndexInfo
	for cursor.Next(ctx) {
		var indexDoc bson.M
		if err := cursor.Decode(&indexDoc); err != nil {
			continue
		}

		indexInfo := &mongorpcv1.IndexInfo{
			Name: indexDoc["name"].(string),
		}

		// Parse key specification
		if key, ok := indexDoc["key"].(bson.M); ok {
			for field, direction := range key {
				indexKey := &mongorpcv1.IndexKey{
					Field: field,
				}

				switch v := direction.(type) {
				case int32:
					if v == 1 {
						indexKey.KeyType = &mongorpcv1.IndexKey_Direction{Direction: mongorpcv1.SortDirection_ASCENDING}
					} else {
						indexKey.KeyType = &mongorpcv1.IndexKey_Direction{Direction: mongorpcv1.SortDirection_DESCENDING}
					}
				case int:
					if v == 1 {
						indexKey.KeyType = &mongorpcv1.IndexKey_Direction{Direction: mongorpcv1.SortDirection_ASCENDING}
					} else {
						indexKey.KeyType = &mongorpcv1.IndexKey_Direction{Direction: mongorpcv1.SortDirection_DESCENDING}
					}
				case string:
					indexKey.KeyType = &mongorpcv1.IndexKey_Type{Type: v}
				}

				indexInfo.Keys = append(indexInfo.Keys, indexKey)
			}
		}

		// Parse options
		if unique, ok := indexDoc["unique"].(bool); ok {
			indexInfo.Unique = unique
		}
		if sparse, ok := indexDoc["sparse"].(bool); ok {
			indexInfo.Sparse = sparse
		}
		if ttl, ok := indexDoc["expireAfterSeconds"].(int32); ok {
			indexInfo.ExpireAfterSeconds = int64(ttl)
		}

		indexes = append(indexes, indexInfo)
	}

	return &mongorpcv1.ListIndexesResponse{
		Indexes: indexes,
	}, nil
}

// CreateIndex creates an index on a collection.
func (s *Server) CreateIndex(ctx context.Context, req *mongorpcv1.CreateIndexRequest) (*mongorpcv1.CreateIndexResponse, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}
	if len(req.Keys) == 0 {
		return nil, status.Error(codes.InvalidArgument, "at least one index key is required")
	}

	// Admin-only operation
	if !s.isAdmin(ctx) {
		return nil, status.Error(codes.PermissionDenied, "admin privileges required for index operations")
	}

	slog.Debug("CreateIndex", "database", req.Database, "collection", req.Collection, "keys", len(req.Keys))

	// Build index keys
	keys := bson.D{}
	for _, k := range req.Keys {
		switch keyType := k.KeyType.(type) {
		case *mongorpcv1.IndexKey_Direction:
			if keyType.Direction == mongorpcv1.SortDirection_ASCENDING {
				keys = append(keys, bson.E{Key: k.Field, Value: 1})
			} else {
				keys = append(keys, bson.E{Key: k.Field, Value: -1})
			}
		case *mongorpcv1.IndexKey_Type:
			keys = append(keys, bson.E{Key: k.Field, Value: keyType.Type})
		default:
			keys = append(keys, bson.E{Key: k.Field, Value: 1})
		}
	}

	// Build index options
	indexOpts := options.Index()
	if req.Options != nil {
		if req.Options.Name != "" {
			indexOpts.SetName(req.Options.Name)
		}
		if req.Options.Unique {
			indexOpts.SetUnique(true)
		}
		if req.Options.Sparse {
			indexOpts.SetSparse(true)
		}
		if req.Options.ExpireAfterSeconds > 0 {
			indexOpts.SetExpireAfterSeconds(int32(req.Options.ExpireAfterSeconds))
		}
		if req.Options.Hidden {
			indexOpts.SetHidden(true)
		}
	}

	// Create the index
	collection := s.db.Client().Database(req.Database).Collection(req.Collection)
	indexModel := mongo.IndexModel{
		Keys:    keys,
		Options: indexOpts,
	}

	indexName, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create index: %v", err)
	}

	slog.Info("Index created", "database", req.Database, "collection", req.Collection, "name", indexName)

	return &mongorpcv1.CreateIndexResponse{
		IndexName: indexName,
	}, nil
}

// DropIndex removes an index from a collection.
func (s *Server) DropIndex(ctx context.Context, req *mongorpcv1.DropIndexRequest) (*emptypb.Empty, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}
	if req.IndexName == "" {
		return nil, status.Error(codes.InvalidArgument, "index name is required")
	}

	// Admin-only operation
	if !s.isAdmin(ctx) {
		return nil, status.Error(codes.PermissionDenied, "admin privileges required for index operations")
	}

	slog.Debug("DropIndex", "database", req.Database, "collection", req.Collection, "index", req.IndexName)

	collection := s.db.Client().Database(req.Database).Collection(req.Collection)
	err := collection.Indexes().DropOne(ctx, req.IndexName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to drop index: %v", err)
	}

	slog.Info("Index dropped", "database", req.Database, "collection", req.Collection, "name", req.IndexName)

	return &emptypb.Empty{}, nil
}
