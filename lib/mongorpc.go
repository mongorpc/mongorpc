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

type MongoRPCAdminServer struct {
	mongorpc.UnimplementedMongoRPCAdminServer
	DB *mongo.Client
}

// NewMongoRPCServer creates a new MongoRPCServer
func NewMongoRPCServer(db *mongo.Client) *MongoRPCServer {
	srv := &MongoRPCServer{DB: db}
	return srv
}

// NewMongoRPCAdminServer creates a new MongoRPCAdminServer
func NewMongoRPCAdminServer(db *mongo.Client) *MongoRPCAdminServer {
	srv := &MongoRPCAdminServer{DB: db}
	return srv
}

// New gRPC Server
func NewGRPCServer(opt ...grpc.ServerOption) *grpc.Server {
	srv := grpc.NewServer(opt...)
	return srv
}

// New MongoRPC and MongoRPCAdmin Server
func NewServer(db *mongo.Client, opt ...grpc.ServerOption) *grpc.Server {
	srv := NewGRPCServer(opt...)
	mongorpc.RegisterMongoRPCServer(srv, NewMongoRPCServer(db))
	mongorpc.RegisterMongoRPCAdminServer(srv, NewMongoRPCAdminServer(db))
	return srv
}
