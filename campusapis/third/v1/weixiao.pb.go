// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: campusapis/third/v1/weixiao.proto

package thirdv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommonWeixiaoPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawData string `protobuf:"bytes,1,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
	AppKey  string `protobuf:"bytes,2,opt,name=app_key,json=appKey,proto3" json:"app_key,omitempty"`
}

func (x *CommonWeixiaoPostRequest) Reset() {
	*x = CommonWeixiaoPostRequest{}
	mi := &file_campusapis_third_v1_weixiao_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommonWeixiaoPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonWeixiaoPostRequest) ProtoMessage() {}

func (x *CommonWeixiaoPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_third_v1_weixiao_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonWeixiaoPostRequest.ProtoReflect.Descriptor instead.
func (*CommonWeixiaoPostRequest) Descriptor() ([]byte, []int) {
	return file_campusapis_third_v1_weixiao_proto_rawDescGZIP(), []int{0}
}

func (x *CommonWeixiaoPostRequest) GetRawData() string {
	if x != nil {
		return x.RawData
	}
	return ""
}

func (x *CommonWeixiaoPostRequest) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

type CommonWeixiaoPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	RawData string `protobuf:"bytes,3,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
	AppKey  string `protobuf:"bytes,4,opt,name=app_key,json=appKey,proto3" json:"app_key,omitempty"`
}

func (x *CommonWeixiaoPostResponse) Reset() {
	*x = CommonWeixiaoPostResponse{}
	mi := &file_campusapis_third_v1_weixiao_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommonWeixiaoPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonWeixiaoPostResponse) ProtoMessage() {}

func (x *CommonWeixiaoPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_third_v1_weixiao_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonWeixiaoPostResponse.ProtoReflect.Descriptor instead.
func (*CommonWeixiaoPostResponse) Descriptor() ([]byte, []int) {
	return file_campusapis_third_v1_weixiao_proto_rawDescGZIP(), []int{1}
}

func (x *CommonWeixiaoPostResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonWeixiaoPostResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CommonWeixiaoPostResponse) GetRawData() string {
	if x != nil {
		return x.RawData
	}
	return ""
}

func (x *CommonWeixiaoPostResponse) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

var File_campusapis_third_v1_weixiao_proto protoreflect.FileDescriptor

var file_campusapis_third_v1_weixiao_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x69, 0x78, 0x69, 0x61, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x57, 0x65, 0x69,
	0x78, 0x69, 0x61, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70,
	0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70,
	0x4b, 0x65, 0x79, 0x22, 0x7d, 0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x57, 0x65, 0x69,
	0x78, 0x69, 0x61, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x70,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x4b,
	0x65, 0x79, 0x32, 0xc0, 0x06, 0x0a, 0x0e, 0x57, 0x65, 0x69, 0x78, 0x69, 0x61, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa2, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x43, 0x61, 0x72, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2d,
	0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x57, 0x65, 0x69, 0x78, 0x69,
	0x61, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x57, 0x65, 0x69, 0x78, 0x69, 0x61,
	0x6f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x69, 0x78, 0x69, 0x61, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x12,
	0x2d, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x57, 0x65, 0x69, 0x78,
	0x69, 0x61, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x57, 0x65, 0x69, 0x78, 0x69,
	0x61, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x69, 0x78, 0x69, 0x61, 0x6f, 0x2f, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x12, 0x9d, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12,
	0x2d, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x57, 0x65, 0x69, 0x78,
	0x69, 0x61, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x57, 0x65, 0x69, 0x78, 0x69,
	0x61, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x69, 0x78, 0x69, 0x61, 0x6f, 0x2f, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0xa7, 0x01, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x57, 0x65, 0x69, 0x78, 0x69, 0x61, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x57, 0x65, 0x69, 0x78, 0x69, 0x61, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24,
	0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x69, 0x78, 0x69, 0x61,
	0x6f, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x2d, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x57, 0x65, 0x69, 0x78, 0x69, 0x61, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x57, 0x65, 0x69, 0x78, 0x69, 0x61, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01,
	0x2a, 0x22, 0x22, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x69,
	0x78, 0x69, 0x61, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x45, 0x78, 0x61, 0x6d, 0x42, 0xd2, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x61,
	0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x57, 0x65, 0x69, 0x78, 0x69, 0x61, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64,
	0x75, 0x68, 0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73,
	0x64, 0x6b, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68, 0x69, 0x72, 0x64, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x54, 0x58, 0xaa, 0x02, 0x13, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x54, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x43, 0x61, 0x6d,
	0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x54, 0x68, 0x69, 0x72, 0x64, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1f, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x54, 0x68,
	0x69, 0x72, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x15, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x3a,
	0x3a, 0x54, 0x68, 0x69, 0x72, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_campusapis_third_v1_weixiao_proto_rawDescOnce sync.Once
	file_campusapis_third_v1_weixiao_proto_rawDescData = file_campusapis_third_v1_weixiao_proto_rawDesc
)

func file_campusapis_third_v1_weixiao_proto_rawDescGZIP() []byte {
	file_campusapis_third_v1_weixiao_proto_rawDescOnce.Do(func() {
		file_campusapis_third_v1_weixiao_proto_rawDescData = protoimpl.X.CompressGZIP(file_campusapis_third_v1_weixiao_proto_rawDescData)
	})
	return file_campusapis_third_v1_weixiao_proto_rawDescData
}

var file_campusapis_third_v1_weixiao_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_campusapis_third_v1_weixiao_proto_goTypes = []any{
	(*CommonWeixiaoPostRequest)(nil),  // 0: campusapis.third.v1.CommonWeixiaoPostRequest
	(*CommonWeixiaoPostResponse)(nil), // 1: campusapis.third.v1.CommonWeixiaoPostResponse
}
var file_campusapis_third_v1_weixiao_proto_depIdxs = []int32{
	0, // 0: campusapis.third.v1.WeixiaoService.GetSchoolCardBalance:input_type -> campusapis.third.v1.CommonWeixiaoPostRequest
	0, // 1: campusapis.third.v1.WeixiaoService.GetLibraryBorrow:input_type -> campusapis.third.v1.CommonWeixiaoPostRequest
	0, // 2: campusapis.third.v1.WeixiaoService.GetStudentCourse:input_type -> campusapis.third.v1.CommonWeixiaoPostRequest
	0, // 3: campusapis.third.v1.WeixiaoService.GetStudentCourseScore:input_type -> campusapis.third.v1.CommonWeixiaoPostRequest
	0, // 4: campusapis.third.v1.WeixiaoService.GetStudentExam:input_type -> campusapis.third.v1.CommonWeixiaoPostRequest
	1, // 5: campusapis.third.v1.WeixiaoService.GetSchoolCardBalance:output_type -> campusapis.third.v1.CommonWeixiaoPostResponse
	1, // 6: campusapis.third.v1.WeixiaoService.GetLibraryBorrow:output_type -> campusapis.third.v1.CommonWeixiaoPostResponse
	1, // 7: campusapis.third.v1.WeixiaoService.GetStudentCourse:output_type -> campusapis.third.v1.CommonWeixiaoPostResponse
	1, // 8: campusapis.third.v1.WeixiaoService.GetStudentCourseScore:output_type -> campusapis.third.v1.CommonWeixiaoPostResponse
	1, // 9: campusapis.third.v1.WeixiaoService.GetStudentExam:output_type -> campusapis.third.v1.CommonWeixiaoPostResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_campusapis_third_v1_weixiao_proto_init() }
func file_campusapis_third_v1_weixiao_proto_init() {
	if File_campusapis_third_v1_weixiao_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_campusapis_third_v1_weixiao_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_campusapis_third_v1_weixiao_proto_goTypes,
		DependencyIndexes: file_campusapis_third_v1_weixiao_proto_depIdxs,
		MessageInfos:      file_campusapis_third_v1_weixiao_proto_msgTypes,
	}.Build()
	File_campusapis_third_v1_weixiao_proto = out.File
	file_campusapis_third_v1_weixiao_proto_rawDesc = nil
	file_campusapis_third_v1_weixiao_proto_goTypes = nil
	file_campusapis_third_v1_weixiao_proto_depIdxs = nil
}
