// Package mongorpc provides the gRPC service implementations.
package mongorpc

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	mongorpcv1 "github.com/mongorpc/mongorpc/gen/mongorpc/v1"
)

// ListCollections lists all collections in a database.
func (s *Server) ListCollections(ctx context.Context, req *mongorpcv1.ListCollectionsRequest) (*mongorpcv1.ListCollectionsResponse, error) {
	if req.Database == "" {
		return nil, status.Error(codes.InvalidArgument, "database is required")
	}

	// Admin-only operation
	if !s.isAdmin(ctx) {
		return nil, status.Error(codes.PermissionDenied, "admin privileges required for collection management")
	}

	slog.Debug("ListCollections", "database", req.Database)

	db := s.db.Client().Database(req.Database)

	// Build filter
	filter := bson.M{}
	if req.NameFilter != "" {
		filter["name"] = bson.M{"$regex": req.NameFilter}
	}

	cursor, err := db.ListCollections(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list collections: %v", err)
	}
	defer cursor.Close(ctx)

	var collections []*mongorpcv1.CollectionInfo
	for cursor.Next(ctx) {
		var collDoc bson.M
		if err := cursor.Decode(&collDoc); err != nil {
			continue
		}

		info := &mongorpcv1.CollectionInfo{
			Name: collDoc["name"].(string),
		}

		if collType, ok := collDoc["type"].(string); ok {
			info.Type = collType
		}

		// Skip system collections unless explicitly requested
		if !req.IncludeSystem && len(info.Name) > 0 && info.Name[0] == '_' {
			continue
		}

		collections = append(collections, info)
	}

	return &mongorpcv1.ListCollectionsResponse{
		Collections: collections,
	}, nil
}

// CreateCollection creates a new collection.
func (s *Server) CreateCollection(ctx context.Context, req *mongorpcv1.CreateCollectionRequest) (*emptypb.Empty, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}

	// Admin-only operation
	if !s.isAdmin(ctx) {
		return nil, status.Error(codes.PermissionDenied, "admin privileges required for collection management")
	}

	slog.Debug("CreateCollection", "database", req.Database, "collection", req.Collection)

	db := s.db.Client().Database(req.Database)

	// Build options
	opts := options.CreateCollection()
	if req.Options != nil {
		if req.Options.Capped {
			opts.SetCapped(true)
			if req.Options.Size > 0 {
				opts.SetSizeInBytes(req.Options.Size)
			}
			if req.Options.Max > 0 {
				opts.SetMaxDocuments(req.Options.Max)
			}
		}
		if req.Options.ValidationLevel != "" {
			opts.SetValidationLevel(req.Options.ValidationLevel)
		}
		if req.Options.ValidationAction != "" {
			opts.SetValidationAction(req.Options.ValidationAction)
		}
		if req.Options.ExpireAfterSeconds > 0 {
			opts.SetExpireAfterSeconds(req.Options.ExpireAfterSeconds)
		}
	}

	err := db.CreateCollection(ctx, req.Collection, opts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create collection: %v", err)
	}

	slog.Info("Collection created", "database", req.Database, "collection", req.Collection)

	return &emptypb.Empty{}, nil
}

// DropCollection drops a collection.
func (s *Server) DropCollection(ctx context.Context, req *mongorpcv1.DropCollectionRequest) (*emptypb.Empty, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}

	// Admin-only operation
	if !s.isAdmin(ctx) {
		return nil, status.Error(codes.PermissionDenied, "admin privileges required for collection management")
	}

	slog.Debug("DropCollection", "database", req.Database, "collection", req.Collection)

	collection := s.db.Client().Database(req.Database).Collection(req.Collection)
	err := collection.Drop(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to drop collection: %v", err)
	}

	slog.Info("Collection dropped", "database", req.Database, "collection", req.Collection)

	return &emptypb.Empty{}, nil
}

// RenameCollection renames a collection.
func (s *Server) RenameCollection(ctx context.Context, req *mongorpcv1.RenameCollectionRequest) (*emptypb.Empty, error) {
	if req.Database == "" || req.Collection == "" || req.NewName == "" {
		return nil, status.Error(codes.InvalidArgument, "database, collection, and new_name are required")
	}

	// Admin-only operation
	if !s.isAdmin(ctx) {
		return nil, status.Error(codes.PermissionDenied, "admin privileges required for collection management")
	}

	slog.Debug("RenameCollection", "database", req.Database, "from", req.Collection, "to", req.NewName)

	// Use the admin command to rename
	adminDB := s.db.Client().Database("admin")

	renameCmd := bson.D{
		{Key: "renameCollection", Value: req.Database + "." + req.Collection},
		{Key: "to", Value: req.Database + "." + req.NewName},
		{Key: "dropTarget", Value: req.DropTarget},
	}

	var result bson.M
	err := adminDB.RunCommand(ctx, renameCmd).Decode(&result)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to rename collection: %v", err)
	}

	slog.Info("Collection renamed", "database", req.Database, "from", req.Collection, "to", req.NewName)

	return &emptypb.Empty{}, nil
}
