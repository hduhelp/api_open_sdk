// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DingTalkService_CreateClassChatGroup_FullMethodName   = "/consultingapis.dingtalk.v1.DingTalkService/CreateClassChatGroup"
	DingTalkService_GetClassChatGroup_FullMethodName      = "/consultingapis.dingtalk.v1.DingTalkService/GetClassChatGroup"
	DingTalkService_MergeClassChatGroup_FullMethodName    = "/consultingapis.dingtalk.v1.DingTalkService/MergeClassChatGroup"
	DingTalkService_SyncClassChatGroup_FullMethodName     = "/consultingapis.dingtalk.v1.DingTalkService/SyncClassChatGroup"
	DingTalkService_CreateGradeChatGroup_FullMethodName   = "/consultingapis.dingtalk.v1.DingTalkService/CreateGradeChatGroup"
	DingTalkService_GetGradeChatGroup_FullMethodName      = "/consultingapis.dingtalk.v1.DingTalkService/GetGradeChatGroup"
	DingTalkService_SyncGradeChatGroup_FullMethodName     = "/consultingapis.dingtalk.v1.DingTalkService/SyncGradeChatGroup"
	DingTalkService_CreateRecruitChatGroup_FullMethodName = "/consultingapis.dingtalk.v1.DingTalkService/CreateRecruitChatGroup"
)

// DingTalkServiceClient is the client API for DingTalkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DingTalkServiceClient interface {
	CreateClassChatGroup(ctx context.Context, in *CreateClassChatGroupRequest, opts ...grpc.CallOption) (*CreateClassChatGroupResponse, error)
	GetClassChatGroup(ctx context.Context, in *GetClassChatGroupRequest, opts ...grpc.CallOption) (*GetClassChatGroupResponse, error)
	MergeClassChatGroup(ctx context.Context, in *MergeClassChatGroupRequest, opts ...grpc.CallOption) (*MergeClassChatGroupResponse, error)
	SyncClassChatGroup(ctx context.Context, in *SyncClassChatGroupRequest, opts ...grpc.CallOption) (*SyncClassChatGroupResponse, error)
	CreateGradeChatGroup(ctx context.Context, in *CreateGradeChatGroupRequest, opts ...grpc.CallOption) (*CreateGradeChatGroupResponse, error)
	GetGradeChatGroup(ctx context.Context, in *GetGradeChatGroupRequest, opts ...grpc.CallOption) (*GetGradeChatGroupResponse, error)
	SyncGradeChatGroup(ctx context.Context, in *SyncGradeChatGroupRequest, opts ...grpc.CallOption) (*SyncGradeChatGroupResponse, error)
	CreateRecruitChatGroup(ctx context.Context, in *CreateRecruitChatGroupRequest, opts ...grpc.CallOption) (*CreateRecruitChatGroupResponse, error)
}

type dingTalkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDingTalkServiceClient(cc grpc.ClientConnInterface) DingTalkServiceClient {
	return &dingTalkServiceClient{cc}
}

func (c *dingTalkServiceClient) CreateClassChatGroup(ctx context.Context, in *CreateClassChatGroupRequest, opts ...grpc.CallOption) (*CreateClassChatGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateClassChatGroupResponse)
	err := c.cc.Invoke(ctx, DingTalkService_CreateClassChatGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dingTalkServiceClient) GetClassChatGroup(ctx context.Context, in *GetClassChatGroupRequest, opts ...grpc.CallOption) (*GetClassChatGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetClassChatGroupResponse)
	err := c.cc.Invoke(ctx, DingTalkService_GetClassChatGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dingTalkServiceClient) MergeClassChatGroup(ctx context.Context, in *MergeClassChatGroupRequest, opts ...grpc.CallOption) (*MergeClassChatGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MergeClassChatGroupResponse)
	err := c.cc.Invoke(ctx, DingTalkService_MergeClassChatGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dingTalkServiceClient) SyncClassChatGroup(ctx context.Context, in *SyncClassChatGroupRequest, opts ...grpc.CallOption) (*SyncClassChatGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncClassChatGroupResponse)
	err := c.cc.Invoke(ctx, DingTalkService_SyncClassChatGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dingTalkServiceClient) CreateGradeChatGroup(ctx context.Context, in *CreateGradeChatGroupRequest, opts ...grpc.CallOption) (*CreateGradeChatGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGradeChatGroupResponse)
	err := c.cc.Invoke(ctx, DingTalkService_CreateGradeChatGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dingTalkServiceClient) GetGradeChatGroup(ctx context.Context, in *GetGradeChatGroupRequest, opts ...grpc.CallOption) (*GetGradeChatGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGradeChatGroupResponse)
	err := c.cc.Invoke(ctx, DingTalkService_GetGradeChatGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dingTalkServiceClient) SyncGradeChatGroup(ctx context.Context, in *SyncGradeChatGroupRequest, opts ...grpc.CallOption) (*SyncGradeChatGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncGradeChatGroupResponse)
	err := c.cc.Invoke(ctx, DingTalkService_SyncGradeChatGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dingTalkServiceClient) CreateRecruitChatGroup(ctx context.Context, in *CreateRecruitChatGroupRequest, opts ...grpc.CallOption) (*CreateRecruitChatGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRecruitChatGroupResponse)
	err := c.cc.Invoke(ctx, DingTalkService_CreateRecruitChatGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DingTalkServiceServer is the server API for DingTalkService service.
// All implementations must embed UnimplementedDingTalkServiceServer
// for forward compatibility.
type DingTalkServiceServer interface {
	CreateClassChatGroup(context.Context, *CreateClassChatGroupRequest) (*CreateClassChatGroupResponse, error)
	GetClassChatGroup(context.Context, *GetClassChatGroupRequest) (*GetClassChatGroupResponse, error)
	MergeClassChatGroup(context.Context, *MergeClassChatGroupRequest) (*MergeClassChatGroupResponse, error)
	SyncClassChatGroup(context.Context, *SyncClassChatGroupRequest) (*SyncClassChatGroupResponse, error)
	CreateGradeChatGroup(context.Context, *CreateGradeChatGroupRequest) (*CreateGradeChatGroupResponse, error)
	GetGradeChatGroup(context.Context, *GetGradeChatGroupRequest) (*GetGradeChatGroupResponse, error)
	SyncGradeChatGroup(context.Context, *SyncGradeChatGroupRequest) (*SyncGradeChatGroupResponse, error)
	CreateRecruitChatGroup(context.Context, *CreateRecruitChatGroupRequest) (*CreateRecruitChatGroupResponse, error)
	mustEmbedUnimplementedDingTalkServiceServer()
}

// UnimplementedDingTalkServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDingTalkServiceServer struct{}

func (UnimplementedDingTalkServiceServer) CreateClassChatGroup(context.Context, *CreateClassChatGroupRequest) (*CreateClassChatGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClassChatGroup not implemented")
}
func (UnimplementedDingTalkServiceServer) GetClassChatGroup(context.Context, *GetClassChatGroupRequest) (*GetClassChatGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassChatGroup not implemented")
}
func (UnimplementedDingTalkServiceServer) MergeClassChatGroup(context.Context, *MergeClassChatGroupRequest) (*MergeClassChatGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MergeClassChatGroup not implemented")
}
func (UnimplementedDingTalkServiceServer) SyncClassChatGroup(context.Context, *SyncClassChatGroupRequest) (*SyncClassChatGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncClassChatGroup not implemented")
}
func (UnimplementedDingTalkServiceServer) CreateGradeChatGroup(context.Context, *CreateGradeChatGroupRequest) (*CreateGradeChatGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGradeChatGroup not implemented")
}
func (UnimplementedDingTalkServiceServer) GetGradeChatGroup(context.Context, *GetGradeChatGroupRequest) (*GetGradeChatGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGradeChatGroup not implemented")
}
func (UnimplementedDingTalkServiceServer) SyncGradeChatGroup(context.Context, *SyncGradeChatGroupRequest) (*SyncGradeChatGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncGradeChatGroup not implemented")
}
func (UnimplementedDingTalkServiceServer) CreateRecruitChatGroup(context.Context, *CreateRecruitChatGroupRequest) (*CreateRecruitChatGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecruitChatGroup not implemented")
}
func (UnimplementedDingTalkServiceServer) mustEmbedUnimplementedDingTalkServiceServer() {}
func (UnimplementedDingTalkServiceServer) testEmbeddedByValue()                         {}

// UnsafeDingTalkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DingTalkServiceServer will
// result in compilation errors.
type UnsafeDingTalkServiceServer interface {
	mustEmbedUnimplementedDingTalkServiceServer()
}

func RegisterDingTalkServiceServer(s grpc.ServiceRegistrar, srv DingTalkServiceServer) {
	// If the following call pancis, it indicates UnimplementedDingTalkServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
		FullMethod: DingTalkService_CreateClassChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DingTalkServiceServer).CreateClassChatGroup(ctx, req.(*CreateClassChatGroupRequest))
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
		FullMethod: DingTalkService_GetClassChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DingTalkServiceServer).GetClassChatGroup(ctx, req.(*GetClassChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DingTalkService_MergeClassChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeClassChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DingTalkServiceServer).MergeClassChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DingTalkService_MergeClassChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DingTalkServiceServer).MergeClassChatGroup(ctx, req.(*MergeClassChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DingTalkService_SyncClassChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncClassChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DingTalkServiceServer).SyncClassChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DingTalkService_SyncClassChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DingTalkServiceServer).SyncClassChatGroup(ctx, req.(*SyncClassChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DingTalkService_CreateGradeChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGradeChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DingTalkServiceServer).CreateGradeChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DingTalkService_CreateGradeChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DingTalkServiceServer).CreateGradeChatGroup(ctx, req.(*CreateGradeChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DingTalkService_GetGradeChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGradeChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DingTalkServiceServer).GetGradeChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DingTalkService_GetGradeChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DingTalkServiceServer).GetGradeChatGroup(ctx, req.(*GetGradeChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DingTalkService_SyncGradeChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncGradeChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DingTalkServiceServer).SyncGradeChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DingTalkService_SyncGradeChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DingTalkServiceServer).SyncGradeChatGroup(ctx, req.(*SyncGradeChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DingTalkService_CreateRecruitChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecruitChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DingTalkServiceServer).CreateRecruitChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DingTalkService_CreateRecruitChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DingTalkServiceServer).CreateRecruitChatGroup(ctx, req.(*CreateRecruitChatGroupRequest))
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
			MethodName: "GetClassChatGroup",
			Handler:    _DingTalkService_GetClassChatGroup_Handler,
		},
		{
			MethodName: "MergeClassChatGroup",
			Handler:    _DingTalkService_MergeClassChatGroup_Handler,
		},
		{
			MethodName: "SyncClassChatGroup",
			Handler:    _DingTalkService_SyncClassChatGroup_Handler,
		},
		{
			MethodName: "CreateGradeChatGroup",
			Handler:    _DingTalkService_CreateGradeChatGroup_Handler,
		},
		{
			MethodName: "GetGradeChatGroup",
			Handler:    _DingTalkService_GetGradeChatGroup_Handler,
		},
		{
			MethodName: "SyncGradeChatGroup",
			Handler:    _DingTalkService_SyncGradeChatGroup_Handler,
		},
		{
			MethodName: "CreateRecruitChatGroup",
			Handler:    _DingTalkService_CreateRecruitChatGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consultingapis/dingtalk/v1/dingtalk.proto",
}
