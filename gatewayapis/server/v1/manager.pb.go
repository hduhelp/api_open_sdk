// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gatewayapis/server/v1/manager.proto

package serverv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetHTTPRouterListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHTTPRouterListRequest) Reset() {
	*x = GetHTTPRouterListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHTTPRouterListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHTTPRouterListRequest) ProtoMessage() {}

func (x *GetHTTPRouterListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHTTPRouterListRequest.ProtoReflect.Descriptor instead.
func (*GetHTTPRouterListRequest) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_manager_proto_rawDescGZIP(), []int{0}
}

type GetHTTPRouterListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routers []*HTTPRouter `protobuf:"bytes,1,rep,name=routers,proto3" json:"routers,omitempty"`
}

func (x *GetHTTPRouterListResponse) Reset() {
	*x = GetHTTPRouterListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHTTPRouterListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHTTPRouterListResponse) ProtoMessage() {}

func (x *GetHTTPRouterListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHTTPRouterListResponse.ProtoReflect.Descriptor instead.
func (*GetHTTPRouterListResponse) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_manager_proto_rawDescGZIP(), []int{1}
}

func (x *GetHTTPRouterListResponse) GetRouters() []*HTTPRouter {
	if x != nil {
		return x.Routers
	}
	return nil
}

type GetHTTPRouterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//服务名称
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	//方法
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	//挂载地址
	ServePath string `protobuf:"bytes,3,opt,name=servePath,proto3" json:"servePath,omitempty"`
}

func (x *GetHTTPRouterRequest) Reset() {
	*x = GetHTTPRouterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHTTPRouterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHTTPRouterRequest) ProtoMessage() {}

func (x *GetHTTPRouterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHTTPRouterRequest.ProtoReflect.Descriptor instead.
func (*GetHTTPRouterRequest) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_manager_proto_rawDescGZIP(), []int{2}
}

func (x *GetHTTPRouterRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *GetHTTPRouterRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *GetHTTPRouterRequest) GetServePath() string {
	if x != nil {
		return x.ServePath
	}
	return ""
}

type GetHTTPRouterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error int32       `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg   string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data  *HTTPRouter `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetHTTPRouterResponse) Reset() {
	*x = GetHTTPRouterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHTTPRouterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHTTPRouterResponse) ProtoMessage() {}

func (x *GetHTTPRouterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHTTPRouterResponse.ProtoReflect.Descriptor instead.
func (*GetHTTPRouterResponse) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_manager_proto_rawDescGZIP(), []int{3}
}

func (x *GetHTTPRouterResponse) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *GetHTTPRouterResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetHTTPRouterResponse) GetData() *HTTPRouter {
	if x != nil {
		return x.Data
	}
	return nil
}

type PostOrPutHTTPRouterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Router *HTTPRouter `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
}

func (x *PostOrPutHTTPRouterRequest) Reset() {
	*x = PostOrPutHTTPRouterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostOrPutHTTPRouterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostOrPutHTTPRouterRequest) ProtoMessage() {}

func (x *PostOrPutHTTPRouterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostOrPutHTTPRouterRequest.ProtoReflect.Descriptor instead.
func (*PostOrPutHTTPRouterRequest) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_manager_proto_rawDescGZIP(), []int{4}
}

func (x *PostOrPutHTTPRouterRequest) GetRouter() *HTTPRouter {
	if x != nil {
		return x.Router
	}
	return nil
}

type PostOrPutHTTPRouterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostOrPutHTTPRouterResponse) Reset() {
	*x = PostOrPutHTTPRouterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostOrPutHTTPRouterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostOrPutHTTPRouterResponse) ProtoMessage() {}

func (x *PostOrPutHTTPRouterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostOrPutHTTPRouterResponse.ProtoReflect.Descriptor instead.
func (*PostOrPutHTTPRouterResponse) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_manager_proto_rawDescGZIP(), []int{5}
}

type DeleteHTTPRouterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//服务名称
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	//方法
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	//挂载地址
	ServePath string `protobuf:"bytes,3,opt,name=servePath,proto3" json:"servePath,omitempty"`
}

func (x *DeleteHTTPRouterRequest) Reset() {
	*x = DeleteHTTPRouterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHTTPRouterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHTTPRouterRequest) ProtoMessage() {}

func (x *DeleteHTTPRouterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHTTPRouterRequest.ProtoReflect.Descriptor instead.
func (*DeleteHTTPRouterRequest) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_manager_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteHTTPRouterRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *DeleteHTTPRouterRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *DeleteHTTPRouterRequest) GetServePath() string {
	if x != nil {
		return x.ServePath
	}
	return ""
}

type DeleteHTTPRouterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteHTTPRouterResponse) Reset() {
	*x = DeleteHTTPRouterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHTTPRouterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHTTPRouterResponse) ProtoMessage() {}

func (x *DeleteHTTPRouterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHTTPRouterResponse.ProtoReflect.Descriptor instead.
func (*DeleteHTTPRouterResponse) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_manager_proto_rawDescGZIP(), []int{7}
}

var File_gatewayapis_server_v1_manager_proto protoreflect.FileDescriptor

var file_gatewayapis_server_v1_manager_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x07,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x22, 0x66, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x48, 0x54,
	0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x76, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x57, 0x0a, 0x1a, 0x50, 0x6f, 0x73, 0x74, 0x4f,
	0x72, 0x50, 0x75, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x54,
	0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x22, 0x1d, 0x0a, 0x1b, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x72, 0x50, 0x75, 0x74, 0x48, 0x54, 0x54,
	0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x69, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf6, 0x03, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x78, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7e, 0x0a, 0x13, 0x50, 0x6f, 0x73, 0x74,
	0x4f, 0x72, 0x50, 0x75, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12,
	0x31, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x72, 0x50, 0x75,
	0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4f,
	0x72, 0x50, 0x75, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x75, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0xdf, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68, 0x65, 0x6c,
	0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x47, 0x53, 0x58, 0xaa, 0x02, 0x15, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69,
	0x73, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x61, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gatewayapis_server_v1_manager_proto_rawDescOnce sync.Once
	file_gatewayapis_server_v1_manager_proto_rawDescData = file_gatewayapis_server_v1_manager_proto_rawDesc
)

func file_gatewayapis_server_v1_manager_proto_rawDescGZIP() []byte {
	file_gatewayapis_server_v1_manager_proto_rawDescOnce.Do(func() {
		file_gatewayapis_server_v1_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_gatewayapis_server_v1_manager_proto_rawDescData)
	})
	return file_gatewayapis_server_v1_manager_proto_rawDescData
}

var file_gatewayapis_server_v1_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_gatewayapis_server_v1_manager_proto_goTypes = []interface{}{
	(*GetHTTPRouterListRequest)(nil),    // 0: gatewayapis.server.v1.GetHTTPRouterListRequest
	(*GetHTTPRouterListResponse)(nil),   // 1: gatewayapis.server.v1.GetHTTPRouterListResponse
	(*GetHTTPRouterRequest)(nil),        // 2: gatewayapis.server.v1.GetHTTPRouterRequest
	(*GetHTTPRouterResponse)(nil),       // 3: gatewayapis.server.v1.GetHTTPRouterResponse
	(*PostOrPutHTTPRouterRequest)(nil),  // 4: gatewayapis.server.v1.PostOrPutHTTPRouterRequest
	(*PostOrPutHTTPRouterResponse)(nil), // 5: gatewayapis.server.v1.PostOrPutHTTPRouterResponse
	(*DeleteHTTPRouterRequest)(nil),     // 6: gatewayapis.server.v1.DeleteHTTPRouterRequest
	(*DeleteHTTPRouterResponse)(nil),    // 7: gatewayapis.server.v1.DeleteHTTPRouterResponse
	(*HTTPRouter)(nil),                  // 8: gatewayapis.server.v1.HTTPRouter
}
var file_gatewayapis_server_v1_manager_proto_depIdxs = []int32{
	8, // 0: gatewayapis.server.v1.GetHTTPRouterListResponse.routers:type_name -> gatewayapis.server.v1.HTTPRouter
	8, // 1: gatewayapis.server.v1.GetHTTPRouterResponse.data:type_name -> gatewayapis.server.v1.HTTPRouter
	8, // 2: gatewayapis.server.v1.PostOrPutHTTPRouterRequest.router:type_name -> gatewayapis.server.v1.HTTPRouter
	0, // 3: gatewayapis.server.v1.ServerManageerService.GetHTTPRouterList:input_type -> gatewayapis.server.v1.GetHTTPRouterListRequest
	2, // 4: gatewayapis.server.v1.ServerManageerService.GetHTTPRouter:input_type -> gatewayapis.server.v1.GetHTTPRouterRequest
	4, // 5: gatewayapis.server.v1.ServerManageerService.PostOrPutHTTPRouter:input_type -> gatewayapis.server.v1.PostOrPutHTTPRouterRequest
	6, // 6: gatewayapis.server.v1.ServerManageerService.DeleteHTTPRouter:input_type -> gatewayapis.server.v1.DeleteHTTPRouterRequest
	1, // 7: gatewayapis.server.v1.ServerManageerService.GetHTTPRouterList:output_type -> gatewayapis.server.v1.GetHTTPRouterListResponse
	3, // 8: gatewayapis.server.v1.ServerManageerService.GetHTTPRouter:output_type -> gatewayapis.server.v1.GetHTTPRouterResponse
	5, // 9: gatewayapis.server.v1.ServerManageerService.PostOrPutHTTPRouter:output_type -> gatewayapis.server.v1.PostOrPutHTTPRouterResponse
	7, // 10: gatewayapis.server.v1.ServerManageerService.DeleteHTTPRouter:output_type -> gatewayapis.server.v1.DeleteHTTPRouterResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gatewayapis_server_v1_manager_proto_init() }
func file_gatewayapis_server_v1_manager_proto_init() {
	if File_gatewayapis_server_v1_manager_proto != nil {
		return
	}
	file_gatewayapis_server_v1_defined_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gatewayapis_server_v1_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHTTPRouterListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gatewayapis_server_v1_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHTTPRouterListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gatewayapis_server_v1_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHTTPRouterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gatewayapis_server_v1_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHTTPRouterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gatewayapis_server_v1_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostOrPutHTTPRouterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gatewayapis_server_v1_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostOrPutHTTPRouterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gatewayapis_server_v1_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHTTPRouterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gatewayapis_server_v1_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHTTPRouterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gatewayapis_server_v1_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gatewayapis_server_v1_manager_proto_goTypes,
		DependencyIndexes: file_gatewayapis_server_v1_manager_proto_depIdxs,
		MessageInfos:      file_gatewayapis_server_v1_manager_proto_msgTypes,
	}.Build()
	File_gatewayapis_server_v1_manager_proto = out.File
	file_gatewayapis_server_v1_manager_proto_rawDesc = nil
	file_gatewayapis_server_v1_manager_proto_goTypes = nil
	file_gatewayapis_server_v1_manager_proto_depIdxs = nil
}
