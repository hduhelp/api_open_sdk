// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package healthv1

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

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthServiceClient interface {
	//获取指定日期健康打卡记录
	GetCheckinRecord(ctx context.Context, in *GetCheckinRecordRequest, opts ...grpc.CallOption) (*GetCheckinRecordResponse, error)
	//获取历史健康打卡记录
	GetCheckinRecords(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCheckinRecordsResponse, error)
	//提交今日健康打卡
	PostCheckinRecord(ctx context.Context, in *PostCheckinRecordRequest, opts ...grpc.CallOption) (*PostCheckinRecordResponse, error)
	//获取提交健康码的手机号
	GetCheckinPhone(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCheckinPhoneResponse, error)
	//获取当前健康码
	GetHealthCode(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetHealthCodeResponse, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) GetCheckinRecord(ctx context.Context, in *GetCheckinRecordRequest, opts ...grpc.CallOption) (*GetCheckinRecordResponse, error) {
	out := new(GetCheckinRecordResponse)
	err := c.cc.Invoke(ctx, "/campusapis.health.v1.HealthService/GetCheckinRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) GetCheckinRecords(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCheckinRecordsResponse, error) {
	out := new(GetCheckinRecordsResponse)
	err := c.cc.Invoke(ctx, "/campusapis.health.v1.HealthService/GetCheckinRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) PostCheckinRecord(ctx context.Context, in *PostCheckinRecordRequest, opts ...grpc.CallOption) (*PostCheckinRecordResponse, error) {
	out := new(PostCheckinRecordResponse)
	err := c.cc.Invoke(ctx, "/campusapis.health.v1.HealthService/PostCheckinRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) GetCheckinPhone(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCheckinPhoneResponse, error) {
	out := new(GetCheckinPhoneResponse)
	err := c.cc.Invoke(ctx, "/campusapis.health.v1.HealthService/GetCheckinPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) GetHealthCode(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetHealthCodeResponse, error) {
	out := new(GetHealthCodeResponse)
	err := c.cc.Invoke(ctx, "/campusapis.health.v1.HealthService/GetHealthCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
// All implementations should embed UnimplementedHealthServiceServer
// for forward compatibility
type HealthServiceServer interface {
	//获取指定日期健康打卡记录
	GetCheckinRecord(context.Context, *GetCheckinRecordRequest) (*GetCheckinRecordResponse, error)
	//获取历史健康打卡记录
	GetCheckinRecords(context.Context, *emptypb.Empty) (*GetCheckinRecordsResponse, error)
	//提交今日健康打卡
	PostCheckinRecord(context.Context, *PostCheckinRecordRequest) (*PostCheckinRecordResponse, error)
	//获取提交健康码的手机号
	GetCheckinPhone(context.Context, *emptypb.Empty) (*GetCheckinPhoneResponse, error)
	//获取当前健康码
	GetHealthCode(context.Context, *emptypb.Empty) (*GetHealthCodeResponse, error)
}

// UnimplementedHealthServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (UnimplementedHealthServiceServer) GetCheckinRecord(context.Context, *GetCheckinRecordRequest) (*GetCheckinRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckinRecord not implemented")
}
func (UnimplementedHealthServiceServer) GetCheckinRecords(context.Context, *emptypb.Empty) (*GetCheckinRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckinRecords not implemented")
}
func (UnimplementedHealthServiceServer) PostCheckinRecord(context.Context, *PostCheckinRecordRequest) (*PostCheckinRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCheckinRecord not implemented")
}
func (UnimplementedHealthServiceServer) GetCheckinPhone(context.Context, *emptypb.Empty) (*GetCheckinPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckinPhone not implemented")
}
func (UnimplementedHealthServiceServer) GetHealthCode(context.Context, *emptypb.Empty) (*GetHealthCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealthCode not implemented")
}

// UnsafeHealthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServiceServer will
// result in compilation errors.
type UnsafeHealthServiceServer interface {
	mustEmbedUnimplementedHealthServiceServer()
}

func RegisterHealthServiceServer(s grpc.ServiceRegistrar, srv HealthServiceServer) {
	s.RegisterService(&HealthService_ServiceDesc, srv)
}

func _HealthService_GetCheckinRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCheckinRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).GetCheckinRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.health.v1.HealthService/GetCheckinRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).GetCheckinRecord(ctx, req.(*GetCheckinRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_GetCheckinRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).GetCheckinRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.health.v1.HealthService/GetCheckinRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).GetCheckinRecords(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_PostCheckinRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostCheckinRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).PostCheckinRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.health.v1.HealthService/PostCheckinRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).PostCheckinRecord(ctx, req.(*PostCheckinRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_GetCheckinPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).GetCheckinPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.health.v1.HealthService/GetCheckinPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).GetCheckinPhone(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_GetHealthCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).GetHealthCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.health.v1.HealthService/GetHealthCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).GetHealthCode(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthService_ServiceDesc is the grpc.ServiceDesc for HealthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "campusapis.health.v1.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCheckinRecord",
			Handler:    _HealthService_GetCheckinRecord_Handler,
		},
		{
			MethodName: "GetCheckinRecords",
			Handler:    _HealthService_GetCheckinRecords_Handler,
		},
		{
			MethodName: "PostCheckinRecord",
			Handler:    _HealthService_PostCheckinRecord_Handler,
		},
		{
			MethodName: "GetCheckinPhone",
			Handler:    _HealthService_GetCheckinPhone_Handler,
		},
		{
			MethodName: "GetHealthCode",
			Handler:    _HealthService_GetHealthCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "campusapis/health/v1/health.proto",
}
