// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: campusapis/skl/v1/skl.proto

package sklv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SklService_GetStaffUnitInfo_FullMethodName   = "/campusapis.skl.v1.SklService/GetStaffUnitInfo"
	SklService_GetStaffUnitDetail_FullMethodName = "/campusapis.skl.v1.SklService/GetStaffUnitDetail"
)

// SklServiceClient is the client API for SklService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SklServiceClient interface {
	// 获取某学院某年级的学生信息
	GetStaffUnitInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StaffUnitInfoResponse, error)
	GetStaffUnitDetail(ctx context.Context, in *StaffUnitDetailRequest, opts ...grpc.CallOption) (*StaffUnitDetailResponse, error)
}

type sklServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSklServiceClient(cc grpc.ClientConnInterface) SklServiceClient {
	return &sklServiceClient{cc}
}

func (c *sklServiceClient) GetStaffUnitInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StaffUnitInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StaffUnitInfoResponse)
	err := c.cc.Invoke(ctx, SklService_GetStaffUnitInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sklServiceClient) GetStaffUnitDetail(ctx context.Context, in *StaffUnitDetailRequest, opts ...grpc.CallOption) (*StaffUnitDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StaffUnitDetailResponse)
	err := c.cc.Invoke(ctx, SklService_GetStaffUnitDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SklServiceServer is the server API for SklService service.
// All implementations must embed UnimplementedSklServiceServer
// for forward compatibility.
type SklServiceServer interface {
	// 获取某学院某年级的学生信息
	GetStaffUnitInfo(context.Context, *emptypb.Empty) (*StaffUnitInfoResponse, error)
	GetStaffUnitDetail(context.Context, *StaffUnitDetailRequest) (*StaffUnitDetailResponse, error)
	mustEmbedUnimplementedSklServiceServer()
}

// UnimplementedSklServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSklServiceServer struct{}

func (UnimplementedSklServiceServer) GetStaffUnitInfo(context.Context, *emptypb.Empty) (*StaffUnitInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaffUnitInfo not implemented")
}
func (UnimplementedSklServiceServer) GetStaffUnitDetail(context.Context, *StaffUnitDetailRequest) (*StaffUnitDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaffUnitDetail not implemented")
}
func (UnimplementedSklServiceServer) mustEmbedUnimplementedSklServiceServer() {}
func (UnimplementedSklServiceServer) testEmbeddedByValue()                    {}

// UnsafeSklServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SklServiceServer will
// result in compilation errors.
type UnsafeSklServiceServer interface {
	mustEmbedUnimplementedSklServiceServer()
}

func RegisterSklServiceServer(s grpc.ServiceRegistrar, srv SklServiceServer) {
	// If the following call pancis, it indicates UnimplementedSklServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SklService_ServiceDesc, srv)
}

func _SklService_GetStaffUnitInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SklServiceServer).GetStaffUnitInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SklService_GetStaffUnitInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SklServiceServer).GetStaffUnitInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SklService_GetStaffUnitDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffUnitDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SklServiceServer).GetStaffUnitDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SklService_GetStaffUnitDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SklServiceServer).GetStaffUnitDetail(ctx, req.(*StaffUnitDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SklService_ServiceDesc is the grpc.ServiceDesc for SklService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SklService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "campusapis.skl.v1.SklService",
	HandlerType: (*SklServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStaffUnitInfo",
			Handler:    _SklService_GetStaffUnitInfo_Handler,
		},
		{
			MethodName: "GetStaffUnitDetail",
			Handler:    _SklService_GetStaffUnitDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "campusapis/skl/v1/skl.proto",
}
