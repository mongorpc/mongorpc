package client

import (
	"github.com/mongorpc/mongorpc/lib/mongorpc"
	"google.golang.org/grpc"
)

type MongoRPCClient struct {
	address  string
	client   *grpc.ClientConn
	mongorpc mongorpc.MongoRPCClient
	admin    mongorpc.MongoRPCAdminClient
}

func NewClient(address string) *MongoRPCClient {
	return &MongoRPCClient{
		address: address,
	}
}

func (c *MongoRPCClient) Connect(opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		c.address,
		opts...,
	)
	if err != nil {
		return nil, err
	}

	c.mongorpc = mongorpc.NewMongoRPCClient(conn)
	c.admin = mongorpc.NewMongoRPCAdminClient(conn)
	c.client = conn

	return conn, nil
}
