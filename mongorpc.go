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

	UnaryServerInterceptor  grpc.UnaryServerInterceptor
	StreamServerInterceptor grpc.StreamServerInterceptor
}

func (srv *MongoRPC) NewServer() *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(srv.UnaryServerInterceptor),
		grpc.StreamInterceptor(srv.StreamServerInterceptor),
	)

	proto.RegisterMongoRPCServer(server, srv)
	return server
}
