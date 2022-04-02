// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: campusapis/teaching/v1/teaching.proto

package teachingv1

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

// TeachingServiceClient is the client API for TeachingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeachingServiceClient interface {
	// 获取学生/教师课程表
	GetSchedule(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*GetScheduleResponse, error)
	// 获取学生/教师当前课程表
	GetScheduleNow(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*GetScheduleNowResponse, error)
	// 获取学生/教师当前课程表V2 返回更详细的信息
	GetScheduleNowV2(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*GetScheduleNowV2Response, error)
	// 获取所有教室列表
	GetClassrooms(ctx context.Context, in *GetClassroomsRequest, opts ...grpc.CallOption) (*ClassroomsResponse, error)
	// 获取某一教室的使用情况
	GetClassroomUsages(ctx context.Context, in *GetClassroomUsageRequest, opts ...grpc.CallOption) (*ClassroomUsagesResponse, error)
	// 获取所有空教室列表
	GetUnusedClassrooms(ctx context.Context, in *GetUnusedClassroomsRequest, opts ...grpc.CallOption) (*ClassroomsResponse, error)
}

type teachingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeachingServiceClient(cc grpc.ClientConnInterface) TeachingServiceClient {
	return &teachingServiceClient{cc}
}

func (c *teachingServiceClient) GetSchedule(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*GetScheduleResponse, error) {
	out := new(GetScheduleResponse)
	err := c.cc.Invoke(ctx, "/campusapis.teaching.v1.TeachingService/GetSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachingServiceClient) GetScheduleNow(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*GetScheduleNowResponse, error) {
	out := new(GetScheduleNowResponse)
	err := c.cc.Invoke(ctx, "/campusapis.teaching.v1.TeachingService/GetScheduleNow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachingServiceClient) GetScheduleNowV2(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*GetScheduleNowV2Response, error) {
	out := new(GetScheduleNowV2Response)
	err := c.cc.Invoke(ctx, "/campusapis.teaching.v1.TeachingService/GetScheduleNowV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachingServiceClient) GetClassrooms(ctx context.Context, in *GetClassroomsRequest, opts ...grpc.CallOption) (*ClassroomsResponse, error) {
	out := new(ClassroomsResponse)
	err := c.cc.Invoke(ctx, "/campusapis.teaching.v1.TeachingService/GetClassrooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachingServiceClient) GetClassroomUsages(ctx context.Context, in *GetClassroomUsageRequest, opts ...grpc.CallOption) (*ClassroomUsagesResponse, error) {
	out := new(ClassroomUsagesResponse)
	err := c.cc.Invoke(ctx, "/campusapis.teaching.v1.TeachingService/GetClassroomUsages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachingServiceClient) GetUnusedClassrooms(ctx context.Context, in *GetUnusedClassroomsRequest, opts ...grpc.CallOption) (*ClassroomsResponse, error) {
	out := new(ClassroomsResponse)
	err := c.cc.Invoke(ctx, "/campusapis.teaching.v1.TeachingService/GetUnusedClassrooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeachingServiceServer is the server API for TeachingService service.
// All implementations should embed UnimplementedTeachingServiceServer
// for forward compatibility
type TeachingServiceServer interface {
	// 获取学生/教师课程表
	GetSchedule(context.Context, *GetScheduleRequest) (*GetScheduleResponse, error)
	// 获取学生/教师当前课程表
	GetScheduleNow(context.Context, *GetScheduleRequest) (*GetScheduleNowResponse, error)
	// 获取学生/教师当前课程表V2 返回更详细的信息
	GetScheduleNowV2(context.Context, *GetScheduleRequest) (*GetScheduleNowV2Response, error)
	// 获取所有教室列表
	GetClassrooms(context.Context, *GetClassroomsRequest) (*ClassroomsResponse, error)
	// 获取某一教室的使用情况
	GetClassroomUsages(context.Context, *GetClassroomUsageRequest) (*ClassroomUsagesResponse, error)
	// 获取所有空教室列表
	GetUnusedClassrooms(context.Context, *GetUnusedClassroomsRequest) (*ClassroomsResponse, error)
}

// UnimplementedTeachingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTeachingServiceServer struct {
}

func (UnimplementedTeachingServiceServer) GetSchedule(context.Context, *GetScheduleRequest) (*GetScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchedule not implemented")
}
func (UnimplementedTeachingServiceServer) GetScheduleNow(context.Context, *GetScheduleRequest) (*GetScheduleNowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduleNow not implemented")
}
func (UnimplementedTeachingServiceServer) GetScheduleNowV2(context.Context, *GetScheduleRequest) (*GetScheduleNowV2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduleNowV2 not implemented")
}
func (UnimplementedTeachingServiceServer) GetClassrooms(context.Context, *GetClassroomsRequest) (*ClassroomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassrooms not implemented")
}
func (UnimplementedTeachingServiceServer) GetClassroomUsages(context.Context, *GetClassroomUsageRequest) (*ClassroomUsagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassroomUsages not implemented")
}
func (UnimplementedTeachingServiceServer) GetUnusedClassrooms(context.Context, *GetUnusedClassroomsRequest) (*ClassroomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnusedClassrooms not implemented")
}

// UnsafeTeachingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeachingServiceServer will
// result in compilation errors.
type UnsafeTeachingServiceServer interface {
	mustEmbedUnimplementedTeachingServiceServer()
}

func RegisterTeachingServiceServer(s grpc.ServiceRegistrar, srv TeachingServiceServer) {
	s.RegisterService(&TeachingService_ServiceDesc, srv)
}

func _TeachingService_GetSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachingServiceServer).GetSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.teaching.v1.TeachingService/GetSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachingServiceServer).GetSchedule(ctx, req.(*GetScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachingService_GetScheduleNow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachingServiceServer).GetScheduleNow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.teaching.v1.TeachingService/GetScheduleNow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachingServiceServer).GetScheduleNow(ctx, req.(*GetScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachingService_GetScheduleNowV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachingServiceServer).GetScheduleNowV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.teaching.v1.TeachingService/GetScheduleNowV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachingServiceServer).GetScheduleNowV2(ctx, req.(*GetScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachingService_GetClassrooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassroomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachingServiceServer).GetClassrooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.teaching.v1.TeachingService/GetClassrooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachingServiceServer).GetClassrooms(ctx, req.(*GetClassroomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachingService_GetClassroomUsages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassroomUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachingServiceServer).GetClassroomUsages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.teaching.v1.TeachingService/GetClassroomUsages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachingServiceServer).GetClassroomUsages(ctx, req.(*GetClassroomUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachingService_GetUnusedClassrooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnusedClassroomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachingServiceServer).GetUnusedClassrooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.teaching.v1.TeachingService/GetUnusedClassrooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachingServiceServer).GetUnusedClassrooms(ctx, req.(*GetUnusedClassroomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeachingService_ServiceDesc is the grpc.ServiceDesc for TeachingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeachingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "campusapis.teaching.v1.TeachingService",
	HandlerType: (*TeachingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSchedule",
			Handler:    _TeachingService_GetSchedule_Handler,
		},
		{
			MethodName: "GetScheduleNow",
			Handler:    _TeachingService_GetScheduleNow_Handler,
		},
		{
			MethodName: "GetScheduleNowV2",
			Handler:    _TeachingService_GetScheduleNowV2_Handler,
		},
		{
			MethodName: "GetClassrooms",
			Handler:    _TeachingService_GetClassrooms_Handler,
		},
		{
			MethodName: "GetClassroomUsages",
			Handler:    _TeachingService_GetClassroomUsages_Handler,
		},
		{
			MethodName: "GetUnusedClassrooms",
			Handler:    _TeachingService_GetUnusedClassrooms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "campusapis/teaching/v1/teaching.proto",
}
