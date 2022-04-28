// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
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

//门禁服务名称，以品牌或者第三方服务商名称划分，用于出问题时摔锅
type Service int32

const (
	//未知设备
	Service_UNKNOWN Service = 0
	//海康闸机
	Service_HIKVISION Service = 1
	//宇视闸机
	Service_UNIVIEW Service = 2
	//生活区后勤管的闸机，即宿舍闸机
	Service_BACK_OFFICE Service = 3
)

// Enum value maps for Service.
var (
	Service_name = map[int32]string{
		0: "UNKNOWN",
		1: "HIKVISION",
		2: "UNIVIEW",
		3: "BACK_OFFICE",
	}
	Service_value = map[string]int32{
		"UNKNOWN":     0,
		"HIKVISION":   1,
		"UNIVIEW":     2,
		"BACK_OFFICE": 3,
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

//以门禁限制活动/保护的区域为主体的进出方向
type Direction int32

const (
	//进
	Direction_IN Direction = 0
	//出
	Direction_OUT Direction = 1
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "IN",
		1: "OUT",
	}
	Direction_value = map[string]int32{
		"IN":  0,
		"OUT": 1,
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

//门禁类型
type GateType int32

const (
	//学校大门门禁，包括生活区、教学区、研究生公寓门禁
	GateType_CAMPUS GateType = 0
	//宿舍门禁
	GateType_DORMITORY GateType = 1
	//图书馆门禁
	GateType_LIBRARY GateType = 2
)

// Enum value maps for GateType.
var (
	GateType_name = map[int32]string{
		0: "CAMPUS",
		1: "DORMITORY",
		2: "LIBRARY",
	}
	GateType_value = map[string]int32{
		"CAMPUS":    0,
		"DORMITORY": 1,
		"LIBRARY":   2,
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

//通过方式
type AccessMethod int32

const (
	//校园卡
	AccessMethod_SCHOOL_CARD AccessMethod = 0
	//人脸识别
	AccessMethod_FACE AccessMethod = 1
)

// Enum value maps for AccessMethod.
var (
	AccessMethod_name = map[int32]string{
		0: "SCHOOL_CARD",
		1: "FACE",
	}
	AccessMethod_value = map[string]int32{
		"SCHOOL_CARD": 0,
		"FACE":        1,
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
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2a, 0x43, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x48, 0x49, 0x4b, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x49, 0x56, 0x49, 0x45, 0x57, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x41,
	0x43, 0x4b, 0x5f, 0x4f, 0x46, 0x46, 0x49, 0x43, 0x45, 0x10, 0x03, 0x2a, 0x1c, 0x0a, 0x09, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x4e, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x4f, 0x55, 0x54, 0x10, 0x01, 0x2a, 0x32, 0x0a, 0x08, 0x47, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x41, 0x4d, 0x50, 0x55, 0x53, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x4f, 0x52, 0x4d, 0x49, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49, 0x42, 0x52, 0x41, 0x52, 0x59, 0x10, 0x02, 0x2a, 0x29, 0x0a,
	0x0c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0f, 0x0a,
	0x0b, 0x53, 0x43, 0x48, 0x4f, 0x4f, 0x4c, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x46, 0x41, 0x43, 0x45, 0x10, 0x01, 0x42, 0xcb, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x0c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x64, 0x75, 0x68, 0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e,
	0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x61, 0x74, 0x65, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x47, 0x58, 0xaa, 0x02, 0x12, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x43, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1e, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x47, 0x61, 0x74, 0x65,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x14, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x47, 0x61,
	0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_campusapis_gate_v1_defined_proto_goTypes = []interface{}{
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
