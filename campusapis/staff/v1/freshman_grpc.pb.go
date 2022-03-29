// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package staffv1

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

// FreshmanServiceClient is the client API for FreshmanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FreshmanServiceClient interface {
	//获取新生基本信息
	GetFreshmanBaseInfo(ctx context.Context, in *FreshmanKeywordRequest, opts ...grpc.CallOption) (*GetFreshmanBaseInfoResponse, error)
	//获取新生详细信息
	GetFreshmanInfo(ctx context.Context, in *FreshmanKeywordRequest, opts ...grpc.CallOption) (*GetFreshmanInfoResponse, error)
	//获取新生宿舍信息
	GetFreshmanRoommates(ctx context.Context, in *FreshmanKeywordRequest, opts ...grpc.CallOption) (*GetFreshmanRoommatesResponse, error)
}

type freshmanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFreshmanServiceClient(cc grpc.ClientConnInterface) FreshmanServiceClient {
	return &freshmanServiceClient{cc}
}

func (c *freshmanServiceClient) GetFreshmanBaseInfo(ctx context.Context, in *FreshmanKeywordRequest, opts ...grpc.CallOption) (*GetFreshmanBaseInfoResponse, error) {
	out := new(GetFreshmanBaseInfoResponse)
	err := c.cc.Invoke(ctx, "/campusapis.staff.v1.FreshmanService/GetFreshmanBaseInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freshmanServiceClient) GetFreshmanInfo(ctx context.Context, in *FreshmanKeywordRequest, opts ...grpc.CallOption) (*GetFreshmanInfoResponse, error) {
	out := new(GetFreshmanInfoResponse)
	err := c.cc.Invoke(ctx, "/campusapis.staff.v1.FreshmanService/GetFreshmanInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freshmanServiceClient) GetFreshmanRoommates(ctx context.Context, in *FreshmanKeywordRequest, opts ...grpc.CallOption) (*GetFreshmanRoommatesResponse, error) {
	out := new(GetFreshmanRoommatesResponse)
	err := c.cc.Invoke(ctx, "/campusapis.staff.v1.FreshmanService/GetFreshmanRoommates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FreshmanServiceServer is the server API for FreshmanService service.
// All implementations should embed UnimplementedFreshmanServiceServer
// for forward compatibility
type FreshmanServiceServer interface {
	//获取新生基本信息
	GetFreshmanBaseInfo(context.Context, *FreshmanKeywordRequest) (*GetFreshmanBaseInfoResponse, error)
	//获取新生详细信息
	GetFreshmanInfo(context.Context, *FreshmanKeywordRequest) (*GetFreshmanInfoResponse, error)
	//获取新生宿舍信息
	GetFreshmanRoommates(context.Context, *FreshmanKeywordRequest) (*GetFreshmanRoommatesResponse, error)
}

// UnimplementedFreshmanServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFreshmanServiceServer struct {
}

func (UnimplementedFreshmanServiceServer) GetFreshmanBaseInfo(context.Context, *FreshmanKeywordRequest) (*GetFreshmanBaseInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreshmanBaseInfo not implemented")
}
func (UnimplementedFreshmanServiceServer) GetFreshmanInfo(context.Context, *FreshmanKeywordRequest) (*GetFreshmanInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreshmanInfo not implemented")
}
func (UnimplementedFreshmanServiceServer) GetFreshmanRoommates(context.Context, *FreshmanKeywordRequest) (*GetFreshmanRoommatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreshmanRoommates not implemented")
}

// UnsafeFreshmanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FreshmanServiceServer will
// result in compilation errors.
type UnsafeFreshmanServiceServer interface {
	mustEmbedUnimplementedFreshmanServiceServer()
}

func RegisterFreshmanServiceServer(s grpc.ServiceRegistrar, srv FreshmanServiceServer) {
	s.RegisterService(&FreshmanService_ServiceDesc, srv)
}

func _FreshmanService_GetFreshmanBaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreshmanKeywordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreshmanServiceServer).GetFreshmanBaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.staff.v1.FreshmanService/GetFreshmanBaseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreshmanServiceServer).GetFreshmanBaseInfo(ctx, req.(*FreshmanKeywordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreshmanService_GetFreshmanInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreshmanKeywordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreshmanServiceServer).GetFreshmanInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.staff.v1.FreshmanService/GetFreshmanInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreshmanServiceServer).GetFreshmanInfo(ctx, req.(*FreshmanKeywordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreshmanService_GetFreshmanRoommates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreshmanKeywordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreshmanServiceServer).GetFreshmanRoommates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.staff.v1.FreshmanService/GetFreshmanRoommates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreshmanServiceServer).GetFreshmanRoommates(ctx, req.(*FreshmanKeywordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FreshmanService_ServiceDesc is the grpc.ServiceDesc for FreshmanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FreshmanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "campusapis.staff.v1.FreshmanService",
	HandlerType: (*FreshmanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFreshmanBaseInfo",
			Handler:    _FreshmanService_GetFreshmanBaseInfo_Handler,
		},
		{
			MethodName: "GetFreshmanInfo",
			Handler:    _FreshmanService_GetFreshmanInfo_Handler,
		},
		{
			MethodName: "GetFreshmanRoommates",
			Handler:    _FreshmanService_GetFreshmanRoommates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "campusapis/staff/v1/freshman.proto",
}
