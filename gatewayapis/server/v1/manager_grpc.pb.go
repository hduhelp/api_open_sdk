// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gatewayapis/server/v1/manager.proto

package serverv1

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

// ServerManageerServiceClient is the client API for ServerManageerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerManageerServiceClient interface {
	//获取可管理的所有HTTP路由
	GetHTTPRouterList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetHTTPRouterListResponse, error)
	//查询路由
	GetHTTPRouter(ctx context.Context, in *HTTPRouterKeyRequest, opts ...grpc.CallOption) (*HTTPRouter, error)
	//创建或者更新路由
	PostOrPutHTTPRouter(ctx context.Context, in *HTTPRouter, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//删除路由
	DeleteHTTPRouter(ctx context.Context, in *HTTPRouterKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type serverManageerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerManageerServiceClient(cc grpc.ClientConnInterface) ServerManageerServiceClient {
	return &serverManageerServiceClient{cc}
}

func (c *serverManageerServiceClient) GetHTTPRouterList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetHTTPRouterListResponse, error) {
	out := new(GetHTTPRouterListResponse)
	err := c.cc.Invoke(ctx, "/gatewayapis.server.v1.ServerManageerService/GetHTTPRouterList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverManageerServiceClient) GetHTTPRouter(ctx context.Context, in *HTTPRouterKeyRequest, opts ...grpc.CallOption) (*HTTPRouter, error) {
	out := new(HTTPRouter)
	err := c.cc.Invoke(ctx, "/gatewayapis.server.v1.ServerManageerService/GetHTTPRouter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverManageerServiceClient) PostOrPutHTTPRouter(ctx context.Context, in *HTTPRouter, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/gatewayapis.server.v1.ServerManageerService/PostOrPutHTTPRouter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverManageerServiceClient) DeleteHTTPRouter(ctx context.Context, in *HTTPRouterKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/gatewayapis.server.v1.ServerManageerService/DeleteHTTPRouter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerManageerServiceServer is the server API for ServerManageerService service.
// All implementations should embed UnimplementedServerManageerServiceServer
// for forward compatibility
type ServerManageerServiceServer interface {
	//获取可管理的所有HTTP路由
	GetHTTPRouterList(context.Context, *emptypb.Empty) (*GetHTTPRouterListResponse, error)
	//查询路由
	GetHTTPRouter(context.Context, *HTTPRouterKeyRequest) (*HTTPRouter, error)
	//创建或者更新路由
	PostOrPutHTTPRouter(context.Context, *HTTPRouter) (*emptypb.Empty, error)
	//删除路由
	DeleteHTTPRouter(context.Context, *HTTPRouterKeyRequest) (*emptypb.Empty, error)
}

// UnimplementedServerManageerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedServerManageerServiceServer struct {
}

func (UnimplementedServerManageerServiceServer) GetHTTPRouterList(context.Context, *emptypb.Empty) (*GetHTTPRouterListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHTTPRouterList not implemented")
}
func (UnimplementedServerManageerServiceServer) GetHTTPRouter(context.Context, *HTTPRouterKeyRequest) (*HTTPRouter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHTTPRouter not implemented")
}
func (UnimplementedServerManageerServiceServer) PostOrPutHTTPRouter(context.Context, *HTTPRouter) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostOrPutHTTPRouter not implemented")
}
func (UnimplementedServerManageerServiceServer) DeleteHTTPRouter(context.Context, *HTTPRouterKeyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHTTPRouter not implemented")
}

// UnsafeServerManageerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerManageerServiceServer will
// result in compilation errors.
type UnsafeServerManageerServiceServer interface {
	mustEmbedUnimplementedServerManageerServiceServer()
}

func RegisterServerManageerServiceServer(s grpc.ServiceRegistrar, srv ServerManageerServiceServer) {
	s.RegisterService(&ServerManageerService_ServiceDesc, srv)
}

func _ServerManageerService_GetHTTPRouterList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerManageerServiceServer).GetHTTPRouterList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gatewayapis.server.v1.ServerManageerService/GetHTTPRouterList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerManageerServiceServer).GetHTTPRouterList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerManageerService_GetHTTPRouter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPRouterKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerManageerServiceServer).GetHTTPRouter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gatewayapis.server.v1.ServerManageerService/GetHTTPRouter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerManageerServiceServer).GetHTTPRouter(ctx, req.(*HTTPRouterKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerManageerService_PostOrPutHTTPRouter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPRouter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerManageerServiceServer).PostOrPutHTTPRouter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gatewayapis.server.v1.ServerManageerService/PostOrPutHTTPRouter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerManageerServiceServer).PostOrPutHTTPRouter(ctx, req.(*HTTPRouter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerManageerService_DeleteHTTPRouter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPRouterKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerManageerServiceServer).DeleteHTTPRouter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gatewayapis.server.v1.ServerManageerService/DeleteHTTPRouter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerManageerServiceServer).DeleteHTTPRouter(ctx, req.(*HTTPRouterKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerManageerService_ServiceDesc is the grpc.ServiceDesc for ServerManageerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerManageerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gatewayapis.server.v1.ServerManageerService",
	HandlerType: (*ServerManageerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHTTPRouterList",
			Handler:    _ServerManageerService_GetHTTPRouterList_Handler,
		},
		{
			MethodName: "GetHTTPRouter",
			Handler:    _ServerManageerService_GetHTTPRouter_Handler,
		},
		{
			MethodName: "PostOrPutHTTPRouter",
			Handler:    _ServerManageerService_PostOrPutHTTPRouter_Handler,
		},
		{
			MethodName: "DeleteHTTPRouter",
			Handler:    _ServerManageerService_DeleteHTTPRouter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gatewayapis/server/v1/manager.proto",
}
