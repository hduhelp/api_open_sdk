// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: campusapis/library/v1/library.proto

package libraryv1

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

const (
	LibraryService_GetBookInfo_FullMethodName        = "/campusapis.library.v1.LibraryService/GetBookInfo"
	LibraryService_GetBookInfoList_FullMethodName    = "/campusapis.library.v1.LibraryService/GetBookInfoList"
	LibraryService_GetBookMARC_FullMethodName        = "/campusapis.library.v1.LibraryService/GetBookMARC"
	LibraryService_GetBookMARCList_FullMethodName    = "/campusapis.library.v1.LibraryService/GetBookMARCList"
	LibraryService_GetBookLendLast_FullMethodName    = "/campusapis.library.v1.LibraryService/GetBookLendLast"
	LibraryService_GetBookLendHistory_FullMethodName = "/campusapis.library.v1.LibraryService/GetBookLendHistory"
)

// LibraryServiceClient is the client API for LibraryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibraryServiceClient interface {
	// 查询书本信息
	GetBookInfo(ctx context.Context, in *GetBookInfoRequest, opts ...grpc.CallOption) (*GetBookInfoResponse, error)
	// 查询书本信息列表，只能查询一个书籍的书本信息
	GetBookInfoList(ctx context.Context, in *GetBookInfoListRequest, opts ...grpc.CallOption) (*GetBookInfoListResponse, error)
	// 查询书籍信息
	GetBookMARC(ctx context.Context, in *GetBookMARCRequest, opts ...grpc.CallOption) (*GetBookMARCResponse, error)
	// 查询书籍信息列表
	GetBookMARCList(ctx context.Context, in *GetBookMARCListRequest, opts ...grpc.CallOption) (*GetBookMARCListResponse, error)
	// 查询当前图书借阅记录
	GetBookLendLast(ctx context.Context, in *GetBookLendLastRequest, opts ...grpc.CallOption) (*GetBookLendLastResponse, error)
	// 查询历史图书借阅记录
	GetBookLendHistory(ctx context.Context, in *GetBookLendHistoryRequest, opts ...grpc.CallOption) (*GetBookLendHistoryResponse, error)
}

type libraryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryServiceClient(cc grpc.ClientConnInterface) LibraryServiceClient {
	return &libraryServiceClient{cc}
}

func (c *libraryServiceClient) GetBookInfo(ctx context.Context, in *GetBookInfoRequest, opts ...grpc.CallOption) (*GetBookInfoResponse, error) {
	out := new(GetBookInfoResponse)
	err := c.cc.Invoke(ctx, LibraryService_GetBookInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetBookInfoList(ctx context.Context, in *GetBookInfoListRequest, opts ...grpc.CallOption) (*GetBookInfoListResponse, error) {
	out := new(GetBookInfoListResponse)
	err := c.cc.Invoke(ctx, LibraryService_GetBookInfoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetBookMARC(ctx context.Context, in *GetBookMARCRequest, opts ...grpc.CallOption) (*GetBookMARCResponse, error) {
	out := new(GetBookMARCResponse)
	err := c.cc.Invoke(ctx, LibraryService_GetBookMARC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetBookMARCList(ctx context.Context, in *GetBookMARCListRequest, opts ...grpc.CallOption) (*GetBookMARCListResponse, error) {
	out := new(GetBookMARCListResponse)
	err := c.cc.Invoke(ctx, LibraryService_GetBookMARCList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetBookLendLast(ctx context.Context, in *GetBookLendLastRequest, opts ...grpc.CallOption) (*GetBookLendLastResponse, error) {
	out := new(GetBookLendLastResponse)
	err := c.cc.Invoke(ctx, LibraryService_GetBookLendLast_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetBookLendHistory(ctx context.Context, in *GetBookLendHistoryRequest, opts ...grpc.CallOption) (*GetBookLendHistoryResponse, error) {
	out := new(GetBookLendHistoryResponse)
	err := c.cc.Invoke(ctx, LibraryService_GetBookLendHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServiceServer is the server API for LibraryService service.
// All implementations must embed UnimplementedLibraryServiceServer
// for forward compatibility
type LibraryServiceServer interface {
	// 查询书本信息
	GetBookInfo(context.Context, *GetBookInfoRequest) (*GetBookInfoResponse, error)
	// 查询书本信息列表，只能查询一个书籍的书本信息
	GetBookInfoList(context.Context, *GetBookInfoListRequest) (*GetBookInfoListResponse, error)
	// 查询书籍信息
	GetBookMARC(context.Context, *GetBookMARCRequest) (*GetBookMARCResponse, error)
	// 查询书籍信息列表
	GetBookMARCList(context.Context, *GetBookMARCListRequest) (*GetBookMARCListResponse, error)
	// 查询当前图书借阅记录
	GetBookLendLast(context.Context, *GetBookLendLastRequest) (*GetBookLendLastResponse, error)
	// 查询历史图书借阅记录
	GetBookLendHistory(context.Context, *GetBookLendHistoryRequest) (*GetBookLendHistoryResponse, error)
	mustEmbedUnimplementedLibraryServiceServer()
}

// UnimplementedLibraryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLibraryServiceServer struct {
}

func (UnimplementedLibraryServiceServer) GetBookInfo(context.Context, *GetBookInfoRequest) (*GetBookInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookInfo not implemented")
}
func (UnimplementedLibraryServiceServer) GetBookInfoList(context.Context, *GetBookInfoListRequest) (*GetBookInfoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookInfoList not implemented")
}
func (UnimplementedLibraryServiceServer) GetBookMARC(context.Context, *GetBookMARCRequest) (*GetBookMARCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookMARC not implemented")
}
func (UnimplementedLibraryServiceServer) GetBookMARCList(context.Context, *GetBookMARCListRequest) (*GetBookMARCListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookMARCList not implemented")
}
func (UnimplementedLibraryServiceServer) GetBookLendLast(context.Context, *GetBookLendLastRequest) (*GetBookLendLastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookLendLast not implemented")
}
func (UnimplementedLibraryServiceServer) GetBookLendHistory(context.Context, *GetBookLendHistoryRequest) (*GetBookLendHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookLendHistory not implemented")
}
func (UnimplementedLibraryServiceServer) mustEmbedUnimplementedLibraryServiceServer() {}

// UnsafeLibraryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServiceServer will
// result in compilation errors.
type UnsafeLibraryServiceServer interface {
	mustEmbedUnimplementedLibraryServiceServer()
}

func RegisterLibraryServiceServer(s grpc.ServiceRegistrar, srv LibraryServiceServer) {
	s.RegisterService(&LibraryService_ServiceDesc, srv)
}

func _LibraryService_GetBookInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetBookInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_GetBookInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetBookInfo(ctx, req.(*GetBookInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetBookInfoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookInfoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetBookInfoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_GetBookInfoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetBookInfoList(ctx, req.(*GetBookInfoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetBookMARC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookMARCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetBookMARC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_GetBookMARC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetBookMARC(ctx, req.(*GetBookMARCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetBookMARCList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookMARCListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetBookMARCList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_GetBookMARCList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetBookMARCList(ctx, req.(*GetBookMARCListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetBookLendLast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookLendLastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetBookLendLast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_GetBookLendLast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetBookLendLast(ctx, req.(*GetBookLendLastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetBookLendHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookLendHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetBookLendHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_GetBookLendHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetBookLendHistory(ctx, req.(*GetBookLendHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LibraryService_ServiceDesc is the grpc.ServiceDesc for LibraryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LibraryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "campusapis.library.v1.LibraryService",
	HandlerType: (*LibraryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookInfo",
			Handler:    _LibraryService_GetBookInfo_Handler,
		},
		{
			MethodName: "GetBookInfoList",
			Handler:    _LibraryService_GetBookInfoList_Handler,
		},
		{
			MethodName: "GetBookMARC",
			Handler:    _LibraryService_GetBookMARC_Handler,
		},
		{
			MethodName: "GetBookMARCList",
			Handler:    _LibraryService_GetBookMARCList_Handler,
		},
		{
			MethodName: "GetBookLendLast",
			Handler:    _LibraryService_GetBookLendLast_Handler,
		},
		{
			MethodName: "GetBookLendHistory",
			Handler:    _LibraryService_GetBookLendHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "campusapis/library/v1/library.proto",
}
