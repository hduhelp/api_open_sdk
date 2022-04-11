// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: gatewayapis/auth/v1/auth.proto

package authv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenType int32

const (
	TokenType_ORIGIN TokenType = 0
	TokenType_OAUTH  TokenType = 1
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "ORIGIN",
		1: "OAUTH",
	}
	TokenType_value = map[string]int32{
		"ORIGIN": 0,
		"OAUTH":  1,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_gatewayapis_auth_v1_auth_proto_enumTypes[0].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_gatewayapis_auth_v1_auth_proto_enumTypes[0]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_gatewayapis_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

type GetTokenInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid     bool      `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
	AccessToken string    `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	ExpiredTime uint32    `protobuf:"varint,3,opt,name=expiredTime,proto3" json:"expiredTime,omitempty"`
	GrantType   string    `protobuf:"bytes,4,opt,name=grantType,proto3" json:"grantType,omitempty"`
	ClientID    string    `protobuf:"bytes,5,opt,name=clientID,proto3" json:"clientID,omitempty"`
	TokenType   TokenType `protobuf:"varint,6,opt,name=tokenType,proto3,enum=gatewayapis.auth.v1.TokenType" json:"tokenType,omitempty"`
}

func (x *GetTokenInfoResponse) Reset() {
	*x = GetTokenInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_auth_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenInfoResponse) ProtoMessage() {}

func (x *GetTokenInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_auth_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenInfoResponse.ProtoReflect.Descriptor instead.
func (*GetTokenInfoResponse) Descriptor() ([]byte, []int) {
	return file_gatewayapis_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *GetTokenInfoResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *GetTokenInfoResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GetTokenInfoResponse) GetExpiredTime() uint32 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

func (x *GetTokenInfoResponse) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

func (x *GetTokenInfoResponse) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *GetTokenInfoResponse) GetTokenType() TokenType {
	if x != nil {
		return x.TokenType
	}
	return TokenType_ORIGIN
}

type UserIdListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 助手user_id列表
	UserIds []string `protobuf:"bytes,1,rep,name=userIds,proto3" json:"userIds,omitempty"`
}

func (x *UserIdListRequest) Reset() {
	*x = UserIdListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_auth_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdListRequest) ProtoMessage() {}

func (x *UserIdListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_auth_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdListRequest.ProtoReflect.Descriptor instead.
func (*UserIdListRequest) Descriptor() ([]byte, []int) {
	return file_gatewayapis_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserIdListRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// 用户绑定信息查询响应
type BindListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 绑定信息列表
	Items []*Bind `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *BindListResponse) Reset() {
	*x = BindListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_auth_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindListResponse) ProtoMessage() {}

func (x *BindListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_auth_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindListResponse.ProtoReflect.Descriptor instead.
func (*BindListResponse) Descriptor() ([]byte, []int) {
	return file_gatewayapis_auth_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *BindListResponse) GetItems() []*Bind {
	if x != nil {
		return x.Items
	}
	return nil
}

// 用户绑定信息
type Bind struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 助手用户user_id
	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	// 绑定平台：
	// platform=hdu 杭电学（工）号，uid 为学工号，同一用户可能存在多个学号（二学位，研究生）；
	// platform=qzu 衢州学院学（工）号，uid 为学工号
	// platform=wxmp 微信公众号，uid 为微信open id，uid2 为微信union id；
	// platform=dingtalk 钉钉，uid 为钉钉 user id，同一用户可能存在多个钉钉绑定记录，优先用 uid2 存在的记录条的 uid；
	// platform=idCard 身份证，uid 为身份证号；
	// platform=alipay 支付宝, uid todo
	// platform=apple 苹果，uid todo, uid2 todo
	// platform=github github, uid todo, uid2 todo
	// platform=qq QQ号，uid todo, uid2 todo
	// platform=yiban 易班，uid todo
	Platform string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	// uid 不同平台含义不同，见 platform 解释
	Uid string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	// uid 不同平台含义不同，见 platform 解释
	Uid2 string `protobuf:"bytes,4,opt,name=uid2,proto3" json:"uid2,omitempty"`
	// 是否为主学工号
	Primary bool `protobuf:"varint,5,opt,name=primary,proto3" json:"primary,omitempty"`
	// 备注
	Remark string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	// 额外信息，json 格式
	Attributes *structpb.Struct `protobuf:"bytes,7,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Bind) Reset() {
	*x = Bind{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewayapis_auth_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bind) ProtoMessage() {}

func (x *Bind) ProtoReflect() protoreflect.Message {
	mi := &file_gatewayapis_auth_v1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bind.ProtoReflect.Descriptor instead.
func (*Bind) Descriptor() ([]byte, []int) {
	return file_gatewayapis_auth_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *Bind) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Bind) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Bind) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Bind) GetUid2() string {
	if x != nil {
		return x.Uid2
	}
	return ""
}

func (x *Bind) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *Bind) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Bind) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

var File_gatewayapis_auth_v1_auth_proto protoreflect.FileDescriptor

var file_gatewayapis_auth_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xec, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x3c, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x2d, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x43,
	0x0a, 0x10, 0x42, 0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0xcb, 0x01, 0x0a, 0x04, 0x42, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x69, 0x64, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x69, 0x64, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x2a, 0x22, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x41,
	0x55, 0x54, 0x48, 0x10, 0x01, 0x32, 0xce, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x29, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x42, 0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xce, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x42, 0x09, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68,
	0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x58, 0xaa, 0x02, 0x13, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x41, 0x75, 0x74, 0x68,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x15, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x41,
	0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gatewayapis_auth_v1_auth_proto_rawDescOnce sync.Once
	file_gatewayapis_auth_v1_auth_proto_rawDescData = file_gatewayapis_auth_v1_auth_proto_rawDesc
)

func file_gatewayapis_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_gatewayapis_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_gatewayapis_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_gatewayapis_auth_v1_auth_proto_rawDescData)
	})
	return file_gatewayapis_auth_v1_auth_proto_rawDescData
}

var file_gatewayapis_auth_v1_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gatewayapis_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gatewayapis_auth_v1_auth_proto_goTypes = []interface{}{
	(TokenType)(0),               // 0: gatewayapis.auth.v1.TokenType
	(*GetTokenInfoResponse)(nil), // 1: gatewayapis.auth.v1.GetTokenInfoResponse
	(*UserIdListRequest)(nil),    // 2: gatewayapis.auth.v1.UserIdListRequest
	(*BindListResponse)(nil),     // 3: gatewayapis.auth.v1.BindListResponse
	(*Bind)(nil),                 // 4: gatewayapis.auth.v1.Bind
	(*structpb.Struct)(nil),      // 5: google.protobuf.Struct
	(*emptypb.Empty)(nil),        // 6: google.protobuf.Empty
}
var file_gatewayapis_auth_v1_auth_proto_depIdxs = []int32{
	0, // 0: gatewayapis.auth.v1.GetTokenInfoResponse.tokenType:type_name -> gatewayapis.auth.v1.TokenType
	4, // 1: gatewayapis.auth.v1.BindListResponse.items:type_name -> gatewayapis.auth.v1.Bind
	5, // 2: gatewayapis.auth.v1.Bind.attributes:type_name -> google.protobuf.Struct
	6, // 3: gatewayapis.auth.v1.AuthService.GetTokenInfo:input_type -> google.protobuf.Empty
	2, // 4: gatewayapis.auth.v1.AuthService.GetBindListByUserIdList:input_type -> gatewayapis.auth.v1.UserIdListRequest
	1, // 5: gatewayapis.auth.v1.AuthService.GetTokenInfo:output_type -> gatewayapis.auth.v1.GetTokenInfoResponse
	3, // 6: gatewayapis.auth.v1.AuthService.GetBindListByUserIdList:output_type -> gatewayapis.auth.v1.BindListResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gatewayapis_auth_v1_auth_proto_init() }
func file_gatewayapis_auth_v1_auth_proto_init() {
	if File_gatewayapis_auth_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gatewayapis_auth_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenInfoResponse); i {
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
		file_gatewayapis_auth_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdListRequest); i {
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
		file_gatewayapis_auth_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindListResponse); i {
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
		file_gatewayapis_auth_v1_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bind); i {
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
			RawDescriptor: file_gatewayapis_auth_v1_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gatewayapis_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_gatewayapis_auth_v1_auth_proto_depIdxs,
		EnumInfos:         file_gatewayapis_auth_v1_auth_proto_enumTypes,
		MessageInfos:      file_gatewayapis_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_gatewayapis_auth_v1_auth_proto = out.File
	file_gatewayapis_auth_v1_auth_proto_rawDesc = nil
	file_gatewayapis_auth_v1_auth_proto_goTypes = nil
	file_gatewayapis_auth_v1_auth_proto_depIdxs = nil
}
