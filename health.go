package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
)

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
