// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: campusapis/card/v1/card.proto

package cardv1

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

// CardServiceClient is the client API for CardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardServiceClient interface {
	//查询校园卡信息
	GetCardInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCardInfoResponse, error)
	//查询校园卡余额
	GetCardBalance(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCardBalanceResponse, error)
	//查询校园卡消费记录
	GetCardConsume(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCardConsumeResponse, error)
}

type cardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardServiceClient(cc grpc.ClientConnInterface) CardServiceClient {
	return &cardServiceClient{cc}
}

func (c *cardServiceClient) GetCardInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCardInfoResponse, error) {
	out := new(GetCardInfoResponse)
	err := c.cc.Invoke(ctx, "/campusapis.card.v1.CardService/GetCardInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) GetCardBalance(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCardBalanceResponse, error) {
	out := new(GetCardBalanceResponse)
	err := c.cc.Invoke(ctx, "/campusapis.card.v1.CardService/GetCardBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) GetCardConsume(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCardConsumeResponse, error) {
	out := new(GetCardConsumeResponse)
	err := c.cc.Invoke(ctx, "/campusapis.card.v1.CardService/GetCardConsume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardServiceServer is the server API for CardService service.
// All implementations must embed UnimplementedCardServiceServer
// for forward compatibility
type CardServiceServer interface {
	//查询校园卡信息
	GetCardInfo(context.Context, *emptypb.Empty) (*GetCardInfoResponse, error)
	//查询校园卡余额
	GetCardBalance(context.Context, *emptypb.Empty) (*GetCardBalanceResponse, error)
	//查询校园卡消费记录
	GetCardConsume(context.Context, *emptypb.Empty) (*GetCardConsumeResponse, error)
	mustEmbedUnimplementedCardServiceServer()
}

// UnimplementedCardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCardServiceServer struct {
}

func (UnimplementedCardServiceServer) GetCardInfo(context.Context, *emptypb.Empty) (*GetCardInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardInfo not implemented")
}
func (UnimplementedCardServiceServer) GetCardBalance(context.Context, *emptypb.Empty) (*GetCardBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardBalance not implemented")
}
func (UnimplementedCardServiceServer) GetCardConsume(context.Context, *emptypb.Empty) (*GetCardConsumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardConsume not implemented")
}
func (UnimplementedCardServiceServer) mustEmbedUnimplementedCardServiceServer() {}

// UnsafeCardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardServiceServer will
// result in compilation errors.
type UnsafeCardServiceServer interface {
	mustEmbedUnimplementedCardServiceServer()
}

func RegisterCardServiceServer(s grpc.ServiceRegistrar, srv CardServiceServer) {
	s.RegisterService(&CardService_ServiceDesc, srv)
}

func _CardService_GetCardInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).GetCardInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.card.v1.CardService/GetCardInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).GetCardInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_GetCardBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).GetCardBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.card.v1.CardService/GetCardBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).GetCardBalance(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_GetCardConsume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).GetCardConsume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/campusapis.card.v1.CardService/GetCardConsume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).GetCardConsume(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CardService_ServiceDesc is the grpc.ServiceDesc for CardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "campusapis.card.v1.CardService",
	HandlerType: (*CardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCardInfo",
			Handler:    _CardService_GetCardInfo_Handler,
		},
		{
			MethodName: "GetCardBalance",
			Handler:    _CardService_GetCardBalance_Handler,
		},
		{
			MethodName: "GetCardConsume",
			Handler:    _CardService_GetCardConsume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "campusapis/card/v1/card.proto",
}
