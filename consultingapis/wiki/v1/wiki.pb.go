// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.18.1
// source: consultingapis/wiki/v1/wiki.proto

package wikiv1

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

type DocTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DocTokenRequest) Reset() {
	*x = DocTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consultingapis_wiki_v1_wiki_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocTokenRequest) ProtoMessage() {}

func (x *DocTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_consultingapis_wiki_v1_wiki_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocTokenRequest.ProtoReflect.Descriptor instead.
func (*DocTokenRequest) Descriptor() ([]byte, []int) {
	return file_consultingapis_wiki_v1_wiki_proto_rawDescGZIP(), []int{0}
}

func (x *DocTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DocTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content  string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Revision uint32 `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *DocTokenResp) Reset() {
	*x = DocTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consultingapis_wiki_v1_wiki_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocTokenResp) ProtoMessage() {}

func (x *DocTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_consultingapis_wiki_v1_wiki_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocTokenResp.ProtoReflect.Descriptor instead.
func (*DocTokenResp) Descriptor() ([]byte, []int) {
	return file_consultingapis_wiki_v1_wiki_proto_rawDescGZIP(), []int{1}
}

func (x *DocTokenResp) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *DocTokenResp) GetRevision() uint32 {
	if x != nil {
		return x.Revision
	}
	return 0
}

var File_consultingapis_wiki_v1_wiki_proto protoreflect.FileDescriptor

var file_consultingapis_wiki_v1_wiki_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x77, 0x69, 0x6b, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x66, 0x61, 0x71, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0f, 0x44, 0x6f, 0x63, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x44, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x89, 0x01, 0x0a, 0x0a, 0x46, 0x41, 0x51, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x44, 0x6f, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x66, 0x61,
	0x71, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e,
	0x67, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x66, 0x61, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x12, 0x10, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74,
	0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consultingapis_wiki_v1_wiki_proto_rawDescOnce sync.Once
	file_consultingapis_wiki_v1_wiki_proto_rawDescData = file_consultingapis_wiki_v1_wiki_proto_rawDesc
)

func file_consultingapis_wiki_v1_wiki_proto_rawDescGZIP() []byte {
	file_consultingapis_wiki_v1_wiki_proto_rawDescOnce.Do(func() {
		file_consultingapis_wiki_v1_wiki_proto_rawDescData = protoimpl.X.CompressGZIP(file_consultingapis_wiki_v1_wiki_proto_rawDescData)
	})
	return file_consultingapis_wiki_v1_wiki_proto_rawDescData
}

var file_consultingapis_wiki_v1_wiki_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_consultingapis_wiki_v1_wiki_proto_goTypes = []interface{}{
	(*DocTokenRequest)(nil), // 0: consultingapis.faq.v1.DocTokenRequest
	(*DocTokenResp)(nil),    // 1: consultingapis.faq.v1.DocTokenResp
}
var file_consultingapis_wiki_v1_wiki_proto_depIdxs = []int32{
	0, // 0: consultingapis.faq.v1.FAQService.GetDriveDocContent:input_type -> consultingapis.faq.v1.DocTokenRequest
	1, // 1: consultingapis.faq.v1.FAQService.GetDriveDocContent:output_type -> consultingapis.faq.v1.DocTokenResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_consultingapis_wiki_v1_wiki_proto_init() }
func file_consultingapis_wiki_v1_wiki_proto_init() {
	if File_consultingapis_wiki_v1_wiki_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consultingapis_wiki_v1_wiki_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocTokenRequest); i {
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
		file_consultingapis_wiki_v1_wiki_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocTokenResp); i {
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
			RawDescriptor: file_consultingapis_wiki_v1_wiki_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_consultingapis_wiki_v1_wiki_proto_goTypes,
		DependencyIndexes: file_consultingapis_wiki_v1_wiki_proto_depIdxs,
		MessageInfos:      file_consultingapis_wiki_v1_wiki_proto_msgTypes,
	}.Build()
	File_consultingapis_wiki_v1_wiki_proto = out.File
	file_consultingapis_wiki_v1_wiki_proto_rawDesc = nil
	file_consultingapis_wiki_v1_wiki_proto_goTypes = nil
	file_consultingapis_wiki_v1_wiki_proto_depIdxs = nil
}
