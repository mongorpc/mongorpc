package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ClientInterceptor is a gRPC interceptor that adds the access token to the request
type ClientInterceptor struct {
	accessToken string
}

// UnaryClientInterceptor is a gRPC interceptor that adds the access token to the request
func (interceptor *ClientInterceptor) UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	if interceptor.accessToken != "" {
		md.Set("authorization", interceptor.accessToken)
	}

	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}

// StreamClientInterceptor is a gRPC interceptor that adds the access token to the request
func (interceptor ClientInterceptor) StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	if interceptor.accessToken != "" {
		md.Set("authorization", interceptor.accessToken)
	}

	ctx = metadata.NewOutgoingContext(ctx, md)

	return streamer(ctx, desc, cc, method, opts...)
}
