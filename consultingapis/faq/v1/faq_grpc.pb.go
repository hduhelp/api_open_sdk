// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: consultingapis/faq/v1/faq.proto

package faqv1

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
	CreateFAQ(ctx context.Context, in *CreateFAQRequest, opts ...grpc.CallOption) (*CreateFAQResponse, error)
	GetFAQ(ctx context.Context, in *GetFAQRequest, opts ...grpc.CallOption) (*GetFAQResponse, error)
	UpdateFAQ(ctx context.Context, in *UpdateFAQRequest, opts ...grpc.CallOption) (*UpdateFAQResponse, error)
	DeleteFAQ(ctx context.Context, in *DeleteFAQRequest, opts ...grpc.CallOption) (*DeleteFAQResponse, error)
	GetFAQImage(ctx context.Context, in *GetFAQImageRequest, opts ...grpc.CallOption) (*GetFAQImageResponse, error)
	SearchFAQAnswerRichtext(ctx context.Context, in *SearchFAQAnswerRichtextRequest, opts ...grpc.CallOption) (*SearchFAQAnswerRichtextResponse, error)
}

type fAQServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFAQServiceClient(cc grpc.ClientConnInterface) FAQServiceClient {
	return &fAQServiceClient{cc}
}

func (c *fAQServiceClient) CreateFAQ(ctx context.Context, in *CreateFAQRequest, opts ...grpc.CallOption) (*CreateFAQResponse, error) {
	out := new(CreateFAQResponse)
	err := c.cc.Invoke(ctx, "/consultingapis.faq.v1.FAQService/CreateFAQ", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fAQServiceClient) GetFAQ(ctx context.Context, in *GetFAQRequest, opts ...grpc.CallOption) (*GetFAQResponse, error) {
	out := new(GetFAQResponse)
	err := c.cc.Invoke(ctx, "/consultingapis.faq.v1.FAQService/GetFAQ", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fAQServiceClient) UpdateFAQ(ctx context.Context, in *UpdateFAQRequest, opts ...grpc.CallOption) (*UpdateFAQResponse, error) {
	out := new(UpdateFAQResponse)
	err := c.cc.Invoke(ctx, "/consultingapis.faq.v1.FAQService/UpdateFAQ", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fAQServiceClient) DeleteFAQ(ctx context.Context, in *DeleteFAQRequest, opts ...grpc.CallOption) (*DeleteFAQResponse, error) {
	out := new(DeleteFAQResponse)
	err := c.cc.Invoke(ctx, "/consultingapis.faq.v1.FAQService/DeleteFAQ", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fAQServiceClient) GetFAQImage(ctx context.Context, in *GetFAQImageRequest, opts ...grpc.CallOption) (*GetFAQImageResponse, error) {
	out := new(GetFAQImageResponse)
	err := c.cc.Invoke(ctx, "/consultingapis.faq.v1.FAQService/GetFAQImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fAQServiceClient) SearchFAQAnswerRichtext(ctx context.Context, in *SearchFAQAnswerRichtextRequest, opts ...grpc.CallOption) (*SearchFAQAnswerRichtextResponse, error) {
	out := new(SearchFAQAnswerRichtextResponse)
	err := c.cc.Invoke(ctx, "/consultingapis.faq.v1.FAQService/SearchFAQAnswerRichtext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FAQServiceServer is the server API for FAQService service.
// All implementations must embed UnimplementedFAQServiceServer
// for forward compatibility
type FAQServiceServer interface {
	CreateFAQ(context.Context, *CreateFAQRequest) (*CreateFAQResponse, error)
	GetFAQ(context.Context, *GetFAQRequest) (*GetFAQResponse, error)
	UpdateFAQ(context.Context, *UpdateFAQRequest) (*UpdateFAQResponse, error)
	DeleteFAQ(context.Context, *DeleteFAQRequest) (*DeleteFAQResponse, error)
	GetFAQImage(context.Context, *GetFAQImageRequest) (*GetFAQImageResponse, error)
	SearchFAQAnswerRichtext(context.Context, *SearchFAQAnswerRichtextRequest) (*SearchFAQAnswerRichtextResponse, error)
	mustEmbedUnimplementedFAQServiceServer()
}

// UnimplementedFAQServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFAQServiceServer struct {
}

func (UnimplementedFAQServiceServer) CreateFAQ(context.Context, *CreateFAQRequest) (*CreateFAQResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFAQ not implemented")
}
func (UnimplementedFAQServiceServer) GetFAQ(context.Context, *GetFAQRequest) (*GetFAQResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFAQ not implemented")
}
func (UnimplementedFAQServiceServer) UpdateFAQ(context.Context, *UpdateFAQRequest) (*UpdateFAQResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFAQ not implemented")
}
func (UnimplementedFAQServiceServer) DeleteFAQ(context.Context, *DeleteFAQRequest) (*DeleteFAQResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFAQ not implemented")
}
func (UnimplementedFAQServiceServer) GetFAQImage(context.Context, *GetFAQImageRequest) (*GetFAQImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFAQImage not implemented")
}
func (UnimplementedFAQServiceServer) SearchFAQAnswerRichtext(context.Context, *SearchFAQAnswerRichtextRequest) (*SearchFAQAnswerRichtextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFAQAnswerRichtext not implemented")
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

func _FAQService_CreateFAQ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFAQRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FAQServiceServer).CreateFAQ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consultingapis.faq.v1.FAQService/CreateFAQ",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FAQServiceServer).CreateFAQ(ctx, req.(*CreateFAQRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FAQService_GetFAQ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFAQRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FAQServiceServer).GetFAQ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consultingapis.faq.v1.FAQService/GetFAQ",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FAQServiceServer).GetFAQ(ctx, req.(*GetFAQRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FAQService_UpdateFAQ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFAQRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FAQServiceServer).UpdateFAQ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consultingapis.faq.v1.FAQService/UpdateFAQ",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FAQServiceServer).UpdateFAQ(ctx, req.(*UpdateFAQRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FAQService_DeleteFAQ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFAQRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FAQServiceServer).DeleteFAQ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consultingapis.faq.v1.FAQService/DeleteFAQ",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FAQServiceServer).DeleteFAQ(ctx, req.(*DeleteFAQRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FAQService_GetFAQImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFAQImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FAQServiceServer).GetFAQImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consultingapis.faq.v1.FAQService/GetFAQImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FAQServiceServer).GetFAQImage(ctx, req.(*GetFAQImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FAQService_SearchFAQAnswerRichtext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFAQAnswerRichtextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FAQServiceServer).SearchFAQAnswerRichtext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consultingapis.faq.v1.FAQService/SearchFAQAnswerRichtext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FAQServiceServer).SearchFAQAnswerRichtext(ctx, req.(*SearchFAQAnswerRichtextRequest))
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
			MethodName: "CreateFAQ",
			Handler:    _FAQService_CreateFAQ_Handler,
		},
		{
			MethodName: "GetFAQ",
			Handler:    _FAQService_GetFAQ_Handler,
		},
		{
			MethodName: "UpdateFAQ",
			Handler:    _FAQService_UpdateFAQ_Handler,
		},
		{
			MethodName: "DeleteFAQ",
			Handler:    _FAQService_DeleteFAQ_Handler,
		},
		{
			MethodName: "GetFAQImage",
			Handler:    _FAQService_GetFAQImage_Handler,
		},
		{
			MethodName: "SearchFAQAnswerRichtext",
			Handler:    _FAQService_SearchFAQAnswerRichtext_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consultingapis/faq/v1/faq.proto",
}
