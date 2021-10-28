package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/bson"
)

func (srv *MongoRPCServer) ListCollections(ctx context.Context, in *proto.ListCollectionsRequest) (*proto.ListCollectionsResponse, error) {

	// Get Collections
	collections, err := srv.DB.Database(in.Database).ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var arr []interface{}
	for _, collection := range collections {
		arr = append(arr, collection)
	}

	return &proto.ListCollectionsResponse{
		Collections: &proto.Array{
			Values: EncodeArray(arr).Values,
		},
	}, nil
}
