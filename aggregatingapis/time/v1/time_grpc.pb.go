// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: aggregatingapis/time/v1/time.proto

package timev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TimeServiceClient is the client API for TimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeServiceClient interface {
	GetTime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTimeResponse, error)
}

type timeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeServiceClient(cc grpc.ClientConnInterface) TimeServiceClient {
	return &timeServiceClient{cc}
}

func (c *timeServiceClient) GetTime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTimeResponse, error) {
	out := new(GetTimeResponse)
	err := c.cc.Invoke(ctx, "/aggregatingapis.time.v1.TimeService/GetTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeServiceServer is the server API for TimeService service.
// All implementations must embed UnimplementedTimeServiceServer
// for forward compatibility
type TimeServiceServer interface {
	GetTime(context.Context, *emptypb.Empty) (*GetTimeResponse, error)
	mustEmbedUnimplementedTimeServiceServer()
}

// UnimplementedTimeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimeServiceServer struct {
}

func (UnimplementedTimeServiceServer) GetTime(context.Context, *emptypb.Empty) (*GetTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTime not implemented")
}
func (UnimplementedTimeServiceServer) mustEmbedUnimplementedTimeServiceServer() {}

// UnsafeTimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeServiceServer will
// result in compilation errors.
type UnsafeTimeServiceServer interface {
	mustEmbedUnimplementedTimeServiceServer()
}

func RegisterTimeServiceServer(s grpc.ServiceRegistrar, srv TimeServiceServer) {
	s.RegisterService(&TimeService_ServiceDesc, srv)
}

func _TimeService_GetTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).GetTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregatingapis.time.v1.TimeService/GetTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).GetTime(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeService_ServiceDesc is the grpc.ServiceDesc for TimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aggregatingapis.time.v1.TimeService",
	HandlerType: (*TimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTime",
			Handler:    _TimeService_GetTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aggregatingapis/time/v1/time.proto",
}
