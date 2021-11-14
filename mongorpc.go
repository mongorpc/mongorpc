package mongorpc

import (
	"context"
	"fmt"
	"net"

	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

// MongoRPC Server
type MongoRPC struct {
	proto.UnimplementedMongoRPCServer
	DB *mongo.Client
}

type Config struct {
	MongodbURI string
	Port       int

	UnaryServerInterceptors  []grpc.UnaryServerInterceptor
	StreamServerInterceptors []grpc.StreamServerInterceptor
}

func (srv *MongoRPC) NewServer(opt ...grpc.ServerOption) *grpc.Server {
	server := grpc.NewServer(opt...)
	proto.RegisterMongoRPCServer(server, srv)
	return server
}

func Serve(ctx context.Context, config *Config) error {
	port := fmt.Sprintf("0.0.0.0:%d", config.Port)

	database, err := ConnectDatabase(ctx, config.MongodbURI)
	if err != nil {
		return err
	}

	mongorpc := &MongoRPC{
		DB: database,
	}

	srv := mongorpc.NewServer(
		grpc.ChainUnaryInterceptor(
			config.UnaryServerInterceptors...,
		),
		grpc.ChainStreamInterceptor(
			config.StreamServerInterceptors...,
		),
	)

	// listen on the port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	logrus.Printf("mongorpc server is listening at %v", listener.Addr())

	// start the server
	err = srv.Serve(listener)
	return err
}
