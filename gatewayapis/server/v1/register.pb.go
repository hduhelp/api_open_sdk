// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: gatewayapis/server/v1/register.proto

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

type PostRegistServiceInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instance *ServiceInstance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *PostRegistServiceInstanceRequest) Reset() {
	*x = PostRegistServiceInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRegistServiceInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRegistServiceInstanceRequest) ProtoMessage() {}

func (x *PostRegistServiceInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRegistServiceInstanceRequest.ProtoReflect.Descriptor instead.
func (*PostRegistServiceInstanceRequest) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_register_proto_rawDescGZIP(), []int{0}
}

func (x *PostRegistServiceInstanceRequest) GetInstance() *ServiceInstance {
	if x != nil {
		return x.Instance
	}
	return nil
}

type PostRegistServiceInstanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostRegistServiceInstanceResponse) Reset() {
	*x = PostRegistServiceInstanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRegistServiceInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRegistServiceInstanceResponse) ProtoMessage() {}

func (x *PostRegistServiceInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRegistServiceInstanceResponse.ProtoReflect.Descriptor instead.
func (*PostRegistServiceInstanceResponse) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_register_proto_rawDescGZIP(), []int{1}
}

type PostRegistHTTPRoutersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routers []*HTTPRouter `protobuf:"bytes,1,rep,name=routers,proto3" json:"routers,omitempty"`
}

func (x *PostRegistHTTPRoutersRequest) Reset() {
	*x = PostRegistHTTPRoutersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_register_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRegistHTTPRoutersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRegistHTTPRoutersRequest) ProtoMessage() {}

func (x *PostRegistHTTPRoutersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_register_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRegistHTTPRoutersRequest.ProtoReflect.Descriptor instead.
func (*PostRegistHTTPRoutersRequest) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_register_proto_rawDescGZIP(), []int{2}
}

func (x *PostRegistHTTPRoutersRequest) GetRouters() []*HTTPRouter {
	if x != nil {
		return x.Routers
	}
	return nil
}

type PostRegistHTTPRoutersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostRegistHTTPRoutersResponse) Reset() {
	*x = PostRegistHTTPRoutersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_register_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRegistHTTPRoutersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRegistHTTPRoutersResponse) ProtoMessage() {}

func (x *PostRegistHTTPRoutersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_register_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRegistHTTPRoutersResponse.ProtoReflect.Descriptor instead.
func (*PostRegistHTTPRoutersResponse) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_register_proto_rawDescGZIP(), []int{3}
}

type PostRegistGRPCMethodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Methods []*GRPCMethod `protobuf:"bytes,1,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *PostRegistGRPCMethodsRequest) Reset() {
	*x = PostRegistGRPCMethodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_register_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRegistGRPCMethodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRegistGRPCMethodsRequest) ProtoMessage() {}

func (x *PostRegistGRPCMethodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_register_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRegistGRPCMethodsRequest.ProtoReflect.Descriptor instead.
func (*PostRegistGRPCMethodsRequest) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_register_proto_rawDescGZIP(), []int{4}
}

func (x *PostRegistGRPCMethodsRequest) GetMethods() []*GRPCMethod {
	if x != nil {
		return x.Methods
	}
	return nil
}

type PostRegistGRPCMethodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostRegistGRPCMethodsResponse) Reset() {
	*x = PostRegistGRPCMethodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_register_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRegistGRPCMethodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRegistGRPCMethodsResponse) ProtoMessage() {}

func (x *PostRegistGRPCMethodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_register_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRegistGRPCMethodsResponse.ProtoReflect.Descriptor instead.
func (*PostRegistGRPCMethodsResponse) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_register_proto_rawDescGZIP(), []int{5}
}

var File_gatewayapis_server_v1_register_proto protoreflect.FileDescriptor

var file_gatewayapis_server_v1_register_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x66, 0x0a, 0x20, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x23, 0x0a, 0x21, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x5b, 0x0a, 0x1c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x48, 0x54, 0x54,
	0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3b, 0x0a, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x22, 0x1f, 0x0a, 0x1d,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x0a,
	0x1c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x47, 0x52, 0x50, 0x43, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a,
	0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x22, 0x1f, 0x0a, 0x1d, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x47, 0x52, 0x50, 0x43, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb8, 0x03, 0x0a, 0x15,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x90, 0x01, 0x0a, 0x19, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x37, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x84, 0x01, 0x0a, 0x15, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x33, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x84, 0x01, 0x0a, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x47, 0x52,
	0x50, 0x43, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x33, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x47, 0x52, 0x50, 0x43,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x47, 0x52, 0x50, 0x43, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xe0, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68, 0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x53, 0x58, 0xaa, 0x02, 0x15, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69,
	0x73, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x17, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_gatewayapis_server_v1_register_proto_rawDescOnce sync.Once
	file_gatewayapis_server_v1_register_proto_rawDescData = file_gatewayapis_server_v1_register_proto_rawDesc
)

func file_gatewayapis_server_v1_register_proto_rawDescGZIP() []byte {
	file_gatewayapis_server_v1_register_proto_rawDescOnce.Do(func() {
		file_gatewayapis_server_v1_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_gatewayapis_server_v1_register_proto_rawDescData)
	})
	return file_gatewayapis_server_v1_register_proto_rawDescData
}

var file_gatewayapis_server_v1_register_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gatewayapis_server_v1_register_proto_goTypes = []interface{}{
	(*PostRegistServiceInstanceRequest)(nil),  // 0: gatewayapis.server.v1.PostRegistServiceInstanceRequest
	(*PostRegistServiceInstanceResponse)(nil), // 1: gatewayapis.server.v1.PostRegistServiceInstanceResponse
	(*PostRegistHTTPRoutersRequest)(nil),      // 2: gatewayapis.server.v1.PostRegistHTTPRoutersRequest
	(*PostRegistHTTPRoutersResponse)(nil),     // 3: gatewayapis.server.v1.PostRegistHTTPRoutersResponse
	(*PostRegistGRPCMethodsRequest)(nil),      // 4: gatewayapis.server.v1.PostRegistGRPCMethodsRequest
	(*PostRegistGRPCMethodsResponse)(nil),     // 5: gatewayapis.server.v1.PostRegistGRPCMethodsResponse
	(*ServiceInstance)(nil),                   // 6: gatewayapis.server.v1.ServiceInstance
	(*HTTPRouter)(nil),                        // 7: gatewayapis.server.v1.HTTPRouter
	(*GRPCMethod)(nil),                        // 8: gatewayapis.server.v1.GRPCMethod
}
var file_gatewayapis_server_v1_register_proto_depIdxs = []int32{
	6, // 0: gatewayapis.server.v1.PostRegistServiceInstanceRequest.instance:type_name -> gatewayapis.server.v1.ServiceInstance
	7, // 1: gatewayapis.server.v1.PostRegistHTTPRoutersRequest.routers:type_name -> gatewayapis.server.v1.HTTPRouter
	8, // 2: gatewayapis.server.v1.PostRegistGRPCMethodsRequest.methods:type_name -> gatewayapis.server.v1.GRPCMethod
	0, // 3: gatewayapis.server.v1.ServerRegisterService.PostRegistServiceInstance:input_type -> gatewayapis.server.v1.PostRegistServiceInstanceRequest
	2, // 4: gatewayapis.server.v1.ServerRegisterService.PostRegistHTTPRouters:input_type -> gatewayapis.server.v1.PostRegistHTTPRoutersRequest
	4, // 5: gatewayapis.server.v1.ServerRegisterService.PostRegistGRPCMethods:input_type -> gatewayapis.server.v1.PostRegistGRPCMethodsRequest
	1, // 6: gatewayapis.server.v1.ServerRegisterService.PostRegistServiceInstance:output_type -> gatewayapis.server.v1.PostRegistServiceInstanceResponse
	3, // 7: gatewayapis.server.v1.ServerRegisterService.PostRegistHTTPRouters:output_type -> gatewayapis.server.v1.PostRegistHTTPRoutersResponse
	5, // 8: gatewayapis.server.v1.ServerRegisterService.PostRegistGRPCMethods:output_type -> gatewayapis.server.v1.PostRegistGRPCMethodsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gatewayapis_server_v1_register_proto_init() }
func file_gatewayapis_server_v1_register_proto_init() {
	if File_gatewayapis_server_v1_register_proto != nil {
		return
	}
	file_gatewayapis_server_v1_defined_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gatewayapis_server_v1_register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRegistServiceInstanceRequest); i {
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
		file_gatewayapis_server_v1_register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRegistServiceInstanceResponse); i {
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
		file_gatewayapis_server_v1_register_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRegistHTTPRoutersRequest); i {
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
		file_gatewayapis_server_v1_register_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRegistHTTPRoutersResponse); i {
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
		file_gatewayapis_server_v1_register_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRegistGRPCMethodsRequest); i {
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
		file_gatewayapis_server_v1_register_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRegistGRPCMethodsResponse); i {
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
			RawDescriptor: file_gatewayapis_server_v1_register_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gatewayapis_server_v1_register_proto_goTypes,
		DependencyIndexes: file_gatewayapis_server_v1_register_proto_depIdxs,
		MessageInfos:      file_gatewayapis_server_v1_register_proto_msgTypes,
	}.Build()
	File_gatewayapis_server_v1_register_proto = out.File
	file_gatewayapis_server_v1_register_proto_rawDesc = nil
	file_gatewayapis_server_v1_register_proto_goTypes = nil
	file_gatewayapis_server_v1_register_proto_depIdxs = nil
}
