package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// Index Direction helper function
func IndexDirection(direction proto.IndexDirection) bsonx.Val {
	if direction == proto.IndexDirection_ASCENDING {
		return bsonx.Int32(1)
	} else {
		return bsonx.Int32(-1)
	}
}

// Encode proto.IndexDirection to int32
func EncodeIndexDirection(direction proto.IndexDirection) int32 {
	if direction == proto.IndexDirection_ASCENDING {
		return 1
	} else {
		return -1
	}
}

// Decode int32 to proto.IndexDirection
func DecodeIndexDirection(direction int32) proto.IndexDirection {
	if direction == 1 {
		return proto.IndexDirection_ASCENDING
	} else {
		return proto.IndexDirection_DESCENDING
	}
}

// Create Index in Collection
func (srv *MongoRPC) CreateIndex(ctx context.Context, in *proto.CreateIndexRequest) (*proto.CreateIndexResponse, error) {

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
			Name:   &in.Index.Name,
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

// List all indexes in collection
func (srv *MongoRPC) ListIndexes(ctx context.Context, in *proto.ListIndexesRequest) (*proto.ListIndexesResponse, error) {

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

// delete index in collection
func (srv *MongoRPC) DeleteIndex(ctx context.Context, in *proto.DeleteIndexRequest) (*proto.DeleteIndexResponse, error) {
	_, err := srv.DB.Database(in.Database).Collection(in.Collection).Indexes().DropOne(ctx, in.Name)
	if err != nil {
		return nil, err
	}

	return &proto.DeleteIndexResponse{
		Name: in.Name,
	}, nil
}
