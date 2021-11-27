// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: teaching/teaching.proto

package teaching

import (
	baseStaff "github.com/hduhelp/api_open_sdk/baseStaff"
	schoolTime "github.com/hduhelp/api_open_sdk/schoolTime"
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

type ScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Staff      *baseStaff.Staff `protobuf:"bytes,1,opt,name=Staff,proto3" json:"Staff,omitempty"`
	Semester   int32            `protobuf:"varint,2,opt,name=Semester,proto3" json:"Semester,omitempty"`
	SchoolYear int32            `protobuf:"varint,3,opt,name=SchoolYear,proto3" json:"SchoolYear,omitempty"`
	Timestamp  int64            `protobuf:"varint,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Week       int32            `protobuf:"varint,5,opt,name=Week,proto3" json:"Week,omitempty"`
	Weekday    int32            `protobuf:"varint,6,opt,name=Weekday,proto3" json:"Weekday,omitempty"`
}

func (x *ScheduleRequest) Reset() {
	*x = ScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teaching_teaching_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleRequest) ProtoMessage() {}

func (x *ScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teaching_teaching_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleRequest.ProtoReflect.Descriptor instead.
func (*ScheduleRequest) Descriptor() ([]byte, []int) {
	return file_teaching_teaching_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleRequest) GetStaff() *baseStaff.Staff {
	if x != nil {
		return x.Staff
	}
	return nil
}

func (x *ScheduleRequest) GetSemester() int32 {
	if x != nil {
		return x.Semester
	}
	return 0
}

func (x *ScheduleRequest) GetSchoolYear() int32 {
	if x != nil {
		return x.SchoolYear
	}
	return 0
}

func (x *ScheduleRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ScheduleRequest) GetWeek() int32 {
	if x != nil {
		return x.Week
	}
	return 0
}

func (x *ScheduleRequest) GetWeekday() int32 {
	if x != nil {
		return x.Weekday
	}
	return 0
}

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items map[string]*CourseItem `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teaching_teaching_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_teaching_teaching_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_teaching_teaching_proto_rawDescGZIP(), []int{1}
}

func (x *Course) GetItems() map[string]*CourseItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CourseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassID    string                 `protobuf:"bytes,1,opt,name=ClassID,proto3" json:"ClassID,omitempty"`
	ClassName  string                 `protobuf:"bytes,2,opt,name=ClassName,proto3" json:"ClassName,omitempty"`
	CourseID   string                 `protobuf:"bytes,3,opt,name=CourseID,proto3" json:"CourseID,omitempty"`
	CourseName string                 `protobuf:"bytes,4,opt,name=CourseName,proto3" json:"CourseName,omitempty"`
	ClassTime  string                 `protobuf:"bytes,5,opt,name=ClassTime,proto3" json:"ClassTime,omitempty"`
	Credit     string                 `protobuf:"bytes,6,opt,name=Credit,proto3" json:"Credit,omitempty"`
	SchoolYear *schoolTime.SchoolYear `protobuf:"bytes,7,opt,name=SchoolYear,proto3" json:"SchoolYear,omitempty"`
	Semester   *schoolTime.Semester   `protobuf:"bytes,8,opt,name=Semester,proto3" json:"Semester,omitempty"`
	Schedule   *Schedule              `protobuf:"bytes,9,opt,name=Schedule,proto3" json:"Schedule,omitempty"`
	IsThisWeek bool                   `protobuf:"varint,10,opt,name=IsThisWeek,proto3" json:"IsThisWeek,omitempty"`
}

func (x *CourseItem) Reset() {
	*x = CourseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teaching_teaching_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseItem) ProtoMessage() {}

func (x *CourseItem) ProtoReflect() protoreflect.Message {
	mi := &file_teaching_teaching_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseItem.ProtoReflect.Descriptor instead.
func (*CourseItem) Descriptor() ([]byte, []int) {
	return file_teaching_teaching_proto_rawDescGZIP(), []int{2}
}

func (x *CourseItem) GetClassID() string {
	if x != nil {
		return x.ClassID
	}
	return ""
}

func (x *CourseItem) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *CourseItem) GetCourseID() string {
	if x != nil {
		return x.CourseID
	}
	return ""
}

func (x *CourseItem) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

func (x *CourseItem) GetClassTime() string {
	if x != nil {
		return x.ClassTime
	}
	return ""
}

func (x *CourseItem) GetCredit() string {
	if x != nil {
		return x.Credit
	}
	return ""
}

func (x *CourseItem) GetSchoolYear() *schoolTime.SchoolYear {
	if x != nil {
		return x.SchoolYear
	}
	return nil
}

func (x *CourseItem) GetSemester() *schoolTime.Semester {
	if x != nil {
		return x.Semester
	}
	return nil
}

func (x *CourseItem) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *CourseItem) GetIsThisWeek() bool {
	if x != nil {
		return x.IsThisWeek
	}
	return false
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items map[string]*ScheduleItem `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teaching_teaching_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_teaching_teaching_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_teaching_teaching_proto_rawDescGZIP(), []int{3}
}

func (x *Schedule) GetItems() map[string]*ScheduleItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ScheduleItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Week      []int32                `protobuf:"varint,1,rep,packed,name=Week,proto3" json:"Week,omitempty"`
	WeekDay   int32                  `protobuf:"varint,2,opt,name=WeekDay,proto3" json:"WeekDay,omitempty"`
	Section   []int32                `protobuf:"varint,3,rep,packed,name=Section,proto3" json:"Section,omitempty"`
	StartTime int64                  `protobuf:"varint,4,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime   int64                  `protobuf:"varint,5,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	Teachers  *baseStaff.InfoMapList `protobuf:"bytes,6,opt,name=Teachers,proto3" json:"Teachers,omitempty"`
	Students  *baseStaff.InfoMapList `protobuf:"bytes,7,opt,name=Students,proto3" json:"Students,omitempty"`
	Location  string                 `protobuf:"bytes,8,opt,name=Location,proto3" json:"Location,omitempty"`
	SeatsNum  int32                  `protobuf:"varint,9,opt,name=SeatsNum,proto3" json:"SeatsNum,omitempty"`
	Commit    string                 `protobuf:"bytes,10,opt,name=Commit,proto3" json:"Commit,omitempty"`
}

func (x *ScheduleItem) Reset() {
	*x = ScheduleItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teaching_teaching_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleItem) ProtoMessage() {}

func (x *ScheduleItem) ProtoReflect() protoreflect.Message {
	mi := &file_teaching_teaching_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleItem.ProtoReflect.Descriptor instead.
func (*ScheduleItem) Descriptor() ([]byte, []int) {
	return file_teaching_teaching_proto_rawDescGZIP(), []int{4}
}

func (x *ScheduleItem) GetWeek() []int32 {
	if x != nil {
		return x.Week
	}
	return nil
}

func (x *ScheduleItem) GetWeekDay() int32 {
	if x != nil {
		return x.WeekDay
	}
	return 0
}

func (x *ScheduleItem) GetSection() []int32 {
	if x != nil {
		return x.Section
	}
	return nil
}

func (x *ScheduleItem) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *ScheduleItem) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ScheduleItem) GetTeachers() *baseStaff.InfoMapList {
	if x != nil {
		return x.Teachers
	}
	return nil
}

func (x *ScheduleItem) GetStudents() *baseStaff.InfoMapList {
	if x != nil {
		return x.Students
	}
	return nil
}

func (x *ScheduleItem) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ScheduleItem) GetSeatsNum() int32 {
	if x != nil {
		return x.SeatsNum
	}
	return 0
}

func (x *ScheduleItem) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

type StaffInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffID   string `protobuf:"bytes,1,opt,name=StaffID,proto3" json:"StaffID,omitempty"`
	StaffName string `protobuf:"bytes,2,opt,name=StaffName,proto3" json:"StaffName,omitempty"`
	Commit    string `protobuf:"bytes,3,opt,name=Commit,proto3" json:"Commit,omitempty"`
}

func (x *StaffInfo) Reset() {
	*x = StaffInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teaching_teaching_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffInfo) ProtoMessage() {}

func (x *StaffInfo) ProtoReflect() protoreflect.Message {
	mi := &file_teaching_teaching_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffInfo.ProtoReflect.Descriptor instead.
func (*StaffInfo) Descriptor() ([]byte, []int) {
	return file_teaching_teaching_proto_rawDescGZIP(), []int{5}
}

func (x *StaffInfo) GetStaffID() string {
	if x != nil {
		return x.StaffID
	}
	return ""
}

func (x *StaffInfo) GetStaffName() string {
	if x != nil {
		return x.StaffName
	}
	return ""
}

func (x *StaffInfo) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

var File_teaching_teaching_proto protoreflect.FileDescriptor

var file_teaching_teaching_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x1a, 0x19, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x0f,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x66, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x52, 0x05, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59,
	0x65, 0x61, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x65, 0x65, 0x6b, 0x64, 0x61, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x57, 0x65, 0x65, 0x6b, 0x64, 0x61, 0x79, 0x22,
	0x8b, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x4e, 0x0a,
	0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf0, 0x02,
	0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x59, 0x65, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65,
	0x61, 0x72, 0x52, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x12, 0x30,
	0x0a, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x53, 0x65,
	0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x2e, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x54, 0x68, 0x69, 0x73, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x54, 0x68, 0x69, 0x73, 0x57, 0x65, 0x65, 0x6b,
	0x22, 0x91, 0x01, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x33, 0x0a,
	0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x1a, 0x50, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xc6, 0x02, 0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x04, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x65, 0x65,
	0x6b, 0x44, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x57, 0x65, 0x65, 0x6b,
	0x44, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x08, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x08, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x61,
	0x74, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x65, 0x61,
	0x74, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x5b, 0x0a,
	0x09, 0x53, 0x74, 0x61, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x66, 0x66, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x32, 0x5a, 0x0a, 0x0e, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x17,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68, 0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69,
	0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teaching_teaching_proto_rawDescOnce sync.Once
	file_teaching_teaching_proto_rawDescData = file_teaching_teaching_proto_rawDesc
)

func file_teaching_teaching_proto_rawDescGZIP() []byte {
	file_teaching_teaching_proto_rawDescOnce.Do(func() {
		file_teaching_teaching_proto_rawDescData = protoimpl.X.CompressGZIP(file_teaching_teaching_proto_rawDescData)
	})
	return file_teaching_teaching_proto_rawDescData
}

var file_teaching_teaching_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_teaching_teaching_proto_goTypes = []interface{}{
	(*ScheduleRequest)(nil),       // 0: teaching.ScheduleRequest
	(*Course)(nil),                // 1: teaching.Course
	(*CourseItem)(nil),            // 2: teaching.CourseItem
	(*Schedule)(nil),              // 3: teaching.Schedule
	(*ScheduleItem)(nil),          // 4: teaching.ScheduleItem
	(*StaffInfo)(nil),             // 5: teaching.StaffInfo
	nil,                           // 6: teaching.Course.ItemsEntry
	nil,                           // 7: teaching.Schedule.ItemsEntry
	(*baseStaff.Staff)(nil),       // 8: baseStaff.Staff
	(*schoolTime.SchoolYear)(nil), // 9: schoolTime.SchoolYear
	(*schoolTime.Semester)(nil),   // 10: schoolTime.Semester
	(*baseStaff.InfoMapList)(nil), // 11: baseStaff.InfoMapList
}
var file_teaching_teaching_proto_depIdxs = []int32{
	8,  // 0: teaching.ScheduleRequest.Staff:type_name -> baseStaff.Staff
	6,  // 1: teaching.Course.Items:type_name -> teaching.Course.ItemsEntry
	9,  // 2: teaching.CourseItem.SchoolYear:type_name -> schoolTime.SchoolYear
	10, // 3: teaching.CourseItem.Semester:type_name -> schoolTime.Semester
	3,  // 4: teaching.CourseItem.Schedule:type_name -> teaching.Schedule
	7,  // 5: teaching.Schedule.Items:type_name -> teaching.Schedule.ItemsEntry
	11, // 6: teaching.ScheduleItem.Teachers:type_name -> baseStaff.InfoMapList
	11, // 7: teaching.ScheduleItem.Students:type_name -> baseStaff.InfoMapList
	2,  // 8: teaching.Course.ItemsEntry.value:type_name -> teaching.CourseItem
	4,  // 9: teaching.Schedule.ItemsEntry.value:type_name -> teaching.ScheduleItem
	0,  // 10: teaching.TeachingServer.StaffGetScheduleHandler:input_type -> teaching.ScheduleRequest
	1,  // 11: teaching.TeachingServer.StaffGetScheduleHandler:output_type -> teaching.Course
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_teaching_teaching_proto_init() }
func file_teaching_teaching_proto_init() {
	if File_teaching_teaching_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teaching_teaching_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleRequest); i {
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
		file_teaching_teaching_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
		file_teaching_teaching_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseItem); i {
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
		file_teaching_teaching_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_teaching_teaching_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleItem); i {
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
		file_teaching_teaching_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffInfo); i {
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
			RawDescriptor: file_teaching_teaching_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teaching_teaching_proto_goTypes,
		DependencyIndexes: file_teaching_teaching_proto_depIdxs,
		MessageInfos:      file_teaching_teaching_proto_msgTypes,
	}.Build()
	File_teaching_teaching_proto = out.File
	file_teaching_teaching_proto_rawDesc = nil
	file_teaching_teaching_proto_goTypes = nil
	file_teaching_teaching_proto_depIdxs = nil
}
