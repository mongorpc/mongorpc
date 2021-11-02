package main

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/mongorpc/mongorpc"
	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
	"google.golang.org/grpc"
)

const (
	// port on which the server will listen.
	port = ":50051"
)

func main() {

	// parse mongodb uri
	mongoURI, err := connstring.ParseAndValidate(os.Args[1])
	if err != nil {
		logrus.Fatalf("failed to parse mongodb uri %v", err)
	}

	// create a new mongodb client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect to mongodb
	database, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI.String()))
	if err != nil {
		logrus.Fatalf("failed to Connect: %v", err)
	}

	// listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	// create a new grpc server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(LoggingUnaryInterceptor),
	)

	// register the service
	proto.RegisterMongoRPCServer(s, &mongorpc.MongoRPCServer{
		DB: database,
	})

	logrus.Printf("mongorpc server is listening at %v", lis.Addr())

	// start the server
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

// LoggingUnaryInterceptor is a unary interceptor that logs the request and response.
func LoggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logrus.Infoln(info.FullMethod)
	return handler(ctx, req)
}
