package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
)

func (srv *MongoRPCServer) ListCollections(ctx context.Context, in *proto.ListCollectionsRequest) (*proto.ListCollectionsResponse, error) {
	return &proto.ListCollectionsResponse{}, nil
}
