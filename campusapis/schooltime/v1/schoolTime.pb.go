// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: campusapis/schooltime/v1/schoolTime.proto

package schooltimev1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type GetSchoolTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchoolYear string `protobuf:"bytes,1,opt,name=schoolYear,proto3" json:"schoolYear,omitempty"`
	Section    int32  `protobuf:"varint,2,opt,name=section,proto3" json:"section,omitempty"`
	Semester   string `protobuf:"bytes,3,opt,name=semester,proto3" json:"semester,omitempty"`
	TimeStamp  int64  `protobuf:"varint,4,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	WeekDayNow int32  `protobuf:"varint,5,opt,name=weekDayNow,proto3" json:"weekDayNow,omitempty"`
	WeekNow    int32  `protobuf:"varint,6,opt,name=weekNow,proto3" json:"weekNow,omitempty"`
}

func (x *GetSchoolTimeResponse) Reset() {
	*x = GetSchoolTimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSchoolTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchoolTimeResponse) ProtoMessage() {}

func (x *GetSchoolTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchoolTimeResponse.ProtoReflect.Descriptor instead.
func (*GetSchoolTimeResponse) Descriptor() ([]byte, []int) {
	return file_campusapis_schooltime_v1_schoolTime_proto_rawDescGZIP(), []int{0}
}

func (x *GetSchoolTimeResponse) GetSchoolYear() string {
	if x != nil {
		return x.SchoolYear
	}
	return ""
}

func (x *GetSchoolTimeResponse) GetSection() int32 {
	if x != nil {
		return x.Section
	}
	return 0
}

func (x *GetSchoolTimeResponse) GetSemester() string {
	if x != nil {
		return x.Semester
	}
	return ""
}

func (x *GetSchoolTimeResponse) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

func (x *GetSchoolTimeResponse) GetWeekDayNow() int32 {
	if x != nil {
		return x.WeekDayNow
	}
	return 0
}

func (x *GetSchoolTimeResponse) GetWeekNow() int32 {
	if x != nil {
		return x.WeekNow
	}
	return 0
}

type SchoolTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchoolYear *SchoolYear            `protobuf:"bytes,1,opt,name=SchoolYear,proto3" json:"SchoolYear,omitempty"`
	Semester   *Semester              `protobuf:"bytes,2,opt,name=Semester,proto3" json:"Semester,omitempty"`
	SchoolWeek int32                  `protobuf:"varint,3,opt,name=SchoolWeek,proto3" json:"SchoolWeek,omitempty"`
	DayOfWeek  int32                  `protobuf:"varint,4,opt,name=DayOfWeek,proto3" json:"DayOfWeek,omitempty"`
	Section    *Section               `protobuf:"bytes,5,opt,name=Section,proto3" json:"Section,omitempty"`
	Time       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (x *SchoolTime) Reset() {
	*x = SchoolTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchoolTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchoolTime) ProtoMessage() {}

func (x *SchoolTime) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchoolTime.ProtoReflect.Descriptor instead.
func (*SchoolTime) Descriptor() ([]byte, []int) {
	return file_campusapis_schooltime_v1_schoolTime_proto_rawDescGZIP(), []int{1}
}

func (x *SchoolTime) GetSchoolYear() *SchoolYear {
	if x != nil {
		return x.SchoolYear
	}
	return nil
}

func (x *SchoolTime) GetSemester() *Semester {
	if x != nil {
		return x.Semester
	}
	return nil
}

func (x *SchoolTime) GetSchoolWeek() int32 {
	if x != nil {
		return x.SchoolWeek
	}
	return 0
}

func (x *SchoolTime) GetDayOfWeek() int32 {
	if x != nil {
		return x.DayOfWeek
	}
	return 0
}

func (x *SchoolTime) GetSection() *Section {
	if x != nil {
		return x.Section
	}
	return nil
}

func (x *SchoolTime) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type Section struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section   int32                  `protobuf:"varint,1,opt,name=Section,proto3" json:"Section,omitempty"`
	StartTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *Section) Reset() {
	*x = Section{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Section) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Section) ProtoMessage() {}

func (x *Section) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[2]
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
	return file_campusapis_schooltime_v1_schoolTime_proto_rawDescGZIP(), []int{2}
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

type SchoolYear struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year int32 `protobuf:"varint,1,opt,name=Year,proto3" json:"Year,omitempty"`
}

func (x *SchoolYear) Reset() {
	*x = SchoolYear{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchoolYear) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchoolYear) ProtoMessage() {}

func (x *SchoolYear) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[3]
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
	return file_campusapis_schooltime_v1_schoolTime_proto_rawDescGZIP(), []int{3}
}

func (x *SchoolYear) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type Semester struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (x *Semester) Reset() {
	*x = Semester{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Semester) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Semester) ProtoMessage() {}

func (x *Semester) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[4]
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
	return file_campusapis_schooltime_v1_schoolTime_proto_rawDescGZIP(), []int{4}
}

func (x *Semester) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_campusapis_schooltime_v1_schoolTime_proto protoreflect.FileDescriptor

var file_campusapis_schooltime_v1_schoolTime_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x61, 0x6d,
	0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x1e, 0x0a, 0x0a, 0x77, 0x65, 0x65, 0x6b, 0x44, 0x61, 0x79, 0x4e, 0x6f, 0x77, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x65, 0x65, 0x6b, 0x44, 0x61, 0x79, 0x4e, 0x6f, 0x77, 0x12,
	0x18, 0x0a, 0x07, 0x77, 0x65, 0x65, 0x6b, 0x4e, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x77, 0x65, 0x65, 0x6b, 0x4e, 0x6f, 0x77, 0x22, 0xbd, 0x02, 0x0a, 0x0a, 0x53, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65,
	0x61, 0x72, 0x52, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x12, 0x3e,
	0x0a, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6d, 0x65,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x1c,
	0x0a, 0x09, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x3b, 0x0a, 0x07,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x38, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x20, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x59, 0x65, 0x61,
	0x72, 0x22, 0x1c, 0x0a, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x32,
	0x8e, 0x01, 0x0a, 0x11, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x79, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2f,
	0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x5a, 0x07, 0x12, 0x05, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x0e, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31,
	0x42, 0xf8, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x0f, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x64, 0x75, 0x68, 0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65,
	0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53,
	0x58, 0xaa, 0x02, 0x18, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x53,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x43,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x74, 0x69, 0x6d, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x61, 0x70, 0x69, 0x73, 0x5c, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1a, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x53, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_campusapis_schooltime_v1_schoolTime_proto_rawDescOnce sync.Once
	file_campusapis_schooltime_v1_schoolTime_proto_rawDescData = file_campusapis_schooltime_v1_schoolTime_proto_rawDesc
)

func file_campusapis_schooltime_v1_schoolTime_proto_rawDescGZIP() []byte {
	file_campusapis_schooltime_v1_schoolTime_proto_rawDescOnce.Do(func() {
		file_campusapis_schooltime_v1_schoolTime_proto_rawDescData = protoimpl.X.CompressGZIP(file_campusapis_schooltime_v1_schoolTime_proto_rawDescData)
	})
	return file_campusapis_schooltime_v1_schoolTime_proto_rawDescData
}

var file_campusapis_schooltime_v1_schoolTime_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_campusapis_schooltime_v1_schoolTime_proto_goTypes = []interface{}{
	(*GetSchoolTimeResponse)(nil), // 0: campusapis.schooltime.v1.GetSchoolTimeResponse
	(*SchoolTime)(nil),            // 1: campusapis.schooltime.v1.SchoolTime
	(*Section)(nil),               // 2: campusapis.schooltime.v1.Section
	(*SchoolYear)(nil),            // 3: campusapis.schooltime.v1.SchoolYear
	(*Semester)(nil),              // 4: campusapis.schooltime.v1.Semester
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_campusapis_schooltime_v1_schoolTime_proto_depIdxs = []int32{
	3, // 0: campusapis.schooltime.v1.SchoolTime.SchoolYear:type_name -> campusapis.schooltime.v1.SchoolYear
	4, // 1: campusapis.schooltime.v1.SchoolTime.Semester:type_name -> campusapis.schooltime.v1.Semester
	2, // 2: campusapis.schooltime.v1.SchoolTime.Section:type_name -> campusapis.schooltime.v1.Section
	5, // 3: campusapis.schooltime.v1.SchoolTime.Time:type_name -> google.protobuf.Timestamp
	5, // 4: campusapis.schooltime.v1.Section.StartTime:type_name -> google.protobuf.Timestamp
	5, // 5: campusapis.schooltime.v1.Section.EndTime:type_name -> google.protobuf.Timestamp
	6, // 6: campusapis.schooltime.v1.SchoolTimeService.GetSchoolTime:input_type -> google.protobuf.Empty
	0, // 7: campusapis.schooltime.v1.SchoolTimeService.GetSchoolTime:output_type -> campusapis.schooltime.v1.GetSchoolTimeResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_campusapis_schooltime_v1_schoolTime_proto_init() }
func file_campusapis_schooltime_v1_schoolTime_proto_init() {
	if File_campusapis_schooltime_v1_schoolTime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSchoolTimeResponse); i {
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
		file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchoolTime); i {
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
		file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_campusapis_schooltime_v1_schoolTime_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_campusapis_schooltime_v1_schoolTime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_campusapis_schooltime_v1_schoolTime_proto_goTypes,
		DependencyIndexes: file_campusapis_schooltime_v1_schoolTime_proto_depIdxs,
		MessageInfos:      file_campusapis_schooltime_v1_schoolTime_proto_msgTypes,
	}.Build()
	File_campusapis_schooltime_v1_schoolTime_proto = out.File
	file_campusapis_schooltime_v1_schoolTime_proto_rawDesc = nil
	file_campusapis_schooltime_v1_schoolTime_proto_goTypes = nil
	file_campusapis_schooltime_v1_schoolTime_proto_depIdxs = nil
}
