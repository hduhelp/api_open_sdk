// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: campusapis/schoolTime/defined.proto

package schoolTime

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

//节次信息
type Section struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//第几节
	Section int32 `protobuf:"varint,1,opt,name=Section,proto3" json:"Section,omitempty"`
	//本节开始时间
	StartTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	//本节结束时间
	EndTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *Section) Reset() {
	*x = Section{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_schoolTime_defined_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Section) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Section) ProtoMessage() {}

func (x *Section) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schoolTime_defined_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Section.ProtoReflect.Descriptor instead.
func (*Section) Descriptor() ([]byte, []int) {
	return file_campusapis_schoolTime_defined_proto_rawDescGZIP(), []int{0}
}

func (x *Section) GetSection() int32 {
	if x != nil {
		return x.Section
	}
	return 0
}

func (x *Section) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Section) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

//学年
type SchoolYear struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 学年开始年，如：2021-2022 学年为 2021
	Year int32 `protobuf:"varint,1,opt,name=Year,proto3" json:"Year,omitempty"`
}

func (x *SchoolYear) Reset() {
	*x = SchoolYear{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_schoolTime_defined_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchoolYear) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchoolYear) ProtoMessage() {}

func (x *SchoolYear) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schoolTime_defined_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchoolYear.ProtoReflect.Descriptor instead.
func (*SchoolYear) Descriptor() ([]byte, []int) {
	return file_campusapis_schoolTime_defined_proto_rawDescGZIP(), []int{1}
}

func (x *SchoolYear) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

//学期
type Semester struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (x *Semester) Reset() {
	*x = Semester{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_schoolTime_defined_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Semester) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Semester) ProtoMessage() {}

func (x *Semester) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schoolTime_defined_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Semester.ProtoReflect.Descriptor instead.
func (*Semester) Descriptor() ([]byte, []int) {
	return file_campusapis_schoolTime_defined_proto_rawDescGZIP(), []int{2}
}

func (x *Semester) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_campusapis_schoolTime_defined_proto protoreflect.FileDescriptor

var file_campusapis_schoolTime_defined_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01,
	0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x59, 0x65, 0x61, 0x72, 0x22, 0x1c, 0x0a, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x4e, 0x75, 0x6d, 0x42, 0xd5, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x0c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64,
	0x75, 0x68, 0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73,
	0x64, 0x6b, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02,
	0x15, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0xca, 0x02, 0x15, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61,
	0x70, 0x69, 0x73, 0x5c, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0xe2, 0x02,
	0x21, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x53, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x16, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x3a,
	0x3a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_campusapis_schoolTime_defined_proto_rawDescOnce sync.Once
	file_campusapis_schoolTime_defined_proto_rawDescData = file_campusapis_schoolTime_defined_proto_rawDesc
)

func file_campusapis_schoolTime_defined_proto_rawDescGZIP() []byte {
	file_campusapis_schoolTime_defined_proto_rawDescOnce.Do(func() {
		file_campusapis_schoolTime_defined_proto_rawDescData = protoimpl.X.CompressGZIP(file_campusapis_schoolTime_defined_proto_rawDescData)
	})
	return file_campusapis_schoolTime_defined_proto_rawDescData
}

var file_campusapis_schoolTime_defined_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_campusapis_schoolTime_defined_proto_goTypes = []interface{}{
	(*Section)(nil),               // 0: campusapis.schoolTime.Section
	(*SchoolYear)(nil),            // 1: campusapis.schoolTime.SchoolYear
	(*Semester)(nil),              // 2: campusapis.schoolTime.Semester
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_campusapis_schoolTime_defined_proto_depIdxs = []int32{
	3, // 0: campusapis.schoolTime.Section.StartTime:type_name -> google.protobuf.Timestamp
	3, // 1: campusapis.schoolTime.Section.EndTime:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_campusapis_schoolTime_defined_proto_init() }
func file_campusapis_schoolTime_defined_proto_init() {
	if File_campusapis_schoolTime_defined_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_campusapis_schoolTime_defined_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Section); i {
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
		file_campusapis_schoolTime_defined_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchoolYear); i {
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
		file_campusapis_schoolTime_defined_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Semester); i {
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
			RawDescriptor: file_campusapis_schoolTime_defined_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_campusapis_schoolTime_defined_proto_goTypes,
		DependencyIndexes: file_campusapis_schoolTime_defined_proto_depIdxs,
		MessageInfos:      file_campusapis_schoolTime_defined_proto_msgTypes,
	}.Build()
	File_campusapis_schoolTime_defined_proto = out.File
	file_campusapis_schoolTime_defined_proto_rawDesc = nil
	file_campusapis_schoolTime_defined_proto_goTypes = nil
	file_campusapis_schoolTime_defined_proto_depIdxs = nil
}
