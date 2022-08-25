// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: campusapis/schoolTime/v1/school_time.proto

package schoolTimev1

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

// SchoolTimeServiceClient is the client API for SchoolTimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchoolTimeServiceClient interface {
	// 获取当前学校时间
	GetSchoolTime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSchoolTimeResponse, error)
	// 获取学校学期信息列表
	GetSemesterList(ctx context.Context, in *GetSemesterListRequest, opts ...grpc.CallOption) (*GetSemesterListResponse, error)
}

type schoolTimeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchoolTimeServiceClient(cc grpc.ClientConnInterface) SchoolTimeServiceClient {
	return &schoolTimeServiceClient{cc}
}

func (c *schoolTimeServiceClient) GetSchoolTime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSchoolTimeResponse, error) {
	out := new(GetSchoolTimeResponse)
	err := c.cc.Invoke(ctx, "/campusapis.schoolTime.v1.SchoolTimeService/GetSchoolTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schoolTimeServiceClient) GetSemesterList(ctx context.Context, in *GetSemesterListRequest, opts ...grpc.CallOption) (*GetSemesterListResponse, error) {
	out := new(GetSemesterListResponse)
	err := c.cc.Invoke(ctx, "/campusapis.schoolTime.v1.SchoolTimeService/GetSemesterList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchoolTimeServiceServer is the server API for SchoolTimeService service.
// All implementations must embed UnimplementedSchoolTimeServiceServer
// for forward compatibility
type SchoolTimeServiceServer interface {
	// 获取当前学校时间
	GetSchoolTime(context.Context, *emptypb.Empty) (*GetSchoolTimeResponse, error)
	// 获取学校学期信息列表
	GetSemesterList(context.Context, *GetSemesterListRequest) (*GetSemesterListResponse, error)
	mustEmbedUnimplementedSchoolTimeServiceServer()
}

// UnimplementedSchoolTimeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchoolTimeServiceServer struct {
}

func (UnimplementedSchoolTimeServiceServer) GetSchoolTime(context.Context, *emptypb.Empty) (*GetSchoolTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchoolTime not implemented")
}
func (UnimplementedSchoolTimeServiceServer) GetSemesterList(context.Context, *GetSemesterListRequest) (*GetSemesterListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSemesterList not implemented")
}
func (UnimplementedSchoolTimeServiceServer) mustEmbedUnimplementedSchoolTimeServiceServer() {}

// UnsafeSchoolTimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchoolTimeServiceServer will
// result in compilation errors.
type UnsafeSchoolTimeServiceServer interface {
	mustEmbedUnimplementedSchoolTimeServiceServer()
}

func RegisterSchoolTimeServiceServer(s grpc.ServiceRegistrar, srv SchoolTimeServiceServer) {
	s.RegisterService(&SchoolTimeService_ServiceDesc, srv)
}

func _SchoolTimeService_GetSchoolTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchoolTimeServiceServer).GetSchoolTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.schoolTime.v1.SchoolTimeService/GetSchoolTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchoolTimeServiceServer).GetSchoolTime(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchoolTimeService_GetSemesterList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSemesterListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchoolTimeServiceServer).GetSemesterList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.schoolTime.v1.SchoolTimeService/GetSemesterList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchoolTimeServiceServer).GetSemesterList(ctx, req.(*GetSemesterListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchoolTimeService_ServiceDesc is the grpc.ServiceDesc for SchoolTimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchoolTimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "campusapis.schoolTime.v1.SchoolTimeService",
	HandlerType: (*SchoolTimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSchoolTime",
			Handler:    _SchoolTimeService_GetSchoolTime_Handler,
		},
		{
			MethodName: "GetSemesterList",
			Handler:    _SchoolTimeService_GetSemesterList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "campusapis/schoolTime/v1/school_time.proto",
}
