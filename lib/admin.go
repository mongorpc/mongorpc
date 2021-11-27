package lib

import (
	"context"

	"github.com/mongorpc/mongorpc/lib/encoder"
	"github.com/mongorpc/mongorpc/lib/mongorpc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func (srv *MongoRPCAdminServer) ListDatabases(ctx context.Context, in *mongorpc.Empty) (*mongorpc.Value, error) {
	filter := bson.D{}
	databases, err := srv.DB.ListDatabaseNames(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	return encoder.Encode(databases), nil
}

func (srv *MongoRPCAdminServer) DropDatabase(ctx context.Context, in *mongorpc.DropDatabaseRequest) (*mongorpc.Empty, error) {
	err := srv.DB.Database(in.Database).Drop(ctx)
	if err != nil {
		return nil, err
	}
	return &mongorpc.Empty{}, nil
}

func (srv *MongoRPCAdminServer) CreateCollection(ctx context.Context, in *mongorpc.CreateCollectionRequest) (*mongorpc.Empty, error) {
	options := options.CreateCollectionOptions{}
	if in.Options != nil {
		options.MaxDocuments = &in.Options.Max
		options.SizeInBytes = &in.Options.Size
		// TODO: other options
	}

	err := srv.DB.Database(in.Database).CreateCollection(ctx, in.Collection)
	if err != nil {
		return nil, err
	}
	return &mongorpc.Empty{}, nil
}

func (srv *MongoRPCAdminServer) DropCollection(ctx context.Context, in *mongorpc.DropCollectionRequest) (*mongorpc.Empty, error) {
	err := srv.DB.Database(in.Database).Collection(in.Collection).Drop(ctx)
	if err != nil {
		return nil, err
	}
	return &mongorpc.Empty{}, nil
}

func (srv *MongoRPCAdminServer) ListCollections(ctx context.Context, in *mongorpc.ListCollectionsRequest) (*mongorpc.Value, error) {
	filter := bson.D{}
	collections, err := srv.DB.Database(in.Database).ListCollectionNames(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	return encoder.Encode(collections), nil
}

// Create Index in Collection
func (srv *MongoRPCAdminServer) CreateIndex(ctx context.Context, in *mongorpc.CreateIndexRequest) (*mongorpc.Value, error) {

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

	return encoder.Encode(name), nil
}

// List all indexes in collection
func (srv *MongoRPCAdminServer) ListIndexes(ctx context.Context, in *mongorpc.ListIndexesRequest) (*mongorpc.Value, error) {

	cursor, err := srv.DB.Database(in.Database).Collection(in.Collection).Indexes().List(ctx)
	if err != nil {
		return nil, err
	}

	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	var indexes []*mongorpc.Index

	for _, i := range result {
		index := &mongorpc.Index{
			Name: i["name"].(string),
		}

		key := i["key"].(bson.M)
		for k, v := range key {
			index.Keys = append(index.Keys, &mongorpc.IndexKey{
				Field:     k,
				Direction: DecodeIndexDirection(v.(int32)),
			})
		}

		indexes = append(indexes, index)
	}

	return encoder.Encode(indexes), nil
}

// delete index in collection
func (srv *MongoRPCAdminServer) DropIndex(ctx context.Context, in *mongorpc.DropIndexRequest) (*mongorpc.Empty, error) {
	_, err := srv.DB.Database(in.Database).Collection(in.Collection).Indexes().DropOne(ctx, in.Index)
	if err != nil {
		return nil, err
	}
	return &mongorpc.Empty{}, nil
}

func IndexDirection(direction mongorpc.IndexDirection) bsonx.Val {
	if direction == mongorpc.IndexDirection_ASCENDING {
		return bsonx.Int32(1)
	} else {
		return bsonx.Int32(-1)
	}
}

func DecodeIndexDirection(direction int32) mongorpc.IndexDirection {
	if direction == 1 {
		return mongorpc.IndexDirection_ASCENDING
	} else {
		return mongorpc.IndexDirection_DESCENDING
	}
}
