// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: campusapis/gate/v1/callback.proto

package gatev1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 推送门禁事件回调请求体
type PostGateEventCallbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffId   string `protobuf:"bytes,1,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
	StaffName string `protobuf:"bytes,2,opt,name=staff_name,json=staffName,proto3" json:"staff_name,omitempty"`
	//事件时间
	Time *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	//门禁名称
	GateName  string       `protobuf:"bytes,4,opt,name=gate_name,json=gateName,proto3" json:"gate_name,omitempty"`
	GateType  GateType     `protobuf:"varint,5,opt,name=gate_type,json=gateType,proto3,enum=campusapis.gate.v1.GateType" json:"gate_type,omitempty"`
	Direction Direction    `protobuf:"varint,6,opt,name=direction,proto3,enum=campusapis.gate.v1.Direction" json:"direction,omitempty"`
	Service   Service      `protobuf:"varint,7,opt,name=service,proto3,enum=campusapis.gate.v1.Service" json:"service,omitempty"`
	Method    AccessMethod `protobuf:"varint,8,opt,name=method,proto3,enum=campusapis.gate.v1.AccessMethod" json:"method,omitempty"`
}

func (x *PostGateEventCallbackRequest) Reset() {
	*x = PostGateEventCallbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_gate_v1_callback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostGateEventCallbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostGateEventCallbackRequest) ProtoMessage() {}

func (x *PostGateEventCallbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_gate_v1_callback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostGateEventCallbackRequest.ProtoReflect.Descriptor instead.
func (*PostGateEventCallbackRequest) Descriptor() ([]byte, []int) {
	return file_campusapis_gate_v1_callback_proto_rawDescGZIP(), []int{0}
}

func (x *PostGateEventCallbackRequest) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

func (x *PostGateEventCallbackRequest) GetStaffName() string {
	if x != nil {
		return x.StaffName
	}
	return ""
}

func (x *PostGateEventCallbackRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *PostGateEventCallbackRequest) GetGateName() string {
	if x != nil {
		return x.GateName
	}
	return ""
}

func (x *PostGateEventCallbackRequest) GetGateType() GateType {
	if x != nil {
		return x.GateType
	}
	return GateType_CAMPUS
}

func (x *PostGateEventCallbackRequest) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_IN
}

func (x *PostGateEventCallbackRequest) GetService() Service {
	if x != nil {
		return x.Service
	}
	return Service_UNKNOWN
}

func (x *PostGateEventCallbackRequest) GetMethod() AccessMethod {
	if x != nil {
		return x.Method
	}
	return AccessMethod_SCHOOL_CARD
}

// 推送门禁事件回调返回体
type PostGateEventCallbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostGateEventCallbackResponse) Reset() {
	*x = PostGateEventCallbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_gate_v1_callback_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostGateEventCallbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostGateEventCallbackResponse) ProtoMessage() {}

func (x *PostGateEventCallbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_gate_v1_callback_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostGateEventCallbackResponse.ProtoReflect.Descriptor instead.
func (*PostGateEventCallbackResponse) Descriptor() ([]byte, []int) {
	return file_campusapis_gate_v1_callback_proto_rawDescGZIP(), []int{1}
}

var File_campusapis_gate_v1_callback_proto protoreflect.FileDescriptor

var file_campusapis_gate_v1_callback_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x03, 0x0a, 0x1c, 0x50, 0x6f, 0x73, 0x74, 0x47, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x09, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x67, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x63, 0x61,
	0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x50, 0x6f, 0x73, 0x74, 0x47, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x95, 0x01, 0x0a, 0x13, 0x47, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x7e, 0x0a, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x30, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75,
	0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x61, 0x6d,
	0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0xcc, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68, 0x65, 0x6c, 0x70, 0x2f,
	0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x61, 0x6d,
	0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x67, 0x61, 0x74, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x47, 0x58, 0xaa, 0x02, 0x12, 0x43,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x12, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x47,
	0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61,
	0x70, 0x69, 0x73, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x61, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x47, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_campusapis_gate_v1_callback_proto_rawDescOnce sync.Once
	file_campusapis_gate_v1_callback_proto_rawDescData = file_campusapis_gate_v1_callback_proto_rawDesc
)

func file_campusapis_gate_v1_callback_proto_rawDescGZIP() []byte {
	file_campusapis_gate_v1_callback_proto_rawDescOnce.Do(func() {
		file_campusapis_gate_v1_callback_proto_rawDescData = protoimpl.X.CompressGZIP(file_campusapis_gate_v1_callback_proto_rawDescData)
	})
	return file_campusapis_gate_v1_callback_proto_rawDescData
}

var file_campusapis_gate_v1_callback_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_campusapis_gate_v1_callback_proto_goTypes = []interface{}{
	(*PostGateEventCallbackRequest)(nil),  // 0: campusapis.gate.v1.PostGateEventCallbackRequest
	(*PostGateEventCallbackResponse)(nil), // 1: campusapis.gate.v1.PostGateEventCallbackResponse
	(*timestamppb.Timestamp)(nil),         // 2: google.protobuf.Timestamp
	(GateType)(0),                         // 3: campusapis.gate.v1.GateType
	(Direction)(0),                        // 4: campusapis.gate.v1.Direction
	(Service)(0),                          // 5: campusapis.gate.v1.Service
	(AccessMethod)(0),                     // 6: campusapis.gate.v1.AccessMethod
}
var file_campusapis_gate_v1_callback_proto_depIdxs = []int32{
	2, // 0: campusapis.gate.v1.PostGateEventCallbackRequest.time:type_name -> google.protobuf.Timestamp
	3, // 1: campusapis.gate.v1.PostGateEventCallbackRequest.gate_type:type_name -> campusapis.gate.v1.GateType
	4, // 2: campusapis.gate.v1.PostGateEventCallbackRequest.direction:type_name -> campusapis.gate.v1.Direction
	5, // 3: campusapis.gate.v1.PostGateEventCallbackRequest.service:type_name -> campusapis.gate.v1.Service
	6, // 4: campusapis.gate.v1.PostGateEventCallbackRequest.method:type_name -> campusapis.gate.v1.AccessMethod
	0, // 5: campusapis.gate.v1.GateCallbackService.PostGateEventCallback:input_type -> campusapis.gate.v1.PostGateEventCallbackRequest
	1, // 6: campusapis.gate.v1.GateCallbackService.PostGateEventCallback:output_type -> campusapis.gate.v1.PostGateEventCallbackResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_campusapis_gate_v1_callback_proto_init() }
func file_campusapis_gate_v1_callback_proto_init() {
	if File_campusapis_gate_v1_callback_proto != nil {
		return
	}
	file_campusapis_gate_v1_defined_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_campusapis_gate_v1_callback_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostGateEventCallbackRequest); i {
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
		file_campusapis_gate_v1_callback_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostGateEventCallbackResponse); i {
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
			RawDescriptor: file_campusapis_gate_v1_callback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_campusapis_gate_v1_callback_proto_goTypes,
		DependencyIndexes: file_campusapis_gate_v1_callback_proto_depIdxs,
		MessageInfos:      file_campusapis_gate_v1_callback_proto_msgTypes,
	}.Build()
	File_campusapis_gate_v1_callback_proto = out.File
	file_campusapis_gate_v1_callback_proto_rawDesc = nil
	file_campusapis_gate_v1_callback_proto_goTypes = nil
	file_campusapis_gate_v1_callback_proto_depIdxs = nil
}