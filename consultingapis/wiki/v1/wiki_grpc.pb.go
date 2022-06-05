// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wikiv1

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

// FAQServiceClient is the client API for FAQService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FAQServiceClient interface {
	GetDriveDocContent(ctx context.Context, in *DocTokenRequest, opts ...grpc.CallOption) (*DocTokenResp, error)
}

type fAQServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFAQServiceClient(cc grpc.ClientConnInterface) FAQServiceClient {
	return &fAQServiceClient{cc}
}

func (c *fAQServiceClient) GetDriveDocContent(ctx context.Context, in *DocTokenRequest, opts ...grpc.CallOption) (*DocTokenResp, error) {
	out := new(DocTokenResp)
	err := c.cc.Invoke(ctx, "/consultingapis.faq.v1.FAQService/GetDriveDocContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FAQServiceServer is the server API for FAQService service.
// All implementations must embed UnimplementedFAQServiceServer
// for forward compatibility
type FAQServiceServer interface {
	GetDriveDocContent(context.Context, *DocTokenRequest) (*DocTokenResp, error)
	mustEmbedUnimplementedFAQServiceServer()
}

// UnimplementedFAQServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFAQServiceServer struct {
}

func (UnimplementedFAQServiceServer) GetDriveDocContent(context.Context, *DocTokenRequest) (*DocTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDriveDocContent not implemented")
}
func (UnimplementedFAQServiceServer) mustEmbedUnimplementedFAQServiceServer() {}

// UnsafeFAQServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FAQServiceServer will
// result in compilation errors.
type UnsafeFAQServiceServer interface {
	mustEmbedUnimplementedFAQServiceServer()
}

func RegisterFAQServiceServer(s grpc.ServiceRegistrar, srv FAQServiceServer) {
	s.RegisterService(&FAQService_ServiceDesc, srv)
}

func _FAQService_GetDriveDocContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FAQServiceServer).GetDriveDocContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consultingapis.faq.v1.FAQService/GetDriveDocContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FAQServiceServer).GetDriveDocContent(ctx, req.(*DocTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FAQService_ServiceDesc is the grpc.ServiceDesc for FAQService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FAQService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "consultingapis.faq.v1.FAQService",
	HandlerType: (*FAQServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDriveDocContent",
			Handler:    _FAQService_GetDriveDocContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consultingapis/wiki/v1/wiki.proto",
}
