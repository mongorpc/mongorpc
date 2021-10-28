package main

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/mongorpc/mongorpc"
	pb "github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {

	mongoURI, err := connstring.ParseAndValidate(os.Args[1])
	if err != nil {
		logrus.Fatalf("failed to parse mongodb uri %v", err)
	}

	// connect to mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI.String()))
	if err != nil {
		logrus.Fatalf("failed to Connect: %v", err)
	}

	// start grpc server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(LoggingUnaryInterceptor),
	)
	pb.RegisterMongoRPCServer(s, &mongorpc.MongoRPCServer{
		DB: database,
	})
	logrus.Printf("mongorpc server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func LoggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logrus.Infoln(info.FullMethod)
	return handler(ctx, req)
}
