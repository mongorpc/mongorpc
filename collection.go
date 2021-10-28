package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/bson"
)

// Get List of collections in a database.
func (srv *MongoRPCServer) ListCollections(ctx context.Context, in *proto.ListCollectionsRequest) (*proto.ListCollectionsResponse, error) {

	// Get Collections
	collections, err := srv.DB.Database(in.Database).ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	// Convert []string to []interface{}
	var arr []interface{}
	for _, collection := range collections {
		arr = append(arr, collection)
	}

	// Return Response Object
	return &proto.ListCollectionsResponse{
		Collections: &proto.Array{
			Values: EncodeArray(arr).Values,
		},
	}, nil
}
