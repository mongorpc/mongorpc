// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mongorpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MongoRPCClient is the client API for MongoRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MongoRPCClient interface {
	GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*Value, error)
	InsertDocument(ctx context.Context, in *InsertDocumentRequest, opts ...grpc.CallOption) (*ObjectId, error)
	UpdateDocument(ctx context.Context, in *UpdateDocumentRequest, opts ...grpc.CallOption) (*Value, error)
	DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*Value, error)
	BulkInsertDocuments(ctx context.Context, in *BulkInsertDocumentsRequest, opts ...grpc.CallOption) (*Value, error)
	QueryDocuments(ctx context.Context, in *QueryDocumentsRequest, opts ...grpc.CallOption) (*Value, error)
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (MongoRPC_ListenClient, error)
}

type mongoRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewMongoRPCClient(cc grpc.ClientConnInterface) MongoRPCClient {
	return &mongoRPCClient{cc}
}

func (c *mongoRPCClient) GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/GetDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) InsertDocument(ctx context.Context, in *InsertDocumentRequest, opts ...grpc.CallOption) (*ObjectId, error) {
	out := new(ObjectId)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/InsertDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) UpdateDocument(ctx context.Context, in *UpdateDocumentRequest, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/UpdateDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/DeleteDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) BulkInsertDocuments(ctx context.Context, in *BulkInsertDocumentsRequest, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/BulkInsertDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) QueryDocuments(ctx context.Context, in *QueryDocumentsRequest, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/QueryDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (MongoRPC_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &MongoRPC_ServiceDesc.Streams[0], "/mongorpc.MongoRPC/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &mongoRPCListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MongoRPC_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type mongoRPCListenClient struct {
	grpc.ClientStream
}

func (x *mongoRPCListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MongoRPCServer is the server API for MongoRPC service.
// All implementations must embed UnimplementedMongoRPCServer
// for forward compatibility
type MongoRPCServer interface {
	GetDocument(context.Context, *GetDocumentRequest) (*Value, error)
	InsertDocument(context.Context, *InsertDocumentRequest) (*ObjectId, error)
	UpdateDocument(context.Context, *UpdateDocumentRequest) (*Value, error)
	DeleteDocument(context.Context, *DeleteDocumentRequest) (*Value, error)
	BulkInsertDocuments(context.Context, *BulkInsertDocumentsRequest) (*Value, error)
	QueryDocuments(context.Context, *QueryDocumentsRequest) (*Value, error)
	Listen(*ListenRequest, MongoRPC_ListenServer) error
	mustEmbedUnimplementedMongoRPCServer()
}

// UnimplementedMongoRPCServer must be embedded to have forward compatible implementations.
type UnimplementedMongoRPCServer struct {
}

func (UnimplementedMongoRPCServer) GetDocument(context.Context, *GetDocumentRequest) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDocument not implemented")
}
func (UnimplementedMongoRPCServer) InsertDocument(context.Context, *InsertDocumentRequest) (*ObjectId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertDocument not implemented")
}
func (UnimplementedMongoRPCServer) UpdateDocument(context.Context, *UpdateDocumentRequest) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDocument not implemented")
}
func (UnimplementedMongoRPCServer) DeleteDocument(context.Context, *DeleteDocumentRequest) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDocument not implemented")
}
func (UnimplementedMongoRPCServer) BulkInsertDocuments(context.Context, *BulkInsertDocumentsRequest) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkInsertDocuments not implemented")
}
func (UnimplementedMongoRPCServer) QueryDocuments(context.Context, *QueryDocumentsRequest) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDocuments not implemented")
}
func (UnimplementedMongoRPCServer) Listen(*ListenRequest, MongoRPC_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedMongoRPCServer) mustEmbedUnimplementedMongoRPCServer() {}

// UnsafeMongoRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MongoRPCServer will
// result in compilation errors.
type UnsafeMongoRPCServer interface {
	mustEmbedUnimplementedMongoRPCServer()
}

func RegisterMongoRPCServer(s grpc.ServiceRegistrar, srv MongoRPCServer) {
	s.RegisterService(&MongoRPC_ServiceDesc, srv)
}

func _MongoRPC_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/GetDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).GetDocument(ctx, req.(*GetDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_InsertDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).InsertDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/InsertDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).InsertDocument(ctx, req.(*InsertDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_UpdateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).UpdateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/UpdateDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).UpdateDocument(ctx, req.(*UpdateDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_DeleteDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).DeleteDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/DeleteDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).DeleteDocument(ctx, req.(*DeleteDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_BulkInsertDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkInsertDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).BulkInsertDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/BulkInsertDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).BulkInsertDocuments(ctx, req.(*BulkInsertDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_QueryDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).QueryDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/QueryDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).QueryDocuments(ctx, req.(*QueryDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MongoRPCServer).Listen(m, &mongoRPCListenServer{stream})
}

type MongoRPC_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type mongoRPCListenServer struct {
	grpc.ServerStream
}

func (x *mongoRPCListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MongoRPC_ServiceDesc is the grpc.ServiceDesc for MongoRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MongoRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mongorpc.MongoRPC",
	HandlerType: (*MongoRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDocument",
			Handler:    _MongoRPC_GetDocument_Handler,
		},
		{
			MethodName: "InsertDocument",
			Handler:    _MongoRPC_InsertDocument_Handler,
		},
		{
			MethodName: "UpdateDocument",
			Handler:    _MongoRPC_UpdateDocument_Handler,
		},
		{
			MethodName: "DeleteDocument",
			Handler:    _MongoRPC_DeleteDocument_Handler,
		},
		{
			MethodName: "BulkInsertDocuments",
			Handler:    _MongoRPC_BulkInsertDocuments_Handler,
		},
		{
			MethodName: "QueryDocuments",
			Handler:    _MongoRPC_QueryDocuments_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _MongoRPC_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mongorpc/mongorpc.proto",
}
