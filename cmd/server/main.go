// MongoRPC Server - A gRPC proxy for MongoDB.
package main

import (
	"context"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	mongorpcv1 "github.com/mongorpc/mongorpc/gen/mongorpc/v1"
	mongorpcapi "github.com/mongorpc/mongorpc/internal/api/mongorpc"
	"github.com/mongorpc/mongorpc/internal/config"
	"github.com/mongorpc/mongorpc/internal/repository/mongodb"
)

func main() {
	app := &cli.App{
		Name:  "mongorpc",
		Usage: "A gRPC proxy for MongoDB",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "address",
				Aliases: []string{"a"},
				Usage:   "Server listen address",
				EnvVars: []string{"SERVER_ADDRESS"},
				Value:   "localhost:50051",
			},
			&cli.StringFlag{
				Name:    "mongodb-uri",
				Aliases: []string{"m"},
				Usage:   "MongoDB connection URI",
				EnvVars: []string{"MONGODB_URI"},
				Value:   "mongodb://localhost:27017",
			},
			&cli.StringFlag{
				Name:    "database",
				Aliases: []string{"d"},
				Usage:   "Default MongoDB database",
				EnvVars: []string{"MONGODB_DATABASE"},
				Value:   "mongorpc",
			},
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "Enable debug logging",
				EnvVars: []string{"DEBUG"},
			},
		},
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("Application error", "error", err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	// Configure logging
	logLevel := slog.LevelInfo
	if c.Bool("debug") {
		logLevel = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load configuration
	cfg := config.Load()
	cfg.ServerAddress = c.String("address")
	cfg.MongoDBURI = c.String("mongodb-uri")
	cfg.MongoDBDatabase = c.String("database")

	// Connect to MongoDB
	mongoClient, err := mongodb.New(ctx, cfg.MongoDBURI, cfg.MongoDBDatabase)
	if err != nil {
		return err
	}
	defer func() {
		if err := mongoClient.Close(ctx); err != nil {
			slog.Error("Failed to close MongoDB connection", "error", err)
		}
	}()

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register MongoRPC service
	mongorpcService := mongorpcapi.NewServer(mongoClient)
	mongorpcv1.RegisterMongoRPCServer(grpcServer, mongorpcService)

	// Enable reflection for debugging
	reflection.Register(grpcServer)

	// Start listening
	listener, err := net.Listen("tcp", cfg.ServerAddress)
	if err != nil {
		return err
	}

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		slog.Info("Shutting down server...")
		grpcServer.GracefulStop()
		cancel()
	}()

	slog.Info("MongoRPC server starting", "address", cfg.ServerAddress)
	return grpcServer.Serve(listener)
}
