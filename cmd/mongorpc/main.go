package main

import (
	"fmt"
	"net"
	"os"

	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"github.com/mongorpc/mongorpc"
	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type MongoRPC struct {
	mongoURI        string
	port            int
	jwtSecret       string
	casbinModelPath string
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
			&cli.StringFlag{
				Name:        "casbin-model-path",
				Value:       "rbac_model.conf",
				Usage:       "the path to the casbin model file",
				Destination: &srv.casbinModelPath,
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

	// Initialize a MongoDB adapter and use it in a Casbin enforcer:
	// The adapter will use the database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	adapter, err := mongodbadapter.NewAdapter(srv.mongoURI) // Your MongoDB URL.
	if err != nil {
		return err
	}

	enforcer, err := casbin.NewEnforcer(srv.casbinModelPath, adapter)
	if err != nil {
		return err
	}

	// listen on the port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	// initilize interceptors
	interceptor := mongorpc.Interceptor{
		JWTSecret: srv.jwtSecret,
		Casbin:    enforcer,
	}

	// create a new grpc server
	server := grpc.NewServer(
		// grpc.Creds(tlsCredentials),
		grpc.UnaryInterceptor(interceptor.UnaryInterceptor),
		grpc.StreamInterceptor(interceptor.StreamInterceptor),
	)

	// register the service
	proto.RegisterMongoRPCServer(server, &mongorpc.MongoRPCServer{
		DB: database,
	})

	// Load casbin policy
	err = interceptor.LoadPolicy()
	if err != nil {
		return err
	}

	logrus.Printf("mongorpc server is listening at %v", listener.Addr())

	// start the server
	err = server.Serve(listener)
	return err
}
