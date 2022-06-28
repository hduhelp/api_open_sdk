// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: aggregatingapis/infostream/v1/infostream.proto

package infoStreamv1

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

// InfoStreamServiceClient is the client API for InfoStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfoStreamServiceClient interface {
	GetInfoStream(ctx context.Context, in *InfoStreamReq, opts ...grpc.CallOption) (*InfoStreamResp, error)
}

type infoStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoStreamServiceClient(cc grpc.ClientConnInterface) InfoStreamServiceClient {
	return &infoStreamServiceClient{cc}
}

func (c *infoStreamServiceClient) GetInfoStream(ctx context.Context, in *InfoStreamReq, opts ...grpc.CallOption) (*InfoStreamResp, error) {
	out := new(InfoStreamResp)
	err := c.cc.Invoke(ctx, "/aggregatingapis.infoStream.v1.InfoStreamService/GetInfoStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoStreamServiceServer is the server API for InfoStreamService service.
// All implementations should embed UnimplementedInfoStreamServiceServer
// for forward compatibility
type InfoStreamServiceServer interface {
	GetInfoStream(context.Context, *InfoStreamReq) (*InfoStreamResp, error)
}

// UnimplementedInfoStreamServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInfoStreamServiceServer struct {
}

func (UnimplementedInfoStreamServiceServer) GetInfoStream(context.Context, *InfoStreamReq) (*InfoStreamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoStream not implemented")
}

// UnsafeInfoStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoStreamServiceServer will
// result in compilation errors.
type UnsafeInfoStreamServiceServer interface {
	mustEmbedUnimplementedInfoStreamServiceServer()
}

func RegisterInfoStreamServiceServer(s grpc.ServiceRegistrar, srv InfoStreamServiceServer) {
	s.RegisterService(&InfoStreamService_ServiceDesc, srv)
}

func _InfoStreamService_GetInfoStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoStreamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoStreamServiceServer).GetInfoStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregatingapis.infoStream.v1.InfoStreamService/GetInfoStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoStreamServiceServer).GetInfoStream(ctx, req.(*InfoStreamReq))
	}
	return interceptor(ctx, in, info, handler)
}

// InfoStreamService_ServiceDesc is the grpc.ServiceDesc for InfoStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InfoStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aggregatingapis.infoStream.v1.InfoStreamService",
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