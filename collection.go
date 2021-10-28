package mongorpc

import (
	"context"
	"encoding/json"

	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (server *MongoRPCServer) ListCollections(ctx context.Context, in *proto.ListCollectionsRequest) (*proto.ListCollectionsResponse, error) {
	logrus.Printf("ListCollections: %v", in)
	cur, err := server.DB.Database(in.Database).ListCollections(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var jsonDocument map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}

		temporaryBytes, err := bson.MarshalExtJSON(result, true, true)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(temporaryBytes, &jsonDocument)
		if err != nil {
			return nil, err
		}
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return &proto.ListCollectionsResponse{
		Fields: parseJSON(jsonDocument),
	}, err
}
