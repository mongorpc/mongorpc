package main

import (
	"context"
	"os"

	"github.com/mongorpc/mongorpc"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	config := &mongorpc.Config{}

	app := &cli.App{
		Name:  "mongorpc",
		Usage: "mongorpc is a gRPC server that can be used to call mongodb directly",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mongodb",
				Value:       "mongodb://localhost:27017",
				Usage:       "the mongodb uri",
				Destination: &config.MongodbURI,
				EnvVars:     []string{"MONGO_URI"},
			},
			&cli.IntFlag{
				Name:        "port",
				Value:       9090,
				Usage:       "the port on which the server will listen",
				Destination: &config.Port,
				EnvVars:     []string{"PORT"},
			},
		},
		Action: func(c *cli.Context) error {
			config.UnaryServerInterceptors = append(config.UnaryServerInterceptors, UnaryInterceptor)
			config.StreamServerInterceptors = append(config.StreamServerInterceptors, StreamInterceptor)

			err := mongorpc.Serve(c.Context, config)
			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logrus.Infoln("method: ", info.FullMethod)
	return handler(ctx, req)
}

func StreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logrus.Infoln("method: ", info.FullMethod)
	return handler(srv, stream)
}
