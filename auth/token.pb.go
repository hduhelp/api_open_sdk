// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: auth/token.proto

package auth

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

type TokenType int32

const (
	TokenType_Origin TokenType = 0
	TokenType_OAuth  TokenType = 1
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "Origin",
		1: "OAuth",
	}
	TokenType_value = map[string]int32{
		"Origin": 0,
		"OAuth":  1,
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
	return file_auth_token_proto_enumTypes[0].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_auth_token_proto_enumTypes[0]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_auth_token_proto_rawDescGZIP(), []int{0}
}

var File_auth_token_proto protoreflect.FileDescriptor

var file_auth_token_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x2a, 0x22, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x10, 0x01, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68, 0x65,
	0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_token_proto_rawDescOnce sync.Once
	file_auth_token_proto_rawDescData = file_auth_token_proto_rawDesc
)

func file_auth_token_proto_rawDescGZIP() []byte {
	file_auth_token_proto_rawDescOnce.Do(func() {
		file_auth_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_token_proto_rawDescData)
	})
	return file_auth_token_proto_rawDescData
}

var file_auth_token_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_token_proto_goTypes = []interface{}{
	(TokenType)(0), // 0: auth.TokenType
}
var file_auth_token_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_token_proto_init() }
func file_auth_token_proto_init() {
	if File_auth_token_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_token_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_token_proto_goTypes,
		DependencyIndexes: file_auth_token_proto_depIdxs,
		EnumInfos:         file_auth_token_proto_enumTypes,
	}.Build()
	File_auth_token_proto = out.File
	file_auth_token_proto_rawDesc = nil
	file_auth_token_proto_goTypes = nil
	file_auth_token_proto_depIdxs = nil
}
