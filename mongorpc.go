package mongorpc

import (
	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

// MongoRPC Server
type MongoRPC struct {
	proto.UnimplementedMongoRPCServer
	DB *mongo.Client
}

func (srv *MongoRPC) NewServer(opt ...grpc.ServerOption) *grpc.Server {
	server := grpc.NewServer(opt...)
	proto.RegisterMongoRPCServer(server, srv)
	return server
}
