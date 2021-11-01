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
