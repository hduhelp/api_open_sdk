// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: campusapis/schoolTime/v1/school_time.proto

package schoolTimev1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

	Error int32       `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg   string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data  *SchoolTime `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetSchoolTimeResponse) Reset() {
	*x = GetSchoolTimeResponse{}
	mi := &file_campusapis_schoolTime_v1_school_time_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSchoolTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchoolTimeResponse) ProtoMessage() {}

func (x *GetSchoolTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schoolTime_v1_school_time_proto_msgTypes[0]
	if x != nil {
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
	return file_campusapis_schoolTime_v1_school_time_proto_rawDescGZIP(), []int{0}
}

func (x *GetSchoolTimeResponse) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *GetSchoolTimeResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetSchoolTimeResponse) GetData() *SchoolTime {
	if x != nil {
		return x.Data
	}
	return nil
}

type SchoolTime struct {
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

func (x *SchoolTime) Reset() {
	*x = SchoolTime{}
	mi := &file_campusapis_schoolTime_v1_school_time_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SchoolTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchoolTime) ProtoMessage() {}

func (x *SchoolTime) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schoolTime_v1_school_time_proto_msgTypes[1]
	if x != nil {
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
	return file_campusapis_schoolTime_v1_school_time_proto_rawDescGZIP(), []int{1}
}

func (x *SchoolTime) GetSchoolYear() string {
	if x != nil {
		return x.SchoolYear
	}
	return ""
}

func (x *SchoolTime) GetSection() int32 {
	if x != nil {
		return x.Section
	}
	return 0
}

func (x *SchoolTime) GetSemester() string {
	if x != nil {
		return x.Semester
	}
	return ""
}

func (x *SchoolTime) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

func (x *SchoolTime) GetWeekDayNow() int32 {
	if x != nil {
		return x.WeekDayNow
	}
	return 0
}

func (x *SchoolTime) GetWeekNow() int32 {
	if x != nil {
		return x.WeekNow
	}
	return 0
}

type GetSemesterListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate string `protobuf:"bytes,1,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate   string `protobuf:"bytes,2,opt,name=endDate,proto3" json:"endDate,omitempty"`
}

func (x *GetSemesterListRequest) Reset() {
	*x = GetSemesterListRequest{}
	mi := &file_campusapis_schoolTime_v1_school_time_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSemesterListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSemesterListRequest) ProtoMessage() {}

func (x *GetSemesterListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schoolTime_v1_school_time_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSemesterListRequest.ProtoReflect.Descriptor instead.
func (*GetSemesterListRequest) Descriptor() ([]byte, []int) {
	return file_campusapis_schoolTime_v1_school_time_proto_rawDescGZIP(), []int{2}
}

func (x *GetSemesterListRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetSemesterListRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type GetSemesterListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error int32           `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg   string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data  []*SemesterInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetSemesterListResponse) Reset() {
	*x = GetSemesterListResponse{}
	mi := &file_campusapis_schoolTime_v1_school_time_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSemesterListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSemesterListResponse) ProtoMessage() {}

func (x *GetSemesterListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schoolTime_v1_school_time_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSemesterListResponse.ProtoReflect.Descriptor instead.
func (*GetSemesterListResponse) Descriptor() ([]byte, []int) {
	return file_campusapis_schoolTime_v1_school_time_proto_rawDescGZIP(), []int{3}
}

func (x *GetSemesterListResponse) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *GetSemesterListResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetSemesterListResponse) GetData() []*SemesterInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type SemesterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchoolYear string `protobuf:"bytes,1,opt,name=schoolYear,proto3" json:"schoolYear,omitempty"`
	Semester   string `protobuf:"bytes,2,opt,name=semester,proto3" json:"semester,omitempty"`
	StartDate  string `protobuf:"bytes,3,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate    string `protobuf:"bytes,4,opt,name=endDate,proto3" json:"endDate,omitempty"`
	State      int32  `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *SemesterInfo) Reset() {
	*x = SemesterInfo{}
	mi := &file_campusapis_schoolTime_v1_school_time_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SemesterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SemesterInfo) ProtoMessage() {}

func (x *SemesterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_schoolTime_v1_school_time_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SemesterInfo.ProtoReflect.Descriptor instead.
func (*SemesterInfo) Descriptor() ([]byte, []int) {
	return file_campusapis_schoolTime_v1_school_time_proto_rawDescGZIP(), []int{4}
}

func (x *SemesterInfo) GetSchoolYear() string {
	if x != nil {
		return x.SchoolYear
	}
	return ""
}

func (x *SemesterInfo) GetSemester() string {
	if x != nil {
		return x.Semester
	}
	return ""
}

func (x *SemesterInfo) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *SemesterInfo) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *SemesterInfo) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

var File_campusapis_schoolTime_v1_school_time_proto protoreflect.FileDescriptor

var file_campusapis_schoolTime_v1_school_time_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x61,
	0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x79, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xba, 0x01, 0x0a,
	0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
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
	0x52, 0x07, 0x77, 0x65, 0x65, 0x6b, 0x4e, 0x6f, 0x77, 0x22, 0x50, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x7d, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3a,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x98, 0x01, 0x0a, 0x0c, 0x53,
	0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x32, 0xd0, 0x02, 0x0a, 0x11, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2f, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x5a, 0x07, 0x12,
	0x05, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x13, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x74,
	0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0xba, 0x01, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x30, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x31, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x5a, 0x16, 0x12, 0x14,
	0x2f, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x42, 0xf8, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x53, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68, 0x65, 0x6c, 0x70,
	0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x61,
	0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x18, 0x43, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69,
	0x73, 0x5c, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x24, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x53, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61,
	0x70, 0x69, 0x73, 0x3a, 0x3a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_campusapis_schoolTime_v1_school_time_proto_rawDescOnce sync.Once
	file_campusapis_schoolTime_v1_school_time_proto_rawDescData = file_campusapis_schoolTime_v1_school_time_proto_rawDesc
)

func file_campusapis_schoolTime_v1_school_time_proto_rawDescGZIP() []byte {
	file_campusapis_schoolTime_v1_school_time_proto_rawDescOnce.Do(func() {
		file_campusapis_schoolTime_v1_school_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_campusapis_schoolTime_v1_school_time_proto_rawDescData)
	})
	return file_campusapis_schoolTime_v1_school_time_proto_rawDescData
}

var file_campusapis_schoolTime_v1_school_time_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_campusapis_schoolTime_v1_school_time_proto_goTypes = []any{
	(*GetSchoolTimeResponse)(nil),   // 0: campusapis.schoolTime.v1.GetSchoolTimeResponse
	(*SchoolTime)(nil),              // 1: campusapis.schoolTime.v1.SchoolTime
	(*GetSemesterListRequest)(nil),  // 2: campusapis.schoolTime.v1.GetSemesterListRequest
	(*GetSemesterListResponse)(nil), // 3: campusapis.schoolTime.v1.GetSemesterListResponse
	(*SemesterInfo)(nil),            // 4: campusapis.schoolTime.v1.SemesterInfo
	(*emptypb.Empty)(nil),           // 5: google.protobuf.Empty
}
var file_campusapis_schoolTime_v1_school_time_proto_depIdxs = []int32{
	1, // 0: campusapis.schoolTime.v1.GetSchoolTimeResponse.data:type_name -> campusapis.schoolTime.v1.SchoolTime
	4, // 1: campusapis.schoolTime.v1.GetSemesterListResponse.data:type_name -> campusapis.schoolTime.v1.SemesterInfo
	5, // 2: campusapis.schoolTime.v1.SchoolTimeService.GetSchoolTime:input_type -> google.protobuf.Empty
	2, // 3: campusapis.schoolTime.v1.SchoolTimeService.GetSemesterList:input_type -> campusapis.schoolTime.v1.GetSemesterListRequest
	0, // 4: campusapis.schoolTime.v1.SchoolTimeService.GetSchoolTime:output_type -> campusapis.schoolTime.v1.GetSchoolTimeResponse
	3, // 5: campusapis.schoolTime.v1.SchoolTimeService.GetSemesterList:output_type -> campusapis.schoolTime.v1.GetSemesterListResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_campusapis_schoolTime_v1_school_time_proto_init() }
func file_campusapis_schoolTime_v1_school_time_proto_init() {
	if File_campusapis_schoolTime_v1_school_time_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_campusapis_schoolTime_v1_school_time_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_campusapis_schoolTime_v1_school_time_proto_goTypes,
		DependencyIndexes: file_campusapis_schoolTime_v1_school_time_proto_depIdxs,
		MessageInfos:      file_campusapis_schoolTime_v1_school_time_proto_msgTypes,
	}.Build()
	File_campusapis_schoolTime_v1_school_time_proto = out.File
	file_campusapis_schoolTime_v1_school_time_proto_rawDesc = nil
	file_campusapis_schoolTime_v1_school_time_proto_goTypes = nil
	file_campusapis_schoolTime_v1_school_time_proto_depIdxs = nil
}
