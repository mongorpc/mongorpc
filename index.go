package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func IndexDirection(direction proto.IndexDirection) bsonx.Val {
	if direction == proto.IndexDirection_ASCENDING {
		return bsonx.Int32(1)
	} else {
		return bsonx.Int32(-1)
	}
}

func (srv *MongoRPCServer) CreateIndex(ctx context.Context, in *proto.CreateIndexRequest) (*proto.CreateIndexResponse, error) {

	keys := bsonx.Doc{}
	for _, k := range in.Index.Keys {
		keys = append(keys, bsonx.Elem{
			Key:   k.Field,
			Value: IndexDirection(k.Direction),
		})
	}

	index := mongo.IndexModel{
		Keys: keys,
		Options: &options.IndexOptions{
			Unique: &in.Index.Unique,
		},
	}

	name, err := srv.DB.Database(in.Database).Collection(in.Collection).Indexes().CreateOne(context.Background(), index)
	if err != nil {
		return nil, err
	}

	return &proto.CreateIndexResponse{
		Name: name,
	}, nil
}
