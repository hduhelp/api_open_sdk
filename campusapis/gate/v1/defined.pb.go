// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: campusapis/gate/v1/defined.proto

package gatev1

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

// 门禁服务名称，以品牌或者第三方服务商名称划分，用于出问题时摔锅
type Service int32

const (
	// 未知设备
	Service_SERVICE_UNSPECIFIED Service = 0
	// 海康闸机
	Service_SERVICE_HIKVISION Service = 1
	// 宇视闸机
	Service_SERVICE_UNIVIEW Service = 2
	// 生活区后勤管的闸机，即宿舍闸机
	Service_SERVICE_BACK_OFFICE Service = 3
)

// Enum value maps for Service.
var (
	Service_name = map[int32]string{
		0: "SERVICE_UNSPECIFIED",
		1: "SERVICE_HIKVISION",
		2: "SERVICE_UNIVIEW",
		3: "SERVICE_BACK_OFFICE",
	}
	Service_value = map[string]int32{
		"SERVICE_UNSPECIFIED": 0,
		"SERVICE_HIKVISION":   1,
		"SERVICE_UNIVIEW":     2,
		"SERVICE_BACK_OFFICE": 3,
	}
)

func (x Service) Enum() *Service {
	p := new(Service)
	*p = x
	return p
}

func (x Service) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Service) Descriptor() protoreflect.EnumDescriptor {
	return file_campusapis_gate_v1_defined_proto_enumTypes[0].Descriptor()
}

func (Service) Type() protoreflect.EnumType {
	return &file_campusapis_gate_v1_defined_proto_enumTypes[0]
}

func (x Service) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Service.Descriptor instead.
func (Service) EnumDescriptor() ([]byte, []int) {
	return file_campusapis_gate_v1_defined_proto_rawDescGZIP(), []int{0}
}

// 以门禁限制活动/保护的区域为主体的进出方向
type Direction int32

const (
	Direction_DIRECTION_UNSPECIFIED Direction = 0
	// 进
	Direction_DIRECTION_IN Direction = 1
	// 出
	Direction_DIRECTION_OUT Direction = 2
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "DIRECTION_UNSPECIFIED",
		1: "DIRECTION_IN",
		2: "DIRECTION_OUT",
	}
	Direction_value = map[string]int32{
		"DIRECTION_UNSPECIFIED": 0,
		"DIRECTION_IN":          1,
		"DIRECTION_OUT":         2,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_campusapis_gate_v1_defined_proto_enumTypes[1].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_campusapis_gate_v1_defined_proto_enumTypes[1]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_campusapis_gate_v1_defined_proto_rawDescGZIP(), []int{1}
}

// 门禁类型
type GateType int32

const (
	GateType_GATE_TYPE_UNSPECIFIED GateType = 0
	// 学校大门门禁，包括生活区、教学区、研究生公寓门禁
	GateType_GATE_TYPE_CAMPUS GateType = 1
	// 宿舍门禁
	GateType_GATE_TYPE_DORMITORY GateType = 2
	// 图书馆门禁
	GateType_GATE_TYPE_LIBRARY GateType = 3
)

// Enum value maps for GateType.
var (
	GateType_name = map[int32]string{
		0: "GATE_TYPE_UNSPECIFIED",
		1: "GATE_TYPE_CAMPUS",
		2: "GATE_TYPE_DORMITORY",
		3: "GATE_TYPE_LIBRARY",
	}
	GateType_value = map[string]int32{
		"GATE_TYPE_UNSPECIFIED": 0,
		"GATE_TYPE_CAMPUS":      1,
		"GATE_TYPE_DORMITORY":   2,
		"GATE_TYPE_LIBRARY":     3,
	}
)

func (x GateType) Enum() *GateType {
	p := new(GateType)
	*p = x
	return p
}

func (x GateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GateType) Descriptor() protoreflect.EnumDescriptor {
	return file_campusapis_gate_v1_defined_proto_enumTypes[2].Descriptor()
}

func (GateType) Type() protoreflect.EnumType {
	return &file_campusapis_gate_v1_defined_proto_enumTypes[2]
}

func (x GateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GateType.Descriptor instead.
func (GateType) EnumDescriptor() ([]byte, []int) {
	return file_campusapis_gate_v1_defined_proto_rawDescGZIP(), []int{2}
}

// 通过方式
type AccessMethod int32

const (
	AccessMethod_ACCESS_METHOD_UNSPECIFIED AccessMethod = 0
	// 校园卡
	AccessMethod_ACCESS_METHOD_SCHOOL_CARD AccessMethod = 1
	// 人脸识别
	AccessMethod_ACCESS_METHOD_FACE AccessMethod = 2
)

// Enum value maps for AccessMethod.
var (
	AccessMethod_name = map[int32]string{
		0: "ACCESS_METHOD_UNSPECIFIED",
		1: "ACCESS_METHOD_SCHOOL_CARD",
		2: "ACCESS_METHOD_FACE",
	}
	AccessMethod_value = map[string]int32{
		"ACCESS_METHOD_UNSPECIFIED": 0,
		"ACCESS_METHOD_SCHOOL_CARD": 1,
		"ACCESS_METHOD_FACE":        2,
	}
)

func (x AccessMethod) Enum() *AccessMethod {
	p := new(AccessMethod)
	*p = x
	return p
}

func (x AccessMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccessMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_campusapis_gate_v1_defined_proto_enumTypes[3].Descriptor()
}

func (AccessMethod) Type() protoreflect.EnumType {
	return &file_campusapis_gate_v1_defined_proto_enumTypes[3]
}

func (x AccessMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccessMethod.Descriptor instead.
func (AccessMethod) EnumDescriptor() ([]byte, []int) {
	return file_campusapis_gate_v1_defined_proto_rawDescGZIP(), []int{3}
}

var File_campusapis_gate_v1_defined_proto protoreflect.FileDescriptor

var file_campusapis_gate_v1_defined_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2a, 0x67, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x48, 0x49, 0x4b, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x49,
	0x56, 0x49, 0x45, 0x57, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x4f, 0x46, 0x46, 0x49, 0x43, 0x45, 0x10, 0x03, 0x2a,
	0x4b, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15,
	0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x52, 0x45, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x49, 0x52,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x02, 0x2a, 0x6b, 0x0a, 0x08,
	0x47, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x41, 0x54, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x41, 0x4d, 0x50, 0x55, 0x53, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x41, 0x54,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x52, 0x4d, 0x49, 0x54, 0x4f, 0x52, 0x59,
	0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4c, 0x49, 0x42, 0x52, 0x41, 0x52, 0x59, 0x10, 0x03, 0x2a, 0x64, 0x0a, 0x0c, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x53, 0x43, 0x48, 0x4f, 0x4f, 0x4c,
	0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x46, 0x41, 0x43, 0x45, 0x10, 0x02, 0x42,
	0xcb, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68, 0x65, 0x6c, 0x70, 0x2f, 0x61,
	0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x67,
	0x61, 0x74, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x47, 0x58, 0xaa, 0x02, 0x12, 0x43, 0x61,
	0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x12, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x47, 0x61,
	0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70,
	0x69, 0x73, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61,
	0x70, 0x69, 0x73, 0x3a, 0x3a, 0x47, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_campusapis_gate_v1_defined_proto_rawDescOnce sync.Once
	file_campusapis_gate_v1_defined_proto_rawDescData = file_campusapis_gate_v1_defined_proto_rawDesc
)

func file_campusapis_gate_v1_defined_proto_rawDescGZIP() []byte {
	file_campusapis_gate_v1_defined_proto_rawDescOnce.Do(func() {
		file_campusapis_gate_v1_defined_proto_rawDescData = protoimpl.X.CompressGZIP(file_campusapis_gate_v1_defined_proto_rawDescData)
	})
	return file_campusapis_gate_v1_defined_proto_rawDescData
}

var file_campusapis_gate_v1_defined_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_campusapis_gate_v1_defined_proto_goTypes = []any{
	(Service)(0),      // 0: campusapis.gate.v1.Service
	(Direction)(0),    // 1: campusapis.gate.v1.Direction
	(GateType)(0),     // 2: campusapis.gate.v1.GateType
	(AccessMethod)(0), // 3: campusapis.gate.v1.AccessMethod
}
var file_campusapis_gate_v1_defined_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_campusapis_gate_v1_defined_proto_init() }
func file_campusapis_gate_v1_defined_proto_init() {
	if File_campusapis_gate_v1_defined_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_campusapis_gate_v1_defined_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_campusapis_gate_v1_defined_proto_goTypes,
		DependencyIndexes: file_campusapis_gate_v1_defined_proto_depIdxs,
		EnumInfos:         file_campusapis_gate_v1_defined_proto_enumTypes,
	}.Build()
	File_campusapis_gate_v1_defined_proto = out.File
	file_campusapis_gate_v1_defined_proto_rawDesc = nil
	file_campusapis_gate_v1_defined_proto_goTypes = nil
	file_campusapis_gate_v1_defined_proto_depIdxs = nil
}
