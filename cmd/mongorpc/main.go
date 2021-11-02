package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"github.com/mongorpc/mongorpc"
	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type MongoRPC struct {
	mongoURI  string
	port      int
	jwtSecret string
}

func main() {
	srv := &MongoRPC{}

	app := &cli.App{
		Name:  "mongorpc",
		Usage: "mongorpc is a gRPC server that can be used to call mongodb directly",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mongodb",
				Value:       "mongodb://localhost:27017",
				Usage:       "the mongodb uri",
				Destination: &srv.mongoURI,
			},
			&cli.IntFlag{
				Name:        "port",
				Value:       27051,
				Usage:       "the port on which the server will listen",
				Destination: &srv.port,
			},
			&cli.StringFlag{
				Name:        "jwt-secret",
				Value:       "secret",
				Usage:       "the secret used to validate the jwt token",
				Destination: &srv.jwtSecret,
			},
		},
		Action: srv.serve,
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

func (srv *MongoRPC) serve(c *cli.Context) error {

	port := fmt.Sprintf(":%d", srv.port)

	// connect to mongodb
	database, err := mongo.Connect(c.Context, options.Client().ApplyURI(srv.mongoURI))
	if err != nil {
		return err
	}

	// listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	// initilize interceptors
	interceptor := mongorpc.Interceptor{
		JWTSecret: srv.jwtSecret,
	}

	// tlsCredentials, err := loadTLSCredentials()
	// if err != nil {
	// 	return err
	// }

	// create a new grpc server
	s := grpc.NewServer(
		// grpc.Creds(tlsCredentials),
		grpc.UnaryInterceptor(interceptor.UnaryInterceptor),
		grpc.StreamInterceptor(interceptor.StreamInterceptor),
	)

	// register the service
	proto.RegisterMongoRPCServer(s, &mongorpc.MongoRPCServer{
		DB: database,
	})

	logrus.Printf("mongorpc server is listening at %v", lis.Addr())

	// start the server
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("tls/ca.pem", "tls/ca.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	return credentials.NewTLS(config), nil
}
