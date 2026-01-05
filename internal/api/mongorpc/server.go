// Package mongorpc provides the gRPC service implementations.
package mongorpc

import (
	"context"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	mongorpcv1 "github.com/mongorpc/mongorpc/gen/mongorpc/v1"
	"github.com/mongorpc/mongorpc/internal/repository/mongodb"
	"github.com/mongorpc/mongorpc/internal/rules"
)

// Server implements the MongoRPC gRPC service.
type Server struct {
	mongorpcv1.UnimplementedMongoRPCServer
	db          *mongodb.Client
	rules       *rules.Engine
	adminKey    string
	adminSecret string
}

// NewServer creates a new MongoRPC server.
func NewServer(db *mongodb.Client, rules *rules.Engine, adminKey, adminSecret string) *Server {
	return &Server{
		db:          db,
		rules:       rules,
		adminKey:    adminKey,
		adminSecret: adminSecret,
	}
}

// isAdmin checks if the request has valid admin credentials.
func (s *Server) isAdmin(ctx context.Context) bool {
	if s.adminKey == "" || s.adminSecret == "" {
		return false // Admin not configured
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}

	keys := md.Get("x-admin-key")
	secrets := md.Get("x-admin-secret")

	if len(keys) == 0 || len(secrets) == 0 {
		return false
	}

	return keys[0] == s.adminKey && secrets[0] == s.adminSecret
}

// GetDocument retrieves a single document by ID.
func (s *Server) GetDocument(ctx context.Context, req *mongorpcv1.GetDocumentRequest) (*mongorpcv1.GetDocumentResponse, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}
	if req.Id == nil || req.Id.Hex == "" {
		return nil, status.Error(codes.InvalidArgument, "document id is required")
	}

	slog.Debug("GetDocument", "database", req.Database, "collection", req.Collection, "id", req.Id.Hex)

	// Parse ObjectID
	objectID, err := bson.ObjectIDFromHex(req.Id.Hex)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid object id: %v", err)
	}

	// Build filter
	filter := bson.D{{Key: "_id", Value: objectID}}

	// Get collection
	coll := s.db.Client().Database(req.Database).Collection(req.Collection)

	// Find document
	var result bson.M
	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &mongorpcv1.GetDocumentResponse{Found: false}, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to get document: %v", err)
	}

	// Convert to proto Document
	doc, err := bsonToDocument(result)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert document: %v", err)
	}

	return &mongorpcv1.GetDocumentResponse{
		Document: doc,
		Found:    true,
	}, nil
}

// CreateDocument creates a new document.
func (s *Server) CreateDocument(ctx context.Context, req *mongorpcv1.CreateDocumentRequest) (*mongorpcv1.CreateDocumentResponse, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}
	if req.Document == nil {
		return nil, status.Error(codes.InvalidArgument, "document is required")
	}

	slog.Debug("CreateDocument", "database", req.Database, "collection", req.Collection)

	// Convert proto Document to BSON
	doc, err := documentToBSON(req.Document)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid document: %v", err)
	}

	// Get collection
	coll := s.db.Client().Database(req.Database).Collection(req.Collection)

	// Insert document
	result, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert document: %v", err)
	}

	// Get inserted ID
	insertedID, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, status.Error(codes.Internal, "unexpected inserted id type")
	}

	// Return created document with ID
	if req.Document.Fields == nil {
		req.Document.Fields = make(map[string]*mongorpcv1.Value)
	}
	req.Document.Fields["_id"] = &mongorpcv1.Value{
		ValueType: &mongorpcv1.Value_ObjectIdValue{
			ObjectIdValue: &mongorpcv1.ObjectId{Hex: insertedID.Hex()},
		},
	}
	req.Document.Id = &mongorpcv1.ObjectId{Hex: insertedID.Hex()}

	return &mongorpcv1.CreateDocumentResponse{
		Document: req.Document,
		WriteResult: &mongorpcv1.WriteResult{
			InsertedCount: 1,
		},
	}, nil
}

// DeleteDocument deletes a document by ID.
func (s *Server) DeleteDocument(ctx context.Context, req *mongorpcv1.DeleteDocumentRequest) (*mongorpcv1.DeleteDocumentResponse, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}
	if req.Id == nil || req.Id.Hex == "" {
		return nil, status.Error(codes.InvalidArgument, "document id is required")
	}

	slog.Debug("DeleteDocument", "database", req.Database, "collection", req.Collection, "id", req.Id.Hex)

	// Parse ObjectID
	objectID, err := bson.ObjectIDFromHex(req.Id.Hex)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid object id: %v", err)
	}

	// Build filter
	filter := bson.D{{Key: "_id", Value: objectID}}

	// Get collection
	coll := s.db.Client().Database(req.Database).Collection(req.Collection)

	// Delete document
	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete document: %v", err)
	}

	return &mongorpcv1.DeleteDocumentResponse{
		WriteResult: &mongorpcv1.WriteResult{
			DeletedCount: result.DeletedCount,
		},
	}, nil
}

// ListDocuments lists documents in a collection.
func (s *Server) ListDocuments(ctx context.Context, req *mongorpcv1.ListDocumentsRequest) (*mongorpcv1.ListDocumentsResponse, error) {
	if req.Database == "" || req.Collection == "" {
		return nil, status.Error(codes.InvalidArgument, "database and collection are required")
	}

	slog.Debug("ListDocuments", "database", req.Database, "collection", req.Collection)

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

	// Build find options
	findOpts := options.Find()
	if req.PageSize > 0 {
		findOpts.SetLimit(int64(req.PageSize))
	}

	// Execute query
	cursor, err := coll.Find(ctx, filter, findOpts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to query documents: %v", err)
	}
	defer cursor.Close(ctx)

	// Collect results
	var documents []*mongorpcv1.Document
	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to decode document: %v", err)
		}

		doc, err := bsonToDocument(result)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to convert document: %v", err)
		}
		documents = append(documents, doc)
	}

	if err := cursor.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "cursor error: %v", err)
	}

	return &mongorpcv1.ListDocumentsResponse{
		Documents: documents,
	}, nil
}

// bsonToDocument converts a BSON map to a proto Document.
func bsonToDocument(m bson.M) (*mongorpcv1.Document, error) {
	doc := &mongorpcv1.Document{
		Fields: make(map[string]*mongorpcv1.Value),
	}

	for key, val := range m {
		protoVal, err := bsonToValue(val)
		if err != nil {
			return nil, fmt.Errorf("failed to convert field %s: %w", key, err)
		}
		doc.Fields[key] = protoVal

		if key == "_id" {
			if v, ok := protoVal.ValueType.(*mongorpcv1.Value_ObjectIdValue); ok {
				doc.Id = v.ObjectIdValue
			}
		}
	}

	return doc, nil
}

// bsonToValue converts a BSON value to a proto Value.
func bsonToValue(v interface{}) (*mongorpcv1.Value, error) {
	if v == nil {
		return &mongorpcv1.Value{
			ValueType: &mongorpcv1.Value_NullValue{},
		}, nil
	}

	switch val := v.(type) {
	case bson.ObjectID:
		return &mongorpcv1.Value{
			ValueType: &mongorpcv1.Value_ObjectIdValue{
				ObjectIdValue: &mongorpcv1.ObjectId{Hex: val.Hex()},
			},
		}, nil
	case string:
		return &mongorpcv1.Value{
			ValueType: &mongorpcv1.Value_StringValue{StringValue: val},
		}, nil
	case int32:
		return &mongorpcv1.Value{
			ValueType: &mongorpcv1.Value_Int32Value{Int32Value: val},
		}, nil
	case int64:
		return &mongorpcv1.Value{
			ValueType: &mongorpcv1.Value_Int64Value{Int64Value: val},
		}, nil
	case int:
		return &mongorpcv1.Value{
			ValueType: &mongorpcv1.Value_Int64Value{Int64Value: int64(val)},
		}, nil
	case float64:
		return &mongorpcv1.Value{
			ValueType: &mongorpcv1.Value_DoubleValue{DoubleValue: val},
		}, nil
	case bool:
		return &mongorpcv1.Value{
			ValueType: &mongorpcv1.Value_BooleanValue{BooleanValue: val},
		}, nil
	case bson.M:
		mapVal := &mongorpcv1.MapValue{
			Fields: make(map[string]*mongorpcv1.Value),
		}
		for k, v := range val {
			fieldVal, err := bsonToValue(v)
			if err != nil {
				return nil, err
			}
			mapVal.Fields[k] = fieldVal
		}
		return &mongorpcv1.Value{
			ValueType: &mongorpcv1.Value_MapValue{MapValue: mapVal},
		}, nil
	case bson.A:
		arrayVal := &mongorpcv1.ArrayValue{
			Values: make([]*mongorpcv1.Value, 0, len(val)),
		}
		for _, item := range val {
			itemVal, err := bsonToValue(item)
			if err != nil {
				return nil, err
			}
			arrayVal.Values = append(arrayVal.Values, itemVal)
		}
		return &mongorpcv1.Value{
			ValueType: &mongorpcv1.Value_ArrayValue{ArrayValue: arrayVal},
		}, nil
	default:
		// For unknown types, convert to string representation
		return &mongorpcv1.Value{
			ValueType: &mongorpcv1.Value_StringValue{StringValue: fmt.Sprintf("%v", val)},
		}, nil
	}
}

// documentToBSON converts a proto Document to a BSON document.
func documentToBSON(doc *mongorpcv1.Document) (bson.D, error) {
	if doc == nil || doc.Fields == nil {
		return bson.D{}, nil
	}

	result := bson.D{}
	for key, val := range doc.Fields {
		bsonVal, err := valueToBSON(val)
		if err != nil {
			return nil, fmt.Errorf("failed to convert field %s: %w", key, err)
		}
		result = append(result, bson.E{Key: key, Value: bsonVal})
	}

	return result, nil
}

// valueToBSON converts a proto Value to a BSON value.
func valueToBSON(val *mongorpcv1.Value) (interface{}, error) {
	if val == nil {
		return nil, nil
	}

	switch v := val.ValueType.(type) {
	case *mongorpcv1.Value_NullValue:
		return nil, nil
	case *mongorpcv1.Value_BooleanValue:
		return v.BooleanValue, nil
	case *mongorpcv1.Value_Int32Value:
		return v.Int32Value, nil
	case *mongorpcv1.Value_Int64Value:
		return v.Int64Value, nil
	case *mongorpcv1.Value_DoubleValue:
		return v.DoubleValue, nil
	case *mongorpcv1.Value_StringValue:
		return v.StringValue, nil
	case *mongorpcv1.Value_BytesValue:
		return v.BytesValue, nil
	case *mongorpcv1.Value_ObjectIdValue:
		if v.ObjectIdValue == nil {
			return nil, nil
		}
		return bson.ObjectIDFromHex(v.ObjectIdValue.Hex)
	case *mongorpcv1.Value_MapValue:
		if v.MapValue == nil {
			return bson.M{}, nil
		}
		result := bson.M{}
		for k, fieldVal := range v.MapValue.Fields {
			bsonVal, err := valueToBSON(fieldVal)
			if err != nil {
				return nil, err
			}
			result[k] = bsonVal
		}
		return result, nil
	case *mongorpcv1.Value_ArrayValue:
		if v.ArrayValue == nil {
			return bson.A{}, nil
		}
		result := make(bson.A, 0, len(v.ArrayValue.Values))
		for _, itemVal := range v.ArrayValue.Values {
			bsonVal, err := valueToBSON(itemVal)
			if err != nil {
				return nil, err
			}
			result = append(result, bsonVal)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported value type: %T", val.ValueType)
	}
}

// filterToBSON converts a proto Filter to BSON filter.
func filterToBSON(filter *mongorpcv1.Filter) (bson.D, error) {
	if filter == nil {
		return bson.D{}, nil
	}

	// Handle Raw filter (MapValue)
	if raw := filter.GetRaw(); raw != nil {
		result := bson.D{}
		for k, v := range raw.Fields {
			bsonVal, err := valueToBSON(v)
			if err != nil {
				return nil, fmt.Errorf("failed to convert filter field %s: %w", k, err)
			}
			result = append(result, bson.E{Key: k, Value: bsonVal})
		}
		return result, nil
	}

	return bson.D{}, nil
}
