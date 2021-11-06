package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
)

// Ping MongoDB server
func (srv *MongoRPC) Ping(ctx context.Context, in *proto.Empty) (*proto.Empty, error) {
	err := srv.DB.Ping(ctx, nil)
	return in, err
}
