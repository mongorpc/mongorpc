package client

import (
	"github.com/mongorpc/mongorpc/lib/mongorpc"
	"google.golang.org/grpc"
)

type MongoRPCClient struct {
	address  string
	client   *grpc.ClientConn
	mongorpc mongorpc.MongoRPCClient
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

	m := mongorpc.NewMongoRPCClient(conn)
	c.client = conn
	c.mongorpc = m

	return conn, nil
}
