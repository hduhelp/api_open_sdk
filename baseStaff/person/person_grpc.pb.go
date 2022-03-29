// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package person

import (
	context "context"
	baseStaff "github.com/hduhelp/api_open_sdk/baseStaff"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PersonServerClient is the client API for PersonServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonServerClient interface {
	StaffGetPersonInfoHandler(ctx context.Context, in *baseStaff.Staff, opts ...grpc.CallOption) (*Info, error)
}

type personServerClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServerClient(cc grpc.ClientConnInterface) PersonServerClient {
	return &personServerClient{cc}
}

func (c *personServerClient) StaffGetPersonInfoHandler(ctx context.Context, in *baseStaff.Staff, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/person.PersonServer/StaffGetPersonInfoHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServerServer is the server API for PersonServer service.
// All implementations should embed UnimplementedPersonServerServer
// for forward compatibility
type PersonServerServer interface {
	StaffGetPersonInfoHandler(context.Context, *baseStaff.Staff) (*Info, error)
}

// UnimplementedPersonServerServer should be embedded to have forward compatible implementations.
type UnimplementedPersonServerServer struct {
}

func (UnimplementedPersonServerServer) StaffGetPersonInfoHandler(context.Context, *baseStaff.Staff) (*Info, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffGetPersonInfoHandler not implemented")
}

// UnsafePersonServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServerServer will
// result in compilation errors.
type UnsafePersonServerServer interface {
	mustEmbedUnimplementedPersonServerServer()
}

func RegisterPersonServerServer(s grpc.ServiceRegistrar, srv PersonServerServer) {
	s.RegisterService(&PersonServer_ServiceDesc, srv)
}

func _PersonServer_StaffGetPersonInfoHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(baseStaff.Staff)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServerServer).StaffGetPersonInfoHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/person.PersonServer/StaffGetPersonInfoHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServerServer).StaffGetPersonInfoHandler(ctx, req.(*baseStaff.Staff))
	}
	return interceptor(ctx, in, info, handler)
}

// PersonServer_ServiceDesc is the grpc.ServiceDesc for PersonServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "person.PersonServer",
	HandlerType: (*PersonServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StaffGetPersonInfoHandler",
			Handler:    _PersonServer_StaffGetPersonInfoHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "baseStaff/person/person.proto",
}
