package lib

import (
	"context"

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
