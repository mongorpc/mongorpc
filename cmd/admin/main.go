package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/mongorpc/mongorpc/lib"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	mongoURI string
	port     int
	debug    bool
}

func main() {
	config := &Config{}

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
			&cli.BoolFlag{
				Name:        "debug",
				Usage:       "enable debug mode",
				Value:       false,
				Destination: &config.debug,
				EnvVars:     []string{"DEBUG"},
			},
		},
		Action: func(ctx *cli.Context) error {
			port := fmt.Sprintf("0.0.0.0:%d", config.port)

			if config.debug {
				logrus.SetLevel(logrus.DebugLevel)
			}

			database, err := ConnectDatabase(ctx.Context, config.mongoURI)
			if err != nil {
				return err
			}

			srv := lib.NewServer(database)

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

func ConnectDatabase(ctx context.Context, mongoURI string) (*mongo.Client, error) {
	// connect to mongodb
	database, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	// ping mongodb to check if it's up
	err = database.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return database, nil
}
