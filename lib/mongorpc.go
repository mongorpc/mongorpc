package lib

import (
	"github.com/mongorpc/mongorpc/lib/mongorpc"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type MongoRPCServer struct {
	mongorpc.UnimplementedMongoRPCServer
	DB *mongo.Client
}

func NewMongoRPCServer(db *mongo.Client, opt ...grpc.ServerOption) *grpc.Server {
	srv := &MongoRPCServer{DB: db}
	server := grpc.NewServer(opt...)
	mongorpc.RegisterMongoRPCServer(server, srv)
	return server
}

func NewServer(opt ...grpc.ServerOption) *grpc.Server {
	srv := grpc.NewServer(opt...)
	return srv
}
