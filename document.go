package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get document by id from collection in database
func (srv *MongoRPCServer) GetDocument(ctx context.Context, in *proto.GetDocumentRequest) (*proto.GetDocumentResponse, error) {

	// Convert id to primitive.ObjectID
	docID, err := primitive.ObjectIDFromHex(in.DocumentId)
	if err != nil {
		return nil, err
	}

	// filter document by id
	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: docID,
		},
	}

	// Get document
	opts := options.FindOne()
	result := srv.DB.Database(in.Database).Collection(in.Collection).FindOne(ctx, filter, opts)

	if err := result.Err(); err != nil {

		logrus.Println(err)
		return nil, err
	}

	// Decode result to map
	doc := map[string]interface{}{}
	if err := result.Decode(&doc); err != nil {
		return nil, err
	}

	// Return document
	return &proto.GetDocumentResponse{
		Document: &proto.Value{
			Type: &proto.Value_MapValue{
				MapValue: EncodeMap(doc),
			},
		},
	}, nil
}

// List documents from collection in database
func (srv *MongoRPCServer) ListDocuments(ctx context.Context, in *proto.ListDocumentsRequest) (*proto.ListDocumentsResponse, error) {

	findOptions := options.Find()
	findOptions.SetLimit(int64(in.Limit))
	findOptions.SetSkip(int64(in.Skip))

	// add sort to find options
	findOptions.SetSort(DecodeSort(in.Sort))

	// TODO: add filter in findOptions
	filter := DecodeFilter(in.Filter)

	// Get documents
	results, err := srv.DB.Database(in.Database).Collection(in.Collection).Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	// Decode results to map
	arrDocuments := []map[string]interface{}{}
	if err := results.All(ctx, &arrDocuments); err != nil {
		return nil, err
	}

	// convert []map[string]interface{} to proto.arrary
	docs := proto.ArrayValue{}
	for _, doc := range arrDocuments {
		docs.Values = append(docs.Values, &proto.Value{
			Type: &proto.Value_MapValue{
				MapValue: EncodeMap(doc),
			},
		})
	}

	// Return documents
	return &proto.ListDocumentsResponse{
		Documents: &docs,
	}, nil
}

// Create document in collection in database
func (srv *MongoRPCServer) CreateDocument(ctx context.Context, in *proto.CreateDocumentRequest) (*proto.CreateDocumentResponse, error) {

	// decode proto document to generic interface
	doc := DecodeValue(in.Document)

	// Create document
	result, err := srv.DB.Database(in.Database).Collection(in.Collection).InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	// Return document
	return &proto.CreateDocumentResponse{
		DocumentId: result.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}

// Update document in collection in database
func (srv *MongoRPCServer) UpdateDocument(ctx context.Context, in *proto.UpdateDocumentRequest) (*proto.UpdateDocumentResponse, error) {

	// decode proto document to generic interface
	doc := DecodeValue(in.Document)

	// Convert id to primitive.ObjectID
	docID, err := primitive.ObjectIDFromHex(in.DocumentId)
	if err != nil {
		return nil, err
	}

	// filter document by id
	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: docID,
		},
	}

	// options for update
	opts := bson.D{
		primitive.E{
			Key:   "$set",
			Value: doc,
		},
	}

	// Update document
	result, err := srv.DB.Database(in.Database).Collection(in.Collection).UpdateOne(ctx, filter, opts)
	if err != nil {
		logrus.Println(err)
		return nil, err
	}

	// Return document
	return &proto.UpdateDocumentResponse{
		MatchedCount:  result.MatchedCount,
		ModifiedCount: result.ModifiedCount,
		UpsertedCount: result.UpsertedCount,
		// UpsertedId:    &proto.Value{},
	}, nil
}

// Delete document from collection in database
func (srv *MongoRPCServer) DeleteDocument(ctx context.Context, in *proto.DeleteDocumentRequest) (*proto.DeleteDocumentResponse, error) {

	// Delete document
	res, err := srv.DB.Database(in.Database).Collection(in.Collection).DeleteOne(ctx, in.DocumentId)
	if err != nil {
		return nil, err
	}

	// Return document
	return &proto.DeleteDocumentResponse{
		DeletedCount: res.DeletedCount,
	}, nil
}
