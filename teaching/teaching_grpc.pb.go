// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package teaching

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

// TeachingServerClient is the client API for TeachingServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeachingServerClient interface {
	StaffGetScheduleHandler(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*Courses, error)
}

type teachingServerClient struct {
	cc grpc.ClientConnInterface
}

func NewTeachingServerClient(cc grpc.ClientConnInterface) TeachingServerClient {
	return &teachingServerClient{cc}
}

func (c *teachingServerClient) StaffGetScheduleHandler(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*Courses, error) {
	out := new(Courses)
	err := c.cc.Invoke(ctx, "/teaching.TeachingServer/StaffGetScheduleHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeachingServerServer is the server API for TeachingServer service.
// All implementations must embed UnimplementedTeachingServerServer
// for forward compatibility
type TeachingServerServer interface {
	StaffGetScheduleHandler(context.Context, *ScheduleRequest) (*Courses, error)
	mustEmbedUnimplementedTeachingServerServer()
}

// UnimplementedTeachingServerServer must be embedded to have forward compatible implementations.
type UnimplementedTeachingServerServer struct {
}

func (UnimplementedTeachingServerServer) StaffGetScheduleHandler(context.Context, *ScheduleRequest) (*Courses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffGetScheduleHandler not implemented")
}
func (UnimplementedTeachingServerServer) mustEmbedUnimplementedTeachingServerServer() {}

// UnsafeTeachingServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeachingServerServer will
// result in compilation errors.
type UnsafeTeachingServerServer interface {
	mustEmbedUnimplementedTeachingServerServer()
}

func RegisterTeachingServerServer(s grpc.ServiceRegistrar, srv TeachingServerServer) {
	s.RegisterService(&TeachingServer_ServiceDesc, srv)
}

func _TeachingServer_StaffGetScheduleHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachingServerServer).StaffGetScheduleHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teaching.TeachingServer/StaffGetScheduleHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachingServerServer).StaffGetScheduleHandler(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeachingServer_ServiceDesc is the grpc.ServiceDesc for TeachingServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeachingServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teaching.TeachingServer",
	HandlerType: (*TeachingServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StaffGetScheduleHandler",
			Handler:    _TeachingServer_StaffGetScheduleHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teaching/teaching.proto",
}
