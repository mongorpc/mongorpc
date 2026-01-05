// MongoRPC Server - A gRPC proxy for MongoDB.
package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	mongorpcv1 "github.com/mongorpc/mongorpc/gen/mongorpc/v1"
	mongorpcapi "github.com/mongorpc/mongorpc/internal/api/mongorpc"
	"github.com/mongorpc/mongorpc/internal/config"
	"github.com/mongorpc/mongorpc/internal/middleware"
	"github.com/mongorpc/mongorpc/internal/observability"
	"github.com/mongorpc/mongorpc/internal/repository/mongodb"
	"github.com/mongorpc/mongorpc/internal/rules"
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
			&cli.StringFlag{
				Name:    "metrics-address",
				Usage:   "Metrics server address",
				EnvVars: []string{"METRICS_ADDRESS"},
				Value:   ":9090",
			},
			&cli.Float64Flag{
				Name:    "rate-limit",
				Usage:   "Requests per second rate limit",
				EnvVars: []string{"RATE_LIMIT"},
				Value:   100,
			},
			&cli.IntFlag{
				Name:    "rate-burst",
				Usage:   "Rate limit burst size",
				EnvVars: []string{"RATE_BURST"},
				Value:   200,
			},
			&cli.StringSliceFlag{
				Name:    "api-keys",
				Usage:   "Valid API keys for authentication",
				EnvVars: []string{"API_KEYS"},
			},
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "Enable debug logging",
				EnvVars: []string{"DEBUG"},
			},
			&cli.BoolFlag{
				Name:    "enable-tracing",
				Usage:   "Enable OpenTelemetry tracing",
				EnvVars: []string{"ENABLE_TRACING"},
			},
			&cli.StringFlag{
				Name:    "admin-key",
				Usage:   "Admin API key for privileged operations",
				EnvVars: []string{"MONGORPC_ADMIN_KEY"},
			},
			&cli.StringFlag{
				Name:    "admin-secret",
				Usage:   "Admin API secret for privileged operations",
				EnvVars: []string{"MONGORPC_ADMIN_SECRET"},
			},
			&cli.StringFlag{
				Name:    "rules",
				Usage:   "Path to security rules YAML file",
				EnvVars: []string{"MONGORPC_RULES"},
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

	// Initialize tracing if enabled
	if c.Bool("enable-tracing") {
		shutdown, err := observability.InitTracer("mongorpc", "1.0.0")
		if err != nil {
			slog.Warn("Failed to initialize tracing", "error", err)
		} else {
			defer shutdown(ctx)
			slog.Info("OpenTelemetry tracing enabled")
		}
	}

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

	// Build middleware chain
	unaryChain := []grpc.UnaryServerInterceptor{
		middleware.RecoveryInterceptor(),
		middleware.LoggingInterceptor(),
	}
	streamChain := []grpc.StreamServerInterceptor{
		middleware.StreamRecoveryInterceptor(),
		middleware.StreamLoggingInterceptor(),
	}

	// Add rate limiting
	rateLimiter := middleware.NewRateLimiter(c.Float64("rate-limit"), c.Int("rate-burst"))
	unaryChain = append(unaryChain, middleware.RateLimitInterceptor(rateLimiter))
	streamChain = append(streamChain, middleware.StreamRateLimitInterceptor(rateLimiter))

	// Add authentication if API keys provided
	apiKeys := c.StringSlice("api-keys")
	if len(apiKeys) > 0 {
		authConfig := &middleware.AuthConfig{
			APIKeys:     apiKeys,
			SkipMethods: []string{}, // Add methods to skip auth
		}
		unaryChain = append(unaryChain, middleware.AuthInterceptor(authConfig))
		streamChain = append(streamChain, middleware.StreamAuthInterceptor(authConfig))
		slog.Info("Authentication enabled", "num_keys", len(apiKeys))
	}

	// Add validation
	unaryChain = append(unaryChain, middleware.ValidationInterceptor())

	// Add rules engine
	var rulesEngine *rules.Engine

	rulesPath := c.String("rules")
	if rulesPath != "" {
		slog.Info("Loading rules from file", "path", rulesPath)
		rulesEngine, err = rules.LoadFromFile(rulesPath)
	} else {
		rulesEngine, err = rules.NewEngine()
		if err == nil {
			// Default to allow all if no rules file provided
			rulesEngine.SetDefaultAllow(rules.OpRead, true)
			rulesEngine.SetDefaultAllow(rules.OpList, true)
			rulesEngine.SetDefaultAllow(rules.OpCreate, true)
			rulesEngine.SetDefaultAllow(rules.OpUpdate, true)
			rulesEngine.SetDefaultAllow(rules.OpDelete, true)
		}
	}

	if err != nil {
		slog.Warn("Failed to create rules engine", "error", err)
	} else {
		unaryChain = append(unaryChain, rules.RulesInterceptor(rulesEngine))
		streamChain = append(streamChain, rules.StreamRulesInterceptor(rulesEngine))
	}

	// Create gRPC server with middleware
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(unaryChain...),
		grpc.ChainStreamInterceptor(streamChain...),
	)

	// Register MongoRPC service
	adminKey := c.String("admin-key")
	adminSecret := c.String("admin-secret")
	if adminKey != "" && adminSecret != "" {
		slog.Info("Admin API enabled")
	}
	mongorpcService := mongorpcapi.NewServer(mongoClient, rulesEngine, adminKey, adminSecret)
	mongorpcv1.RegisterMongoRPCServer(grpcServer, mongorpcService)

	// Enable reflection for debugging
	reflection.Register(grpcServer)

	// Start metrics server
	metricsAddr := c.String("metrics-address")
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
		slog.Info("Metrics server starting", "address", metricsAddr)
		if err := http.ListenAndServe(metricsAddr, nil); err != nil {
			slog.Error("Metrics server error", "error", err)
		}
	}()

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

	slog.Info("MongoRPC server starting", "address", cfg.ServerAddress, "metrics", metricsAddr)
	return grpcServer.Serve(listener)
}
