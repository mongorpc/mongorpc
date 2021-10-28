package main

import (
	"context"
	"net"
	"time"

	pb "github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type MongoRPCServer struct {
	pb.UnimplementedMongoRPCServer
	client *mongo.Client
}

func main() {
	// connect to mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:2707"))
	if err != nil {
		logrus.Fatalf("failed to Connect: %v", err)
	}

	// start grpc server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMongoRPCServer(s, &MongoRPCServer{
		client: client,
	})
	logrus.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
