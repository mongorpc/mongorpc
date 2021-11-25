package lib

import (
	"context"

	"github.com/mongorpc/mongorpc/lib/decoder"
	"github.com/mongorpc/mongorpc/lib/mongorpc"

	"github.com/mongorpc/mongorpc/lib/encoder"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (srv *MongoRPCServer) GetDocument(ctx context.Context, in *mongorpc.GetDocumentRequest) (*mongorpc.Value, error) {

	// Convert id to primitive.ObjectID
	docID, err := primitive.ObjectIDFromHex(in.DocumentId.Id)
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
	return encoder.Encode(doc), nil
}

func (srv *MongoRPCServer) InsertDocument(ctx context.Context, in *mongorpc.InsertDocumentRequest) (*mongorpc.ObjectId, error) {
	// decode proto document to generic interface
	doc := decoder.Decode(in.Document)

	// Insert document
	result, err := srv.DB.Database(in.Database).Collection(in.Collection).InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	// Return document
	return &mongorpc.ObjectId{
		Id: result.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}

func (srv *MongoRPCServer) BulkInsertDocuments(ctx context.Context, in *mongorpc.BulkInsertDocumentsRequest) (*mongorpc.Value, error) {
	// decode proto document to generic interface
	docs := []interface{}{}
	for _, doc := range in.Documents {
		docs = append(docs, decoder.Decode(doc))
	}

	// Insert document
	result, err := srv.DB.Database(in.Database).Collection(in.Collection).InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}

	arr := &mongorpc.ArrayValue{}
	for _, id := range result.InsertedIDs {
		arr.Values = append(arr.Values, encoder.Encode(id))
	}

	// Return document
	return &mongorpc.Value{
		Type: &mongorpc.Value_ArrayValue{
			ArrayValue: arr,
		},
	}, nil
}

func (srv *MongoRPCServer) UpdateDocument(ctx context.Context, in *mongorpc.UpdateDocumentRequest) (*mongorpc.Value, error) {
	// decode proto document to generic interface
	doc := decoder.Decode(in.Document)

	// Convert id to primitive.ObjectID
	docID, err := primitive.ObjectIDFromHex(in.DocumentId.Id)
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

	type UpdateResult struct {
		MatchedCount  int64 `bson:"matchedCount"`
		ModifiedCount int64 `bson:"modifiedCount"`
		UpsertedCount int64 `bson:"upsertedCount"`
	}

	// Return document
	return encoder.Encode(UpdateResult{
		MatchedCount:  result.MatchedCount,
		ModifiedCount: result.ModifiedCount,
		UpsertedCount: result.UpsertedCount,
	}), nil
}

func (srv *MongoRPCServer) DeleteDocument(ctx context.Context, in *mongorpc.DeleteDocumentRequest) (*mongorpc.Value, error) {
	// Convert id to primitive.ObjectID
	docID, err := primitive.ObjectIDFromHex(in.DocumentId.Id)
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

	// Delete document
	res, err := srv.DB.Database(in.Database).Collection(in.Collection).DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	type DeleteResult struct {
		DeletedCount int64 `bson:"deletedCount"`
	}

	// Return document
	return encoder.Encode(DeleteResult{
		DeletedCount: res.DeletedCount,
	}), nil
}

func (srv *MongoRPCServer) QueryDocuments(ctx context.Context, in *mongorpc.QueryDocumentsRequest) (*mongorpc.Value, error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(in.Limit))
	findOptions.SetSkip(int64(in.Skip))

	findOptions.SetSort(decoder.Decode(in.Sort))

	filter := decoder.Decode(in.Query)

	// Get documents
	results, err := srv.DB.Database(in.Database).Collection(in.Collection).Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	// Decode results to map
	documents := []map[string]interface{}{}
	if err := results.All(ctx, &documents); err != nil {
		return nil, err
	}

	// Return documents
	return encoder.Encode(documents), nil
}
