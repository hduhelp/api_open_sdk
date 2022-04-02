// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: campusapis/staff/v1/freshman.proto

package staffv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

//新生三要素请求，依次判断（学号）、（录取通知书或准考证号）、（身份证号），要求其中最少两项不为空
type FreshmanKeywordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//学号
	StaffId string `protobuf:"bytes,1,opt,name=staffId,proto3" json:"staffId,omitempty"`
	//录取通知书号或者是准考证号，不同省份并不一样
	OfferId string `protobuf:"bytes,2,opt,name=offerId,proto3" json:"offerId,omitempty"`
	//身份证号
	IdCardNo string `protobuf:"bytes,3,opt,name=idCardNo,proto3" json:"idCardNo,omitempty"`
}

func (x *FreshmanKeywordRequest) Reset() {
	*x = FreshmanKeywordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_staff_v1_freshman_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreshmanKeywordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreshmanKeywordRequest) ProtoMessage() {}

func (x *FreshmanKeywordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_staff_v1_freshman_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreshmanKeywordRequest.ProtoReflect.Descriptor instead.
func (*FreshmanKeywordRequest) Descriptor() ([]byte, []int) {
	return file_campusapis_staff_v1_freshman_proto_rawDescGZIP(), []int{0}
}

func (x *FreshmanKeywordRequest) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

func (x *FreshmanKeywordRequest) GetOfferId() string {
	if x != nil {
		return x.OfferId
	}
	return ""
}

func (x *FreshmanKeywordRequest) GetIdCardNo() string {
	if x != nil {
		return x.IdCardNo
	}
	return ""
}

//新生基本信息
type GetFreshmanBaseInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//学号
	StaffId string `protobuf:"bytes,1,opt,name=staffId,proto3" json:"staffId,omitempty"`
	//姓名
	StaffName string `protobuf:"bytes,2,opt,name=staffName,proto3" json:"staffName,omitempty"`
	//录取通知书号或者是准考证号，不同省份并不一样
	OfferId string `protobuf:"bytes,3,opt,name=offerId,proto3" json:"offerId,omitempty"`
	//身份证号
	IdCardNo string `protobuf:"bytes,4,opt,name=idCardNo,proto3" json:"idCardNo,omitempty"`
	//录取时间
	OfferTime string `protobuf:"bytes,5,opt,name=offerTime,proto3" json:"offerTime,omitempty"`
	//省份
	Province string `protobuf:"bytes,6,opt,name=province,proto3" json:"province,omitempty"`
}

func (x *GetFreshmanBaseInfoResponse) Reset() {
	*x = GetFreshmanBaseInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_staff_v1_freshman_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFreshmanBaseInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFreshmanBaseInfoResponse) ProtoMessage() {}

func (x *GetFreshmanBaseInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_staff_v1_freshman_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFreshmanBaseInfoResponse.ProtoReflect.Descriptor instead.
func (*GetFreshmanBaseInfoResponse) Descriptor() ([]byte, []int) {
	return file_campusapis_staff_v1_freshman_proto_rawDescGZIP(), []int{1}
}

func (x *GetFreshmanBaseInfoResponse) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

func (x *GetFreshmanBaseInfoResponse) GetStaffName() string {
	if x != nil {
		return x.StaffName
	}
	return ""
}

func (x *GetFreshmanBaseInfoResponse) GetOfferId() string {
	if x != nil {
		return x.OfferId
	}
	return ""
}

func (x *GetFreshmanBaseInfoResponse) GetIdCardNo() string {
	if x != nil {
		return x.IdCardNo
	}
	return ""
}

func (x *GetFreshmanBaseInfoResponse) GetOfferTime() string {
	if x != nil {
		return x.OfferTime
	}
	return ""
}

func (x *GetFreshmanBaseInfoResponse) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

//新生详细信息
type GetFreshmanInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//学号
	StaffId string `protobuf:"bytes,1,opt,name=staffId,proto3" json:"staffId,omitempty"`
	//姓名
	StaffName string `protobuf:"bytes,2,opt,name=staffName,proto3" json:"staffName,omitempty"`
	//班级名称
	ClassName string `protobuf:"bytes,3,opt,name=className,proto3" json:"className,omitempty"`
	//宿舍房间名称
	RoomName string `protobuf:"bytes,4,opt,name=roomName,proto3" json:"roomName,omitempty"`
	//注册状态（是否入学）
	RegStatus string `protobuf:"bytes,5,opt,name=regStatus,proto3" json:"regStatus,omitempty"`
	//学院编号
	SchoolID string `protobuf:"bytes,6,opt,name=schoolID,proto3" json:"schoolID,omitempty"`
	//学院名称
	SchoolName string `protobuf:"bytes,7,opt,name=schoolName,proto3" json:"schoolName,omitempty"`
	//专业名称
	MajorName string `protobuf:"bytes,8,opt,name=majorName,proto3" json:"majorName,omitempty"`
	//身份证号
	CardID string `protobuf:"bytes,9,opt,name=cardID,proto3" json:"cardID,omitempty"`
	//辅导员姓名
	TeacherName  string `protobuf:"bytes,10,opt,name=teacherName,proto3" json:"teacherName,omitempty"`
	TeacherPhone string `protobuf:"bytes,11,opt,name=teacherPhone,proto3" json:"teacherPhone,omitempty"`
	//班助姓名
	AssistantName string `protobuf:"bytes,12,opt,name=assistantName,proto3" json:"assistantName,omitempty"`
	//班助电话
	AssistantPhone string `protobuf:"bytes,13,opt,name=assistantPhone,proto3" json:"assistantPhone,omitempty"`
	//是否支付学费
	HasPaid string `protobuf:"bytes,14,opt,name=hasPaid,proto3" json:"hasPaid,omitempty"`
	//绿色通道
	LSTD   string `protobuf:"bytes,15,opt,name=LSTD,proto3" json:"LSTD,omitempty"`
	LSTDXY string `protobuf:"bytes,16,opt,name=LSTDXY,proto3" json:"LSTDXY,omitempty"`
	LSTDZX string `protobuf:"bytes,17,opt,name=LSTDZX,proto3" json:"LSTDZX,omitempty"`
	// 电话
	Phone string `protobuf:"bytes,18,opt,name=phone,proto3" json:"phone,omitempty"`
	// QQ
	Qq string `protobuf:"bytes,19,opt,name=qq,proto3" json:"qq,omitempty"`
	//楼号
	RoomBuilding string `protobuf:"bytes,20,opt,name=roomBuilding,proto3" json:"roomBuilding,omitempty"`
	//寝室号
	RoomRoom string `protobuf:"bytes,21,opt,name=roomRoom,proto3" json:"roomRoom,omitempty"`
	//床号
	RoomBed string `protobuf:"bytes,22,opt,name=roomBed,proto3" json:"roomBed,omitempty"`
	//报到日期
	ComeDate string `protobuf:"bytes,23,opt,name=comeDate,proto3" json:"comeDate,omitempty"`
	// 身高（厘米）
	High string `protobuf:"bytes,24,opt,name=high,proto3" json:"high,omitempty"`
	//火车区间
	Station string `protobuf:"bytes,25,opt,name=station,proto3" json:"station,omitempty"`
}

func (x *GetFreshmanInfoResponse) Reset() {
	*x = GetFreshmanInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_staff_v1_freshman_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFreshmanInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFreshmanInfoResponse) ProtoMessage() {}

func (x *GetFreshmanInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_staff_v1_freshman_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFreshmanInfoResponse.ProtoReflect.Descriptor instead.
func (*GetFreshmanInfoResponse) Descriptor() ([]byte, []int) {
	return file_campusapis_staff_v1_freshman_proto_rawDescGZIP(), []int{2}
}

func (x *GetFreshmanInfoResponse) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetStaffName() string {
	if x != nil {
		return x.StaffName
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetRegStatus() string {
	if x != nil {
		return x.RegStatus
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetSchoolID() string {
	if x != nil {
		return x.SchoolID
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetSchoolName() string {
	if x != nil {
		return x.SchoolName
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetMajorName() string {
	if x != nil {
		return x.MajorName
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetCardID() string {
	if x != nil {
		return x.CardID
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetTeacherName() string {
	if x != nil {
		return x.TeacherName
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetTeacherPhone() string {
	if x != nil {
		return x.TeacherPhone
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetAssistantName() string {
	if x != nil {
		return x.AssistantName
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetAssistantPhone() string {
	if x != nil {
		return x.AssistantPhone
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetHasPaid() string {
	if x != nil {
		return x.HasPaid
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetLSTD() string {
	if x != nil {
		return x.LSTD
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetLSTDXY() string {
	if x != nil {
		return x.LSTDXY
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetLSTDZX() string {
	if x != nil {
		return x.LSTDZX
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetQq() string {
	if x != nil {
		return x.Qq
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetRoomBuilding() string {
	if x != nil {
		return x.RoomBuilding
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetRoomRoom() string {
	if x != nil {
		return x.RoomRoom
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetRoomBed() string {
	if x != nil {
		return x.RoomBed
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetComeDate() string {
	if x != nil {
		return x.ComeDate
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetHigh() string {
	if x != nil {
		return x.High
	}
	return ""
}

func (x *GetFreshmanInfoResponse) GetStation() string {
	if x != nil {
		return x.Station
	}
	return ""
}

//新生室友查询
type GetFreshmanRoommatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 室友列表
	Items []*GetFreshmanRoommatesResponse_Roommate `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetFreshmanRoommatesResponse) Reset() {
	*x = GetFreshmanRoommatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_staff_v1_freshman_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFreshmanRoommatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFreshmanRoommatesResponse) ProtoMessage() {}

func (x *GetFreshmanRoommatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_staff_v1_freshman_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFreshmanRoommatesResponse.ProtoReflect.Descriptor instead.
func (*GetFreshmanRoommatesResponse) Descriptor() ([]byte, []int) {
	return file_campusapis_staff_v1_freshman_proto_rawDescGZIP(), []int{3}
}

func (x *GetFreshmanRoommatesResponse) GetItems() []*GetFreshmanRoommatesResponse_Roommate {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetFreshmanRoommatesResponse_Roommate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//室友姓名
	StaffName string `protobuf:"bytes,1,opt,name=staffName,proto3" json:"staffName,omitempty"`
	Qq        string `protobuf:"bytes,2,opt,name=qq,proto3" json:"qq,omitempty"`
	Phone     string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetFreshmanRoommatesResponse_Roommate) Reset() {
	*x = GetFreshmanRoommatesResponse_Roommate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campusapis_staff_v1_freshman_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFreshmanRoommatesResponse_Roommate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFreshmanRoommatesResponse_Roommate) ProtoMessage() {}

func (x *GetFreshmanRoommatesResponse_Roommate) ProtoReflect() protoreflect.Message {
	mi := &file_campusapis_staff_v1_freshman_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFreshmanRoommatesResponse_Roommate.ProtoReflect.Descriptor instead.
func (*GetFreshmanRoommatesResponse_Roommate) Descriptor() ([]byte, []int) {
	return file_campusapis_staff_v1_freshman_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetFreshmanRoommatesResponse_Roommate) GetStaffName() string {
	if x != nil {
		return x.StaffName
	}
	return ""
}

func (x *GetFreshmanRoommatesResponse_Roommate) GetQq() string {
	if x != nil {
		return x.Qq
	}
	return ""
}

func (x *GetFreshmanRoommatesResponse_Roommate) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_campusapis_staff_v1_freshman_proto protoreflect.FileDescriptor

var file_campusapis_staff_v1_freshman_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x16, 0x46, 0x72, 0x65, 0x73, 0x68,
	0x6d, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4e,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4e,
	0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61,
	0x6e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x12,
	0x1c, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x22, 0xd7, 0x05, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x46, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x49,
	0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x73, 0x73,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68,
	0x61, 0x73, 0x50, 0x61, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61,
	0x73, 0x50, 0x61, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x53, 0x54, 0x44, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x53, 0x54, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x53, 0x54,
	0x44, 0x58, 0x59, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x53, 0x54, 0x44, 0x58,
	0x59, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x53, 0x54, 0x44, 0x5a, 0x58, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x4c, 0x53, 0x54, 0x44, 0x5a, 0x58, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x71, 0x71, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x71, 0x71, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x6f, 0x6f, 0x6d, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x6f, 0x6f, 0x6d, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x42, 0x65, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x42, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xc0, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x73, 0x68,
	0x6d, 0x61, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x6d, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65,
	0x73, 0x68, 0x6d, 0x61, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x6d, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x6d, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x4e, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x6d, 0x61,
	0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x71, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x71, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x32, 0x98, 0x04, 0x0a, 0x0f, 0x46, 0x72, 0x65, 0x73, 0x68,
	0x6d, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa7, 0x01, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x46, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2b, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61,
	0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61,
	0x6e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x17, 0x2f, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x5a, 0x10, 0x12, 0x0e, 0x2f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x73,
	0x68, 0x6d, 0x61, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75,
	0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x72, 0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x72, 0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x17, 0x2f, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x2f,
	0x69, 0x6e, 0x66, 0x6f, 0x5a, 0x10, 0x12, 0x0e, 0x2f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61,
	0x6e, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0xb8, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x72,
	0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x6d, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x2b, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x52, 0x6f,
	0x6f, 0x6d, 0x6d, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x40, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3a, 0x12, 0x1b, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x2f, 0x72, 0x6f, 0x6f, 0x6d,
	0x6d, 0x61, 0x74, 0x65, 0x5a, 0x14, 0x12, 0x12, 0x2f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x6d, 0x61,
	0x6e, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x6d, 0x61, 0x74, 0x65, 0x62, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x42, 0xd3, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x46,
	0x72, 0x65, 0x73, 0x68, 0x6d, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68, 0x65,
	0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f,
	0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x61, 0x66, 0x66, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53,
	0x58, 0xaa, 0x02, 0x13, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x61, 0x70, 0x69, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f,
	0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x15, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_campusapis_staff_v1_freshman_proto_rawDescOnce sync.Once
	file_campusapis_staff_v1_freshman_proto_rawDescData = file_campusapis_staff_v1_freshman_proto_rawDesc
)

func file_campusapis_staff_v1_freshman_proto_rawDescGZIP() []byte {
	file_campusapis_staff_v1_freshman_proto_rawDescOnce.Do(func() {
		file_campusapis_staff_v1_freshman_proto_rawDescData = protoimpl.X.CompressGZIP(file_campusapis_staff_v1_freshman_proto_rawDescData)
	})
	return file_campusapis_staff_v1_freshman_proto_rawDescData
}

var file_campusapis_staff_v1_freshman_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_campusapis_staff_v1_freshman_proto_goTypes = []interface{}{
	(*FreshmanKeywordRequest)(nil),                // 0: campusapis.staff.v1.FreshmanKeywordRequest
	(*GetFreshmanBaseInfoResponse)(nil),           // 1: campusapis.staff.v1.GetFreshmanBaseInfoResponse
	(*GetFreshmanInfoResponse)(nil),               // 2: campusapis.staff.v1.GetFreshmanInfoResponse
	(*GetFreshmanRoommatesResponse)(nil),          // 3: campusapis.staff.v1.GetFreshmanRoommatesResponse
	(*GetFreshmanRoommatesResponse_Roommate)(nil), // 4: campusapis.staff.v1.GetFreshmanRoommatesResponse.Roommate
}
var file_campusapis_staff_v1_freshman_proto_depIdxs = []int32{
	4, // 0: campusapis.staff.v1.GetFreshmanRoommatesResponse.items:type_name -> campusapis.staff.v1.GetFreshmanRoommatesResponse.Roommate
	0, // 1: campusapis.staff.v1.FreshmanService.GetFreshmanBaseInfo:input_type -> campusapis.staff.v1.FreshmanKeywordRequest
	0, // 2: campusapis.staff.v1.FreshmanService.GetFreshmanInfo:input_type -> campusapis.staff.v1.FreshmanKeywordRequest
	0, // 3: campusapis.staff.v1.FreshmanService.GetFreshmanRoommates:input_type -> campusapis.staff.v1.FreshmanKeywordRequest
	1, // 4: campusapis.staff.v1.FreshmanService.GetFreshmanBaseInfo:output_type -> campusapis.staff.v1.GetFreshmanBaseInfoResponse
	2, // 5: campusapis.staff.v1.FreshmanService.GetFreshmanInfo:output_type -> campusapis.staff.v1.GetFreshmanInfoResponse
	3, // 6: campusapis.staff.v1.FreshmanService.GetFreshmanRoommates:output_type -> campusapis.staff.v1.GetFreshmanRoommatesResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_campusapis_staff_v1_freshman_proto_init() }
func file_campusapis_staff_v1_freshman_proto_init() {
	if File_campusapis_staff_v1_freshman_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_campusapis_staff_v1_freshman_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreshmanKeywordRequest); i {
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
		file_campusapis_staff_v1_freshman_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFreshmanBaseInfoResponse); i {
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
		file_campusapis_staff_v1_freshman_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFreshmanInfoResponse); i {
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
		file_campusapis_staff_v1_freshman_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFreshmanRoommatesResponse); i {
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
		file_campusapis_staff_v1_freshman_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFreshmanRoommatesResponse_Roommate); i {
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
			RawDescriptor: file_campusapis_staff_v1_freshman_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_campusapis_staff_v1_freshman_proto_goTypes,
		DependencyIndexes: file_campusapis_staff_v1_freshman_proto_depIdxs,
		MessageInfos:      file_campusapis_staff_v1_freshman_proto_msgTypes,
	}.Build()
	File_campusapis_staff_v1_freshman_proto = out.File
	file_campusapis_staff_v1_freshman_proto_rawDesc = nil
	file_campusapis_staff_v1_freshman_proto_goTypes = nil
	file_campusapis_staff_v1_freshman_proto_depIdxs = nil
}
