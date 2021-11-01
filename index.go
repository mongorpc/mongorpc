package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/bson"
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

func EncodeIndexDirection(direction proto.IndexDirection) int32 {
	if direction == proto.IndexDirection_ASCENDING {
		return 1
	} else {
		return -1
	}
}

func DecodeIndexDirection(direction int32) proto.IndexDirection {
	if direction == 1 {
		return proto.IndexDirection_ASCENDING
	} else {
		return proto.IndexDirection_DESCENDING
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

	name, err := srv.DB.Database(in.Database).Collection(in.Collection).Indexes().CreateOne(ctx, index)
	if err != nil {
		return nil, err
	}

	return &proto.CreateIndexResponse{
		Name: name,
	}, nil
}

func (srv *MongoRPCServer) ListIndexes(ctx context.Context, in *proto.ListIndexesRequest) (*proto.ListIndexesResponse, error) {

	cursor, err := srv.DB.Database(in.Database).Collection(in.Collection).Indexes().List(ctx)
	if err != nil {
		return nil, err
	}

	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	var indexes []*proto.Index

	for _, i := range result {
		index := &proto.Index{
			Name: i["name"].(string),
		}

		key := i["key"].(bson.M)
		for k, v := range key {
			index.Keys = append(index.Keys, &proto.IndexKey{
				Field:     k,
				Direction: DecodeIndexDirection(v.(int32)),
			})
		}

		indexes = append(indexes, index)
	}

	return &proto.ListIndexesResponse{
		Indexes: indexes,
	}, nil
}
