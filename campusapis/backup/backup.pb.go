// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: campusapis/backup/backup.proto

package backup

import (
	v1 "github.com/hduhelp/api_open_sdk/campusapis/staff/v1"
	v11 "github.com/hduhelp/api_open_sdk/campusapis/teaching/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_campusapis_backup_backup_proto protoreflect.FileDescriptor

var file_campusapis_backup_backup_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x1a, 0x20, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x83, 0x06, 0x0a, 0x13, 0x43, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x82, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x63, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x15, 0x2f,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f,
	0x69, 0x6e, 0x66, 0x6f, 0x5a, 0x0e, 0x12, 0x0c, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x86, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x2b, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x29, 0x12, 0x16, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x5a, 0x0f, 0x12, 0x0d,
	0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x9f, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x12, 0x2b, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2b, 0x12, 0x17, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5a, 0x10, 0x12,
	0x0e, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12,
	0x9a, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x78,
	0x61, 0x6d, 0x12, 0x2a, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x45,
	0x78, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x29, 0x12, 0x16, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x5a, 0x0f, 0x12, 0x0d, 0x2f,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x12, 0x9e, 0x01, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x6f, 0x77, 0x56,
	0x32, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x30, 0x2e, 0x63, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x6f,
	0x77, 0x56, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x3a, 0x12, 0x1c, 0x2f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x6e, 0x6f,
	0x77, 0x5a, 0x1a, 0x12, 0x18, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x6e, 0x6f, 0x77, 0x42, 0xbc, 0x01,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x42, 0x0b, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x75, 0x68, 0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6f,
	0x70, 0x65, 0x6e, 0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0xa2, 0x02, 0x03, 0x43, 0x42, 0x58, 0xaa,
	0x02, 0x11, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0xca, 0x02, 0x11, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x61, 0x70, 0x69, 0x73,
	0x5c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0xe2, 0x02, 0x1d, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x61, 0x70, 0x69, 0x73, 0x5c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x61, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_campusapis_backup_backup_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),                // 0: google.protobuf.Empty
	(*v1.GetStudentGradeRequest)(nil),    // 1: campusapis.staff.v1.GetStudentGradeRequest
	(*v1.GetStudentExamRequest)(nil),     // 2: campusapis.staff.v1.GetStudentExamRequest
	(*v1.GetPersonInfoResponse)(nil),     // 3: campusapis.staff.v1.GetPersonInfoResponse
	(*v1.GetStudentInfoResponse)(nil),    // 4: campusapis.staff.v1.GetStudentInfoResponse
	(*v1.GetStudentGradeResponse)(nil),   // 5: campusapis.staff.v1.GetStudentGradeResponse
	(*v1.GetStudentExamResponse)(nil),    // 6: campusapis.staff.v1.GetStudentExamResponse
	(*v11.GetScheduleNowV2Response)(nil), // 7: campusapis.teaching.v1.GetScheduleNowV2Response
}
var file_campusapis_backup_backup_proto_depIdxs = []int32{
	0, // 0: campusapis.backup.CampusBackupService.GetPersonInfo:input_type -> google.protobuf.Empty
	0, // 1: campusapis.backup.CampusBackupService.GetStudentInfo:input_type -> google.protobuf.Empty
	1, // 2: campusapis.backup.CampusBackupService.GetStudentGrade:input_type -> campusapis.staff.v1.GetStudentGradeRequest
	2, // 3: campusapis.backup.CampusBackupService.GetStudentExam:input_type -> campusapis.staff.v1.GetStudentExamRequest
	0, // 4: campusapis.backup.CampusBackupService.GetScheduleNowV2:input_type -> google.protobuf.Empty
	3, // 5: campusapis.backup.CampusBackupService.GetPersonInfo:output_type -> campusapis.staff.v1.GetPersonInfoResponse
	4, // 6: campusapis.backup.CampusBackupService.GetStudentInfo:output_type -> campusapis.staff.v1.GetStudentInfoResponse
	5, // 7: campusapis.backup.CampusBackupService.GetStudentGrade:output_type -> campusapis.staff.v1.GetStudentGradeResponse
	6, // 8: campusapis.backup.CampusBackupService.GetStudentExam:output_type -> campusapis.staff.v1.GetStudentExamResponse
	7, // 9: campusapis.backup.CampusBackupService.GetScheduleNowV2:output_type -> campusapis.teaching.v1.GetScheduleNowV2Response
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_campusapis_backup_backup_proto_init() }
func file_campusapis_backup_backup_proto_init() {
	if File_campusapis_backup_backup_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_campusapis_backup_backup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_campusapis_backup_backup_proto_goTypes,
		DependencyIndexes: file_campusapis_backup_backup_proto_depIdxs,
	}.Build()
	File_campusapis_backup_backup_proto = out.File
	file_campusapis_backup_backup_proto_rawDesc = nil
	file_campusapis_backup_backup_proto_goTypes = nil
	file_campusapis_backup_backup_proto_depIdxs = nil
}
