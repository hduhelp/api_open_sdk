// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: consultingapis/status.proto

package consultingapis

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

// 请求状态 错误定义
type Status int32

const (
	Status_UNSPECIFIED Status = 0
	// 成功
	Status_OK Status = 20000
	// 通用参数错误
	Status_INVALID_ARGUMENT Status = 40000
	// 人员类型参数缺失
	Status_STAFF_TYPE_NOT_FOUND Status = 40001
	// 不支持的人员类型
	Status_UNSUPPORTED_STAFF_TYPE Status = 40003
	// 人员工号缺失
	Status_STAFF_ID_CANNOT_BE_EMPTY Status = 40004
	// 通用无法找到
	Status_NOT_FOUND Status = 40400
	// 未找到人员
	Status_STAFF_NOT_FOUND Status = 40401
	// 服务错误
	Status_SERVICE_ERROR Status = 50000
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:     "UNSPECIFIED",
		20000: "OK",
		40000: "INVALID_ARGUMENT",
		40001: "STAFF_TYPE_NOT_FOUND",
		40003: "UNSUPPORTED_STAFF_TYPE",
		40004: "STAFF_ID_CANNOT_BE_EMPTY",
		40400: "NOT_FOUND",
		40401: "STAFF_NOT_FOUND",
		50000: "SERVICE_ERROR",
	}
	Status_value = map[string]int32{
		"UNSPECIFIED":              0,
		"OK":                       20000,
		"INVALID_ARGUMENT":         40000,
		"STAFF_TYPE_NOT_FOUND":     40001,
		"UNSUPPORTED_STAFF_TYPE":   40003,
		"STAFF_ID_CANNOT_BE_EMPTY": 40004,
		"NOT_FOUND":                40400,
		"STAFF_NOT_FOUND":          40401,
		"SERVICE_ERROR":            50000,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_consultingapis_status_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_consultingapis_status_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_consultingapis_status_proto_rawDescGZIP(), []int{0}
}

var File_consultingapis_status_proto protoreflect.FileDescriptor

var file_consultingapis_status_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x73, 0x2a, 0xd2, 0x01,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x02, 0x4f, 0x4b, 0x10,
	0xa0, 0x9c, 0x01, 0x12, 0x16, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41,
	0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0xc0, 0xb8, 0x02, 0x12, 0x1a, 0x0a, 0x14, 0x53,
	0x54, 0x41, 0x46, 0x46, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0xc1, 0xb8, 0x02, 0x12, 0x1c, 0x0a, 0x16, 0x55, 0x4e, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x46, 0x46, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0xc3, 0xb8, 0x02, 0x12, 0x1e, 0x0a, 0x18, 0x53, 0x54, 0x41, 0x46, 0x46, 0x5f, 0x49,
	0x44, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x45, 0x4d, 0x50, 0x54,
	0x59, 0x10, 0xc4, 0xb8, 0x02, 0x12, 0x0f, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0xd0, 0xbb, 0x02, 0x12, 0x15, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x46, 0x46, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xd1, 0xbb, 0x02, 0x12, 0x13, 0x0a,
	0x0d, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xd0,
	0x86, 0x03, 0x42, 0xa9, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x73, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68, 0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69,
	0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x73, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02,
	0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x73, 0xca,
	0x02, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x73,
	0xe2, 0x02, 0x1a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consultingapis_status_proto_rawDescOnce sync.Once
	file_consultingapis_status_proto_rawDescData = file_consultingapis_status_proto_rawDesc
)

func file_consultingapis_status_proto_rawDescGZIP() []byte {
	file_consultingapis_status_proto_rawDescOnce.Do(func() {
		file_consultingapis_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_consultingapis_status_proto_rawDescData)
	})
	return file_consultingapis_status_proto_rawDescData
}

var file_consultingapis_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_consultingapis_status_proto_goTypes = []any{
	(Status)(0), // 0: consultingapis.Status
}
var file_consultingapis_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_consultingapis_status_proto_init() }
func file_consultingapis_status_proto_init() {
	if File_consultingapis_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_consultingapis_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_consultingapis_status_proto_goTypes,
		DependencyIndexes: file_consultingapis_status_proto_depIdxs,
		EnumInfos:         file_consultingapis_status_proto_enumTypes,
	}.Build()
	File_consultingapis_status_proto = out.File
	file_consultingapis_status_proto_rawDesc = nil
	file_consultingapis_status_proto_goTypes = nil
	file_consultingapis_status_proto_depIdxs = nil
}
