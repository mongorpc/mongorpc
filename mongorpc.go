package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoRPC Server
type MongoRPCServer struct {
	proto.UnimplementedMongoRPCServer
	DB *mongo.Client
}

// Ping MongoDB server
func (srv *MongoRPCServer) Ping(ctx context.Context, in *proto.Empty) (*proto.Empty, error) {
	err := srv.DB.Ping(ctx, nil)
	return in, err
}

// RunCommand executes the given command against the database.
func (srv *MongoRPCServer) RunDatabaseCommand(ctx context.Context, database string, command interface{}) (*mongo.SingleResult, error) {
	result := srv.DB.Database(database).RunCommand(ctx, command)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return result, nil
}
