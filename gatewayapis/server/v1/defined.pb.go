// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: gatewayapis/server/v1/defined.proto

package serverv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//服务
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//服务名称 同时作为ID使用
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//服务描述
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	//服务路由前缀
	RoutePrefix string `protobuf:"bytes,3,opt,name=route_prefix,json=routePrefix,proto3" json:"route_prefix,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_defined_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_defined_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_defined_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Service) GetRoutePrefix() string {
	if x != nil {
		return x.RoutePrefix
	}
	return ""
}

//服务实例
type ServiceInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//实例ID
	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	//实例名称
	InstanceName string `protobuf:"bytes,2,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	//服务名称
	ServiceName string `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *ServiceInstance) Reset() {
	*x = ServiceInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_defined_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInstance) ProtoMessage() {}

func (x *ServiceInstance) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_defined_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInstance.ProtoReflect.Descriptor instead.
func (*ServiceInstance) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_defined_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceInstance) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *ServiceInstance) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *ServiceInstance) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

//HTTP路由注册
type HTTPRouter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//服务名称
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	//方法
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	//远端path，即服务内地址
	RemotePath string `protobuf:"bytes,3,opt,name=remotePath,proto3" json:"remotePath,omitempty"`
	//挂载地址
	ServePath string `protobuf:"bytes,4,opt,name=servePath,proto3" json:"servePath,omitempty"`
	//是否需要鉴权后才能访问
	NeedAuthorized bool `protobuf:"varint,5,opt,name=needAuthorized,proto3" json:"needAuthorized,omitempty"`
	//是否开启只读缓存
	ReadOnlyCache bool `protobuf:"varint,6,opt,name=readOnlyCache,proto3" json:"readOnlyCache,omitempty"`
	//公共缓存
	PublicCache bool `protobuf:"varint,7,opt,name=publicCache,proto3" json:"publicCache,omitempty"`
	//缓存过期时间
	CacheExpire uint32 `protobuf:"varint,8,opt,name=cacheExpire,proto3" json:"cacheExpire,omitempty"`
	//客户端到服务请求可传递的请求头
	DirectThroughRequestHeaders string `protobuf:"bytes,9,opt,name=directThroughRequestHeaders,proto3" json:"directThroughRequestHeaders,omitempty"`
	//服务到客户端返回可传递的响应头
	DirectThroughResponseHeaders string `protobuf:"bytes,10,opt,name=directThroughResponseHeaders,proto3" json:"directThroughResponseHeaders,omitempty"`
	//是否启用
	Enable bool `protobuf:"varint,11,opt,name=enable,proto3" json:"enable,omitempty"`
}

func (x *HTTPRouter) Reset() {
	*x = HTTPRouter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_defined_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPRouter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPRouter) ProtoMessage() {}

func (x *HTTPRouter) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_defined_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPRouter.ProtoReflect.Descriptor instead.
func (*HTTPRouter) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_defined_proto_rawDescGZIP(), []int{2}
}

func (x *HTTPRouter) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *HTTPRouter) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HTTPRouter) GetRemotePath() string {
	if x != nil {
		return x.RemotePath
	}
	return ""
}

func (x *HTTPRouter) GetServePath() string {
	if x != nil {
		return x.ServePath
	}
	return ""
}

func (x *HTTPRouter) GetNeedAuthorized() bool {
	if x != nil {
		return x.NeedAuthorized
	}
	return false
}

func (x *HTTPRouter) GetReadOnlyCache() bool {
	if x != nil {
		return x.ReadOnlyCache
	}
	return false
}

func (x *HTTPRouter) GetPublicCache() bool {
	if x != nil {
		return x.PublicCache
	}
	return false
}

func (x *HTTPRouter) GetCacheExpire() uint32 {
	if x != nil {
		return x.CacheExpire
	}
	return 0
}

func (x *HTTPRouter) GetDirectThroughRequestHeaders() string {
	if x != nil {
		return x.DirectThroughRequestHeaders
	}
	return ""
}

func (x *HTTPRouter) GetDirectThroughResponseHeaders() string {
	if x != nil {
		return x.DirectThroughResponseHeaders
	}
	return ""
}

func (x *HTTPRouter) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

//GRPC路由 GRPC方法
type GRPCMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//服务名称
	//Service names include the package names, in the form of <package>.<service>.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	//方法名称
	MethodName     string `protobuf:"bytes,2,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	IsClientStream bool   `protobuf:"varint,3,opt,name=isClientStream,proto3" json:"isClientStream,omitempty"`
	IsServerStream bool   `protobuf:"varint,4,opt,name=isServerStream,proto3" json:"isServerStream,omitempty"`
}

func (x *GRPCMethod) Reset() {
	*x = GRPCMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_server_v1_defined_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCMethod) ProtoMessage() {}

func (x *GRPCMethod) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_server_v1_defined_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCMethod.ProtoReflect.Descriptor instead.
func (*GRPCMethod) Descriptor() ([]byte, []int) {
	return file_gatewayapis_server_v1_defined_proto_rawDescGZIP(), []int{3}
}

func (x *GRPCMethod) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *GRPCMethod) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *GRPCMethod) GetIsClientStream() bool {
	if x != nil {
		return x.IsClientStream
	}
	return false
}

func (x *GRPCMethod) GetIsServerStream() bool {
	if x != nil {
		return x.IsServerStream
	}
	return false
}

var File_gatewayapis_server_v1_defined_proto protoreflect.FileDescriptor

var file_gatewayapis_server_v1_defined_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x7a, 0x0a, 0x0f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xac, 0x03, 0x0a, 0x0a, 0x48, 0x54, 0x54, 0x50, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x65, 0x65, 0x64, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6e,
	0x65, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x40, 0x0a, 0x1b, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x54, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x54, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x42, 0x0a, 0x1c, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x54, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x1c, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x54, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x0a, 0x47, 0x52, 0x50, 0x43, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x73, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x69, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x26, 0x0a, 0x0e, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0xdf, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68, 0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f,
	0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x53, 0x58, 0xaa, 0x02, 0x15, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70,
	0x69, 0x73, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x17, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x3a, 0x3a,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_gatewayapis_server_v1_defined_proto_rawDescOnce sync.Once
	file_gatewayapis_server_v1_defined_proto_rawDescData = file_gatewayapis_server_v1_defined_proto_rawDesc
)

func file_gatewayapis_server_v1_defined_proto_rawDescGZIP() []byte {
	file_gatewayapis_server_v1_defined_proto_rawDescOnce.Do(func() {
		file_gatewayapis_server_v1_defined_proto_rawDescData = protoimpl.X.CompressGZIP(file_gatewayapis_server_v1_defined_proto_rawDescData)
	})
	return file_gatewayapis_server_v1_defined_proto_rawDescData
}

var file_gatewayapis_server_v1_defined_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gatewayapis_server_v1_defined_proto_goTypes = []interface{}{
	(*Service)(nil),         // 0: gatewayapis.server.v1.Service
	(*ServiceInstance)(nil), // 1: gatewayapis.server.v1.ServiceInstance
	(*HTTPRouter)(nil),      // 2: gatewayapis.server.v1.HTTPRouter
	(*GRPCMethod)(nil),      // 3: gatewayapis.server.v1.GRPCMethod
}
var file_gatewayapis_server_v1_defined_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gatewayapis_server_v1_defined_proto_init() }
func file_gatewayapis_server_v1_defined_proto_init() {
	if File_gatewayapis_server_v1_defined_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gatewayapis_server_v1_defined_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_gatewayapis_server_v1_defined_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInstance); i {
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
		file_gatewayapis_server_v1_defined_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPRouter); i {
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
		file_gatewayapis_server_v1_defined_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCMethod); i {
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
			RawDescriptor: file_gatewayapis_server_v1_defined_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gatewayapis_server_v1_defined_proto_goTypes,
		DependencyIndexes: file_gatewayapis_server_v1_defined_proto_depIdxs,
		MessageInfos:      file_gatewayapis_server_v1_defined_proto_msgTypes,
	}.Build()
	File_gatewayapis_server_v1_defined_proto = out.File
	file_gatewayapis_server_v1_defined_proto_rawDesc = nil
	file_gatewayapis_server_v1_defined_proto_goTypes = nil
	file_gatewayapis_server_v1_defined_proto_depIdxs = nil
}
