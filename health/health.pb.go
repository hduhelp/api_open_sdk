// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: health/health.proto

package health

import (
	baseStaff "github.com/hduhelp/api_open_sdk/baseStaff"
	types "github.com/hduhelp/api_open_sdk/types"
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

type StaffIDWithTimestampRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffID string `protobuf:"bytes,1,opt,name=StaffID,proto3" json:"StaffID,omitempty"`
	Date    string `protobuf:"bytes,2,opt,name=Date,proto3" json:"Date,omitempty"`
}

func (x *StaffIDWithTimestampRequest) Reset() {
	*x = StaffIDWithTimestampRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffIDWithTimestampRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffIDWithTimestampRequest) ProtoMessage() {}

func (x *StaffIDWithTimestampRequest) ProtoReflect() protoreflect.Message {
	mi := &file_health_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffIDWithTimestampRequest.ProtoReflect.Descriptor instead.
func (*StaffIDWithTimestampRequest) Descriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{0}
}

func (x *StaffIDWithTimestampRequest) GetStaffID() string {
	if x != nil {
		return x.StaffID
	}
	return ""
}

func (x *StaffIDWithTimestampRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type CheckinRecords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CheckinRecord `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *CheckinRecords) Reset() {
	*x = CheckinRecords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckinRecords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinRecords) ProtoMessage() {}

func (x *CheckinRecords) ProtoReflect() protoreflect.Message {
	mi := &file_health_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinRecords.ProtoReflect.Descriptor instead.
func (*CheckinRecords) Descriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{1}
}

func (x *CheckinRecords) GetItems() []*CheckinRecord {
	if x != nil {
		return x.Items
	}
	return nil
}

//健康打卡记录
type CheckinRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int64            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Creator    string           `protobuf:"bytes,2,opt,name=Creator,proto3" json:"Creator,omitempty"`
	CreatedAt  *types.Timestamp `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Modified   *types.Timestamp `protobuf:"bytes,4,opt,name=Modified,proto3" json:"Modified,omitempty"`
	Modifier   string           `protobuf:"bytes,5,opt,name=Modifier,proto3" json:"Modifier,omitempty"`
	Content    *types.JSON      `protobuf:"bytes,6,opt,name=Content,proto3" json:"Content,omitempty"` //健康打卡内容JSON
	ReportTime *types.Timestamp `protobuf:"bytes,7,opt,name=ReportTime,proto3" json:"ReportTime,omitempty"`
	Province   string           `protobuf:"bytes,8,opt,name=Province,proto3" json:"Province,omitempty"`
	City       string           `protobuf:"bytes,9,opt,name=City,proto3" json:"City,omitempty"`
	Country    string           `protobuf:"bytes,10,opt,name=Country,proto3" json:"Country,omitempty"`
	StaffID    string           `protobuf:"bytes,11,opt,name=StaffID,proto3" json:"StaffID,omitempty"`
	StaffName  string           `protobuf:"bytes,12,opt,name=StaffName,proto3" json:"StaffName,omitempty"`
	StaffType  string           `protobuf:"bytes,13,opt,name=StaffType,proto3" json:"StaffType,omitempty"`
	Unit       string           `protobuf:"bytes,14,opt,name=Unit,proto3" json:"Unit,omitempty"`
	SchoolID   string           `protobuf:"bytes,15,opt,name=SchoolID,proto3" json:"SchoolID,omitempty"`
}

func (x *CheckinRecord) Reset() {
	*x = CheckinRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckinRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinRecord) ProtoMessage() {}

func (x *CheckinRecord) ProtoReflect() protoreflect.Message {
	mi := &file_health_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinRecord.ProtoReflect.Descriptor instead.
func (*CheckinRecord) Descriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{2}
}

func (x *CheckinRecord) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CheckinRecord) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *CheckinRecord) GetCreatedAt() *types.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CheckinRecord) GetModified() *types.Timestamp {
	if x != nil {
		return x.Modified
	}
	return nil
}

func (x *CheckinRecord) GetModifier() string {
	if x != nil {
		return x.Modifier
	}
	return ""
}

func (x *CheckinRecord) GetContent() *types.JSON {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *CheckinRecord) GetReportTime() *types.Timestamp {
	if x != nil {
		return x.ReportTime
	}
	return nil
}

func (x *CheckinRecord) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *CheckinRecord) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CheckinRecord) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CheckinRecord) GetStaffID() string {
	if x != nil {
		return x.StaffID
	}
	return ""
}

func (x *CheckinRecord) GetStaffName() string {
	if x != nil {
		return x.StaffName
	}
	return ""
}

func (x *CheckinRecord) GetStaffType() string {
	if x != nil {
		return x.StaffType
	}
	return ""
}

func (x *CheckinRecord) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *CheckinRecord) GetSchoolID() string {
	if x != nil {
		return x.SchoolID
	}
	return ""
}

//健康码
type HealthCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodeID     string           `protobuf:"bytes,1,opt,name=CodeID,proto3" json:"CodeID,omitempty"`
	ApplyTime  *types.Timestamp `protobuf:"bytes,2,opt,name=ApplyTime,proto3" json:"ApplyTime,omitempty"`
	GrantTime  *types.Timestamp `protobuf:"bytes,3,opt,name=GrantTime,proto3" json:"GrantTime,omitempty"`
	UpdateTime *types.Timestamp `protobuf:"bytes,4,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	Name       string           `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	Region     string           `protobuf:"bytes,6,opt,name=Region,proto3" json:"Region,omitempty"`
	Status     string           `protobuf:"bytes,7,opt,name=Status,proto3" json:"Status,omitempty"`
	Comment    string           `protobuf:"bytes,8,opt,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *HealthCode) Reset() {
	*x = HealthCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_health_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCode) ProtoMessage() {}

func (x *HealthCode) ProtoReflect() protoreflect.Message {
	mi := &file_health_health_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCode.ProtoReflect.Descriptor instead.
func (*HealthCode) Descriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{3}
}

func (x *HealthCode) GetCodeID() string {
	if x != nil {
		return x.CodeID
	}
	return ""
}

func (x *HealthCode) GetApplyTime() *types.Timestamp {
	if x != nil {
		return x.ApplyTime
	}
	return nil
}

func (x *HealthCode) GetGrantTime() *types.Timestamp {
	if x != nil {
		return x.GrantTime
	}
	return nil
}

func (x *HealthCode) GetUpdateTime() *types.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *HealthCode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HealthCode) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *HealthCode) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *HealthCode) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

//健康码老
type StudentHealthCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffId          string `protobuf:"bytes,1,opt,name=StaffId,proto3" json:"StaffId,omitempty"`
	StaffName        string `protobuf:"bytes,2,opt,name=StaffName,proto3" json:"StaffName,omitempty"`
	PhoneNum         string `protobuf:"bytes,3,opt,name=PhoneNum,proto3" json:"PhoneNum,omitempty"`
	CodeStatus       string `protobuf:"bytes,4,opt,name=CodeStatus,proto3" json:"CodeStatus,omitempty"`
	Location         string `protobuf:"bytes,5,opt,name=Location,proto3" json:"Location,omitempty"`
	Reason           string `protobuf:"bytes,6,opt,name=Reason,proto3" json:"Reason,omitempty"`
	FirstApplyTime   string `protobuf:"bytes,7,opt,name=FirstApplyTime,proto3" json:"FirstApplyTime,omitempty"`
	FirstGrantTime   string `protobuf:"bytes,8,opt,name=FirstGrantTime,proto3" json:"FirstGrantTime,omitempty"`
	LastUpdateTime   string `protobuf:"bytes,9,opt,name=LastUpdateTime,proto3" json:"LastUpdateTime,omitempty"`
	ServerUpdateTime string `protobuf:"bytes,10,opt,name=ServerUpdateTime,proto3" json:"ServerUpdateTime,omitempty"`
}

func (x *StudentHealthCode) Reset() {
	*x = StudentHealthCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_health_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentHealthCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentHealthCode) ProtoMessage() {}

func (x *StudentHealthCode) ProtoReflect() protoreflect.Message {
	mi := &file_health_health_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentHealthCode.ProtoReflect.Descriptor instead.
func (*StudentHealthCode) Descriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{4}
}

func (x *StudentHealthCode) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

func (x *StudentHealthCode) GetStaffName() string {
	if x != nil {
		return x.StaffName
	}
	return ""
}

func (x *StudentHealthCode) GetPhoneNum() string {
	if x != nil {
		return x.PhoneNum
	}
	return ""
}

func (x *StudentHealthCode) GetCodeStatus() string {
	if x != nil {
		return x.CodeStatus
	}
	return ""
}

func (x *StudentHealthCode) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *StudentHealthCode) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *StudentHealthCode) GetFirstApplyTime() string {
	if x != nil {
		return x.FirstApplyTime
	}
	return ""
}

func (x *StudentHealthCode) GetFirstGrantTime() string {
	if x != nil {
		return x.FirstGrantTime
	}
	return ""
}

func (x *StudentHealthCode) GetLastUpdateTime() string {
	if x != nil {
		return x.LastUpdateTime
	}
	return ""
}

func (x *StudentHealthCode) GetServerUpdateTime() string {
	if x != nil {
		return x.ServerUpdateTime
	}
	return ""
}

var File_health_health_proto protoreflect.FileDescriptor

var file_health_health_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x1a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68, 0x65, 0x6c,
	0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75,
	0x68, 0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64,
	0x6b, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x1b, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x57, 0x69, 0x74, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x22, 0x3d, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xdc, 0x03, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x25, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4a, 0x53, 0x4f, 0x4e, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x53, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x49, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x49, 0x44, 0x22, 0x94, 0x02, 0x0a, 0x0a, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x2e, 0x0a,
	0x09, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a,
	0x09, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xdf, 0x02,
	0x0a, 0x11, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x53, 0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x53, 0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x64, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x64,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x46,
	0x69, 0x72, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x46, 0x69, 0x72, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x46, 0x69, 0x72, 0x73, 0x74, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x4c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32,
	0x9d, 0x02, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x57, 0x69, 0x74, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69,
	0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x1a, 0x19, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x6f, 0x64, 0x65, 0x56, 0x32, 0x12, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x1a, 0x12, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64,
	0x75, 0x68, 0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73,
	0x64, 0x6b, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_health_health_proto_rawDescOnce sync.Once
	file_health_health_proto_rawDescData = file_health_health_proto_rawDesc
)

func file_health_health_proto_rawDescGZIP() []byte {
	file_health_health_proto_rawDescOnce.Do(func() {
		file_health_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_health_proto_rawDescData)
	})
	return file_health_health_proto_rawDescData
}

var file_health_health_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_health_health_proto_goTypes = []interface{}{
	(*StaffIDWithTimestampRequest)(nil), // 0: health.StaffIDWithTimestampRequest
	(*CheckinRecords)(nil),              // 1: health.CheckinRecords
	(*CheckinRecord)(nil),               // 2: health.CheckinRecord
	(*HealthCode)(nil),                  // 3: health.HealthCode
	(*StudentHealthCode)(nil),           // 4: health.StudentHealthCode
	(*types.Timestamp)(nil),             // 5: types.Timestamp
	(*types.JSON)(nil),                  // 6: types.JSON
	(*baseStaff.Staff)(nil),             // 7: baseStaff.Staff
}
var file_health_health_proto_depIdxs = []int32{
	2,  // 0: health.CheckinRecords.Items:type_name -> health.CheckinRecord
	5,  // 1: health.CheckinRecord.CreatedAt:type_name -> types.Timestamp
	5,  // 2: health.CheckinRecord.Modified:type_name -> types.Timestamp
	6,  // 3: health.CheckinRecord.Content:type_name -> types.JSON
	5,  // 4: health.CheckinRecord.ReportTime:type_name -> types.Timestamp
	5,  // 5: health.HealthCode.ApplyTime:type_name -> types.Timestamp
	5,  // 6: health.HealthCode.GrantTime:type_name -> types.Timestamp
	5,  // 7: health.HealthCode.UpdateTime:type_name -> types.Timestamp
	0,  // 8: health.HealthService.GetCheckinRecord:input_type -> health.StaffIDWithTimestampRequest
	7,  // 9: health.HealthService.GetCheckinRecords:input_type -> baseStaff.Staff
	7,  // 10: health.HealthService.GetHealthCode:input_type -> baseStaff.Staff
	7,  // 11: health.HealthService.GetHealthCodeV2:input_type -> baseStaff.Staff
	2,  // 12: health.HealthService.GetCheckinRecord:output_type -> health.CheckinRecord
	1,  // 13: health.HealthService.GetCheckinRecords:output_type -> health.CheckinRecords
	4,  // 14: health.HealthService.GetHealthCode:output_type -> health.StudentHealthCode
	3,  // 15: health.HealthService.GetHealthCodeV2:output_type -> health.HealthCode
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_health_health_proto_init() }
func file_health_health_proto_init() {
	if File_health_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_health_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffIDWithTimestampRequest); i {
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
		file_health_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckinRecords); i {
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
		file_health_health_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckinRecord); i {
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
		file_health_health_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCode); i {
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
		file_health_health_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentHealthCode); i {
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
			RawDescriptor: file_health_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_health_proto_goTypes,
		DependencyIndexes: file_health_health_proto_depIdxs,
		MessageInfos:      file_health_health_proto_msgTypes,
	}.Build()
	File_health_health_proto = out.File
	file_health_health_proto_rawDesc = nil
	file_health_health_proto_goTypes = nil
	file_health_health_proto_depIdxs = nil
}