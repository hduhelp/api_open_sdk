// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: consultingapis/dingtalk/v1/dingtalk.proto

package dingtalkv1

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

// DingTalkServiceClient is the client API for DingTalkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DingTalkServiceClient interface {
	CreateClassChatGroup(ctx context.Context, in *CreateClassChatGroupRequest, opts ...grpc.CallOption) (*CreateClassChatGroupResponse, error)
	CreateClassChatGroupV2(ctx context.Context, in *CreateClassChatGroupV2Request, opts ...grpc.CallOption) (*CreateClassChatGroupV2Response, error)
	GetClassChatGroup(ctx context.Context, in *GetClassChatGroupRequest, opts ...grpc.CallOption) (*GetClassChatGroupResponse, error)
	JoinClassChatGroup(ctx context.Context, in *JoinClassChatGroupRequest, opts ...grpc.CallOption) (*JoinClassChatGroupResponse, error)
}

type dingTalkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDingTalkServiceClient(cc grpc.ClientConnInterface) DingTalkServiceClient {
	return &dingTalkServiceClient{cc}
}

func (c *dingTalkServiceClient) CreateClassChatGroup(ctx context.Context, in *CreateClassChatGroupRequest, opts ...grpc.CallOption) (*CreateClassChatGroupResponse, error) {
	out := new(CreateClassChatGroupResponse)
	err := c.cc.Invoke(ctx, "/consultingapis.dingtalk.v1.DingTalkService/CreateClassChatGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dingTalkServiceClient) CreateClassChatGroupV2(ctx context.Context, in *CreateClassChatGroupV2Request, opts ...grpc.CallOption) (*CreateClassChatGroupV2Response, error) {
	out := new(CreateClassChatGroupV2Response)
	err := c.cc.Invoke(ctx, "/consultingapis.dingtalk.v1.DingTalkService/CreateClassChatGroupV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dingTalkServiceClient) GetClassChatGroup(ctx context.Context, in *GetClassChatGroupRequest, opts ...grpc.CallOption) (*GetClassChatGroupResponse, error) {
	out := new(GetClassChatGroupResponse)
	err := c.cc.Invoke(ctx, "/consultingapis.dingtalk.v1.DingTalkService/GetClassChatGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dingTalkServiceClient) JoinClassChatGroup(ctx context.Context, in *JoinClassChatGroupRequest, opts ...grpc.CallOption) (*JoinClassChatGroupResponse, error) {
	out := new(JoinClassChatGroupResponse)
	err := c.cc.Invoke(ctx, "/consultingapis.dingtalk.v1.DingTalkService/JoinClassChatGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DingTalkServiceServer is the server API for DingTalkService service.
// All implementations must embed UnimplementedDingTalkServiceServer
// for forward compatibility
type DingTalkServiceServer interface {
	CreateClassChatGroup(context.Context, *CreateClassChatGroupRequest) (*CreateClassChatGroupResponse, error)
	CreateClassChatGroupV2(context.Context, *CreateClassChatGroupV2Request) (*CreateClassChatGroupV2Response, error)
	GetClassChatGroup(context.Context, *GetClassChatGroupRequest) (*GetClassChatGroupResponse, error)
	JoinClassChatGroup(context.Context, *JoinClassChatGroupRequest) (*JoinClassChatGroupResponse, error)
	mustEmbedUnimplementedDingTalkServiceServer()
}

// UnimplementedDingTalkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDingTalkServiceServer struct {
}

func (UnimplementedDingTalkServiceServer) CreateClassChatGroup(context.Context, *CreateClassChatGroupRequest) (*CreateClassChatGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClassChatGroup not implemented")
}
func (UnimplementedDingTalkServiceServer) CreateClassChatGroupV2(context.Context, *CreateClassChatGroupV2Request) (*CreateClassChatGroupV2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClassChatGroupV2 not implemented")
}
func (UnimplementedDingTalkServiceServer) GetClassChatGroup(context.Context, *GetClassChatGroupRequest) (*GetClassChatGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassChatGroup not implemented")
}
func (UnimplementedDingTalkServiceServer) JoinClassChatGroup(context.Context, *JoinClassChatGroupRequest) (*JoinClassChatGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinClassChatGroup not implemented")
}
func (UnimplementedDingTalkServiceServer) mustEmbedUnimplementedDingTalkServiceServer() {}

// UnsafeDingTalkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DingTalkServiceServer will
// result in compilation errors.
type UnsafeDingTalkServiceServer interface {
	mustEmbedUnimplementedDingTalkServiceServer()
}

func RegisterDingTalkServiceServer(s grpc.ServiceRegistrar, srv DingTalkServiceServer) {
	s.RegisterService(&DingTalkService_ServiceDesc, srv)
}

func _DingTalkService_CreateClassChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClassChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DingTalkServiceServer).CreateClassChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consultingapis.dingtalk.v1.DingTalkService/CreateClassChatGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DingTalkServiceServer).CreateClassChatGroup(ctx, req.(*CreateClassChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DingTalkService_CreateClassChatGroupV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClassChatGroupV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DingTalkServiceServer).CreateClassChatGroupV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consultingapis.dingtalk.v1.DingTalkService/CreateClassChatGroupV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DingTalkServiceServer).CreateClassChatGroupV2(ctx, req.(*CreateClassChatGroupV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DingTalkService_GetClassChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DingTalkServiceServer).GetClassChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consultingapis.dingtalk.v1.DingTalkService/GetClassChatGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DingTalkServiceServer).GetClassChatGroup(ctx, req.(*GetClassChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DingTalkService_JoinClassChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinClassChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DingTalkServiceServer).JoinClassChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consultingapis.dingtalk.v1.DingTalkService/JoinClassChatGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DingTalkServiceServer).JoinClassChatGroup(ctx, req.(*JoinClassChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DingTalkService_ServiceDesc is the grpc.ServiceDesc for DingTalkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DingTalkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "consultingapis.dingtalk.v1.DingTalkService",
	HandlerType: (*DingTalkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClassChatGroup",
			Handler:    _DingTalkService_CreateClassChatGroup_Handler,
		},
		{
			MethodName: "CreateClassChatGroupV2",
			Handler:    _DingTalkService_CreateClassChatGroupV2_Handler,
		},
		{
			MethodName: "GetClassChatGroup",
			Handler:    _DingTalkService_GetClassChatGroup_Handler,
		},
		{
			MethodName: "JoinClassChatGroup",
			Handler:    _DingTalkService_JoinClassChatGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consultingapis/dingtalk/v1/dingtalk.proto",
}
