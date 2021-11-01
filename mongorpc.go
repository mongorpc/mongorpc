package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoRPC Server
type MongoRPCServer struct {
	proto.UnimplementedMongoRPCServer
	DB *mongo.Client
}

// Ping MongoDB server
func (srv *MongoRPCServer) Ping(ctx context.Context, in *proto.Empty) (*proto.Empty, error) {
	err := srv.DB.Ping(ctx, nil)
	return &proto.Empty{}, err
}

// Health Check MongoDB and rpc server
func (srv *MongoRPCServer) HealthCheck(ctx context.Context, in *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {

	// check mongodb health
	err := srv.DB.Ping(ctx, nil)

	// if mongo is down return not serving response
	if err != nil {
		return &proto.HealthCheckResponse{
			Status: proto.HealthCheckResponse_NOT_SERVING,
		}, nil
	} else {

		// if mongo is up return serving response
		return &proto.HealthCheckResponse{
			Status: proto.HealthCheckResponse_SERVING,
		}, nil
	}
}
