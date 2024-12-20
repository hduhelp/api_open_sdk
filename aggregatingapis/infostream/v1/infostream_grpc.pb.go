// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: aggregatingapis/infostream/v1/infostream.proto

package infostreamv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	InfoStreamService_GetInfoStream_FullMethodName = "/aggregatingapis.infostream.v1.InfoStreamService/GetInfoStream"
)

// InfoStreamServiceClient is the client API for InfoStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfoStreamServiceClient interface {
	GetInfoStream(ctx context.Context, in *GetInfoStreamRequest, opts ...grpc.CallOption) (*GetInfoStreamResponse, error)
}

type infoStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoStreamServiceClient(cc grpc.ClientConnInterface) InfoStreamServiceClient {
	return &infoStreamServiceClient{cc}
}

func (c *infoStreamServiceClient) GetInfoStream(ctx context.Context, in *GetInfoStreamRequest, opts ...grpc.CallOption) (*GetInfoStreamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInfoStreamResponse)
	err := c.cc.Invoke(ctx, InfoStreamService_GetInfoStream_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoStreamServiceServer is the server API for InfoStreamService service.
// All implementations must embed UnimplementedInfoStreamServiceServer
// for forward compatibility.
type InfoStreamServiceServer interface {
	GetInfoStream(context.Context, *GetInfoStreamRequest) (*GetInfoStreamResponse, error)
	mustEmbedUnimplementedInfoStreamServiceServer()
}

// UnimplementedInfoStreamServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInfoStreamServiceServer struct{}

func (UnimplementedInfoStreamServiceServer) GetInfoStream(context.Context, *GetInfoStreamRequest) (*GetInfoStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoStream not implemented")
}
func (UnimplementedInfoStreamServiceServer) mustEmbedUnimplementedInfoStreamServiceServer() {}
func (UnimplementedInfoStreamServiceServer) testEmbeddedByValue()                           {}

// UnsafeInfoStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoStreamServiceServer will
// result in compilation errors.
type UnsafeInfoStreamServiceServer interface {
	mustEmbedUnimplementedInfoStreamServiceServer()
}

func RegisterInfoStreamServiceServer(s grpc.ServiceRegistrar, srv InfoStreamServiceServer) {
	// If the following call pancis, it indicates UnimplementedInfoStreamServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InfoStreamService_ServiceDesc, srv)
}

func _InfoStreamService_GetInfoStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoStreamServiceServer).GetInfoStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InfoStreamService_GetInfoStream_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoStreamServiceServer).GetInfoStream(ctx, req.(*GetInfoStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InfoStreamService_ServiceDesc is the grpc.ServiceDesc for InfoStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InfoStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aggregatingapis.infostream.v1.InfoStreamService",
	HandlerType: (*InfoStreamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfoStream",
			Handler:    _InfoStreamService_GetInfoStream_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aggregatingapis/infostream/v1/infostream.proto",
}
