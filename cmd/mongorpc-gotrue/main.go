package main

import (
	"fmt"
	"net"
	"os"

	"github.com/mongorpc/mongorpc"
	"github.com/mongorpc/mongorpc/cmd/mongorpc-gotrue/interceptor"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type MongoRPCConfig struct {
	mongoURI  string
	port      int
	jwtSecret string
}

func main() {
	config := &MongoRPCConfig{}

	app := &cli.App{
		Name:  "mongorpc",
		Usage: "mongorpc is a gRPC server that can be used to call mongodb directly",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mongodb",
				Value:       "mongodb://localhost:27017",
				Usage:       "the mongodb uri",
				Destination: &config.mongoURI,
				EnvVars:     []string{"MONGO_URI"},
			},
			&cli.IntFlag{
				Name:        "port",
				Value:       9090,
				Usage:       "the port on which the server will listen",
				Destination: &config.port,
				EnvVars:     []string{"PORT"},
			},
			&cli.StringFlag{
				Name:        "secret",
				Usage:       "the jwt secret",
				Destination: &config.jwtSecret,
				Required:    true,
				EnvVars:     []string{"JWT_SECRET"},
			},
		},
		Action: func(c *cli.Context) error {
			port := fmt.Sprintf("0.0.0.0:%d", config.port)

			// connect to mongodb
			database, err := mongo.Connect(c.Context, options.Client().ApplyURI(config.mongoURI))
			if err != nil {
				return err
			}

			// ping mongodb to check if it's up
			err = database.Ping(c.Context, nil)
			if err != nil {
				return err
			}

			mongorpc := &mongorpc.MongoRPC{
				DB: database,
			}

			i := interceptor.Interceptor{
				JWTSecret: config.jwtSecret,
			}

			srv := mongorpc.NewServer(
				grpc.UnaryInterceptor(i.UnaryInterceptor),
				grpc.StreamInterceptor(i.StreamInterceptor),
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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
