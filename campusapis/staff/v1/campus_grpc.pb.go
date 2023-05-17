// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: campusapis/staff/v1/campus.proto

package staffv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CampusService_GetPersonInfo_FullMethodName              = "/campusapis.staff.v1.CampusService/GetPersonInfo"
	CampusService_GetStudentInfo_FullMethodName             = "/campusapis.staff.v1.CampusService/GetStudentInfo"
	CampusService_GetStudentSchoolRollStatus_FullMethodName = "/campusapis.staff.v1.CampusService/GetStudentSchoolRollStatus"
	CampusService_GetStudentDormInfo_FullMethodName         = "/campusapis.staff.v1.CampusService/GetStudentDormInfo"
	CampusService_GetStudentBirthdayInfo_FullMethodName     = "/campusapis.staff.v1.CampusService/GetStudentBirthdayInfo"
	CampusService_GetStudentBirthdaysIn_FullMethodName      = "/campusapis.staff.v1.CampusService/GetStudentBirthdaysIn"
	CampusService_GetStudentRewards_FullMethodName          = "/campusapis.staff.v1.CampusService/GetStudentRewards"
	CampusService_GetStudentCourseSelections_FullMethodName = "/campusapis.staff.v1.CampusService/GetStudentCourseSelections"
	CampusService_GetStudentGrade_FullMethodName            = "/campusapis.staff.v1.CampusService/GetStudentGrade"
	CampusService_GetGlobalStudentGrade_FullMethodName      = "/campusapis.staff.v1.CampusService/GetGlobalStudentGrade"
	CampusService_GetStudentExam_FullMethodName             = "/campusapis.staff.v1.CampusService/GetStudentExam"
	CampusService_GetStudentGPA_FullMethodName              = "/campusapis.staff.v1.CampusService/GetStudentGPA"
	CampusService_PostStudentGateAccess_FullMethodName      = "/campusapis.staff.v1.CampusService/PostStudentGateAccess"
	CampusService_GetStudentStaySchoolInfo_FullMethodName   = "/campusapis.staff.v1.CampusService/GetStudentStaySchoolInfo"
	CampusService_GetFreshmanBaseInfo_FullMethodName        = "/campusapis.staff.v1.CampusService/GetFreshmanBaseInfo"
	CampusService_GetFreshmanDetail_FullMethodName          = "/campusapis.staff.v1.CampusService/GetFreshmanDetail"
	CampusService_GetFreshmanRoommates_FullMethodName       = "/campusapis.staff.v1.CampusService/GetFreshmanRoommates"
)

// CampusServiceClient is the client API for CampusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampusServiceClient interface {
	// 获取员工信息
	GetPersonInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPersonInfoResponse, error)
	// 获取学生基本信息
	GetStudentInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentInfoResponse, error)
	// 获取学生学籍状态
	GetStudentSchoolRollStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentSchoolRollStatusResponse, error)
	//	迁移至校外部分
	//	//获取学生贫困生申请记录
	//	rpc GetStudentNeedyInfo(google.protobuf.Empty) returns (GetStudentNeedyInfoResponse) {
	//	  option (google.api.http) = {
	//	    get: "/staff/v1/student/needy"
	//	    additional_bindings {
	//	      get: "/student/needy"
	//	    }
	//	  };
	//	}
	//
	// 获取学生宿舍信息
	GetStudentDormInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentDormInfoResponse, error)
	// 获取学生生日信息
	GetStudentBirthdayInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentBirthdayInfoResponse, error)
	// 获取生日为指定日期的学生列表
	GetStudentBirthdaysIn(ctx context.Context, in *GetStudentBirthdaysInRequest, opts ...grpc.CallOption) (*GetStudentBirthdaysInResponse, error)
	// 获取学生奖学金信息
	GetStudentRewards(ctx context.Context, in *GetStudentRewardsRequest, opts ...grpc.CallOption) (*GetStudentRewardsResponse, error)
	// 获取学生选课信息
	GetStudentCourseSelections(ctx context.Context, in *GetStudentCourseSelectionsRequest, opts ...grpc.CallOption) (*GetStudentCourseSelectionsResponse, error)
	// 获取学生成绩信息
	GetStudentGrade(ctx context.Context, in *GetStudentGradeRequest, opts ...grpc.CallOption) (*GetStudentGradeResponse, error)
	// 获取全局成绩，用于推送成绩信息，仅内部使用
	GetGlobalStudentGrade(ctx context.Context, in *GetGlobalStudentGradeRequest, opts ...grpc.CallOption) (*GetGlobalStudentGradeResponse, error)
	// 获取学生考试信息
	GetStudentExam(ctx context.Context, in *GetStudentExamRequest, opts ...grpc.CallOption) (*GetStudentExamResponse, error)
	// 获取学生GPA信息
	GetStudentGPA(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentGPAResponse, error)
	// 推送学生门禁通行信息
	PostStudentGateAccess(ctx context.Context, in *PostStudentGateAccessRequest, opts ...grpc.CallOption) (*PostStudentGateAccessResponse, error)
	// 获取学生留校信息
	GetStudentStaySchoolInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentStaySchoolInfoResponse, error)
	// 获取新生基本信息
	GetFreshmanBaseInfo(ctx context.Context, in *GetFreshmanBaseInfoRequest, opts ...grpc.CallOption) (*GetFreshmanBaseInfoResponse, error)
	// 获取新生详细信息
	GetFreshmanDetail(ctx context.Context, in *GetFreshmanDetailRequest, opts ...grpc.CallOption) (*GetFreshmanDetailResponse, error)
	// 获取新生宿舍信息
	GetFreshmanRoommates(ctx context.Context, in *GetFreshmanRoommatesRequest, opts ...grpc.CallOption) (*GetFreshmanRoommatesResponse, error)
}

type campusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampusServiceClient(cc grpc.ClientConnInterface) CampusServiceClient {
	return &campusServiceClient{cc}
}

func (c *campusServiceClient) GetPersonInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPersonInfoResponse, error) {
	out := new(GetPersonInfoResponse)
	err := c.cc.Invoke(ctx, CampusService_GetPersonInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetStudentInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentInfoResponse, error) {
	out := new(GetStudentInfoResponse)
	err := c.cc.Invoke(ctx, CampusService_GetStudentInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetStudentSchoolRollStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentSchoolRollStatusResponse, error) {
	out := new(GetStudentSchoolRollStatusResponse)
	err := c.cc.Invoke(ctx, CampusService_GetStudentSchoolRollStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetStudentDormInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentDormInfoResponse, error) {
	out := new(GetStudentDormInfoResponse)
	err := c.cc.Invoke(ctx, CampusService_GetStudentDormInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetStudentBirthdayInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentBirthdayInfoResponse, error) {
	out := new(GetStudentBirthdayInfoResponse)
	err := c.cc.Invoke(ctx, CampusService_GetStudentBirthdayInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetStudentBirthdaysIn(ctx context.Context, in *GetStudentBirthdaysInRequest, opts ...grpc.CallOption) (*GetStudentBirthdaysInResponse, error) {
	out := new(GetStudentBirthdaysInResponse)
	err := c.cc.Invoke(ctx, CampusService_GetStudentBirthdaysIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetStudentRewards(ctx context.Context, in *GetStudentRewardsRequest, opts ...grpc.CallOption) (*GetStudentRewardsResponse, error) {
	out := new(GetStudentRewardsResponse)
	err := c.cc.Invoke(ctx, CampusService_GetStudentRewards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetStudentCourseSelections(ctx context.Context, in *GetStudentCourseSelectionsRequest, opts ...grpc.CallOption) (*GetStudentCourseSelectionsResponse, error) {
	out := new(GetStudentCourseSelectionsResponse)
	err := c.cc.Invoke(ctx, CampusService_GetStudentCourseSelections_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetStudentGrade(ctx context.Context, in *GetStudentGradeRequest, opts ...grpc.CallOption) (*GetStudentGradeResponse, error) {
	out := new(GetStudentGradeResponse)
	err := c.cc.Invoke(ctx, CampusService_GetStudentGrade_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetGlobalStudentGrade(ctx context.Context, in *GetGlobalStudentGradeRequest, opts ...grpc.CallOption) (*GetGlobalStudentGradeResponse, error) {
	out := new(GetGlobalStudentGradeResponse)
	err := c.cc.Invoke(ctx, CampusService_GetGlobalStudentGrade_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetStudentExam(ctx context.Context, in *GetStudentExamRequest, opts ...grpc.CallOption) (*GetStudentExamResponse, error) {
	out := new(GetStudentExamResponse)
	err := c.cc.Invoke(ctx, CampusService_GetStudentExam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetStudentGPA(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentGPAResponse, error) {
	out := new(GetStudentGPAResponse)
	err := c.cc.Invoke(ctx, CampusService_GetStudentGPA_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) PostStudentGateAccess(ctx context.Context, in *PostStudentGateAccessRequest, opts ...grpc.CallOption) (*PostStudentGateAccessResponse, error) {
	out := new(PostStudentGateAccessResponse)
	err := c.cc.Invoke(ctx, CampusService_PostStudentGateAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetStudentStaySchoolInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentStaySchoolInfoResponse, error) {
	out := new(GetStudentStaySchoolInfoResponse)
	err := c.cc.Invoke(ctx, CampusService_GetStudentStaySchoolInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetFreshmanBaseInfo(ctx context.Context, in *GetFreshmanBaseInfoRequest, opts ...grpc.CallOption) (*GetFreshmanBaseInfoResponse, error) {
	out := new(GetFreshmanBaseInfoResponse)
	err := c.cc.Invoke(ctx, CampusService_GetFreshmanBaseInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetFreshmanDetail(ctx context.Context, in *GetFreshmanDetailRequest, opts ...grpc.CallOption) (*GetFreshmanDetailResponse, error) {
	out := new(GetFreshmanDetailResponse)
	err := c.cc.Invoke(ctx, CampusService_GetFreshmanDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusServiceClient) GetFreshmanRoommates(ctx context.Context, in *GetFreshmanRoommatesRequest, opts ...grpc.CallOption) (*GetFreshmanRoommatesResponse, error) {
	out := new(GetFreshmanRoommatesResponse)
	err := c.cc.Invoke(ctx, CampusService_GetFreshmanRoommates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampusServiceServer is the server API for CampusService service.
// All implementations must embed UnimplementedCampusServiceServer
// for forward compatibility
type CampusServiceServer interface {
	// 获取员工信息
	GetPersonInfo(context.Context, *emptypb.Empty) (*GetPersonInfoResponse, error)
	// 获取学生基本信息
	GetStudentInfo(context.Context, *emptypb.Empty) (*GetStudentInfoResponse, error)
	// 获取学生学籍状态
	GetStudentSchoolRollStatus(context.Context, *emptypb.Empty) (*GetStudentSchoolRollStatusResponse, error)
	//	迁移至校外部分
	//	//获取学生贫困生申请记录
	//	rpc GetStudentNeedyInfo(google.protobuf.Empty) returns (GetStudentNeedyInfoResponse) {
	//	  option (google.api.http) = {
	//	    get: "/staff/v1/student/needy"
	//	    additional_bindings {
	//	      get: "/student/needy"
	//	    }
	//	  };
	//	}
	//
	// 获取学生宿舍信息
	GetStudentDormInfo(context.Context, *emptypb.Empty) (*GetStudentDormInfoResponse, error)
	// 获取学生生日信息
	GetStudentBirthdayInfo(context.Context, *emptypb.Empty) (*GetStudentBirthdayInfoResponse, error)
	// 获取生日为指定日期的学生列表
	GetStudentBirthdaysIn(context.Context, *GetStudentBirthdaysInRequest) (*GetStudentBirthdaysInResponse, error)
	// 获取学生奖学金信息
	GetStudentRewards(context.Context, *GetStudentRewardsRequest) (*GetStudentRewardsResponse, error)
	// 获取学生选课信息
	GetStudentCourseSelections(context.Context, *GetStudentCourseSelectionsRequest) (*GetStudentCourseSelectionsResponse, error)
	// 获取学生成绩信息
	GetStudentGrade(context.Context, *GetStudentGradeRequest) (*GetStudentGradeResponse, error)
	// 获取全局成绩，用于推送成绩信息，仅内部使用
	GetGlobalStudentGrade(context.Context, *GetGlobalStudentGradeRequest) (*GetGlobalStudentGradeResponse, error)
	// 获取学生考试信息
	GetStudentExam(context.Context, *GetStudentExamRequest) (*GetStudentExamResponse, error)
	// 获取学生GPA信息
	GetStudentGPA(context.Context, *emptypb.Empty) (*GetStudentGPAResponse, error)
	// 推送学生门禁通行信息
	PostStudentGateAccess(context.Context, *PostStudentGateAccessRequest) (*PostStudentGateAccessResponse, error)
	// 获取学生留校信息
	GetStudentStaySchoolInfo(context.Context, *emptypb.Empty) (*GetStudentStaySchoolInfoResponse, error)
	// 获取新生基本信息
	GetFreshmanBaseInfo(context.Context, *GetFreshmanBaseInfoRequest) (*GetFreshmanBaseInfoResponse, error)
	// 获取新生详细信息
	GetFreshmanDetail(context.Context, *GetFreshmanDetailRequest) (*GetFreshmanDetailResponse, error)
	// 获取新生宿舍信息
	GetFreshmanRoommates(context.Context, *GetFreshmanRoommatesRequest) (*GetFreshmanRoommatesResponse, error)
	mustEmbedUnimplementedCampusServiceServer()
}

// UnimplementedCampusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampusServiceServer struct {
}

func (UnimplementedCampusServiceServer) GetPersonInfo(context.Context, *emptypb.Empty) (*GetPersonInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonInfo not implemented")
}
func (UnimplementedCampusServiceServer) GetStudentInfo(context.Context, *emptypb.Empty) (*GetStudentInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentInfo not implemented")
}
func (UnimplementedCampusServiceServer) GetStudentSchoolRollStatus(context.Context, *emptypb.Empty) (*GetStudentSchoolRollStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentSchoolRollStatus not implemented")
}
func (UnimplementedCampusServiceServer) GetStudentDormInfo(context.Context, *emptypb.Empty) (*GetStudentDormInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentDormInfo not implemented")
}
func (UnimplementedCampusServiceServer) GetStudentBirthdayInfo(context.Context, *emptypb.Empty) (*GetStudentBirthdayInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentBirthdayInfo not implemented")
}
func (UnimplementedCampusServiceServer) GetStudentBirthdaysIn(context.Context, *GetStudentBirthdaysInRequest) (*GetStudentBirthdaysInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentBirthdaysIn not implemented")
}
func (UnimplementedCampusServiceServer) GetStudentRewards(context.Context, *GetStudentRewardsRequest) (*GetStudentRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentRewards not implemented")
}
func (UnimplementedCampusServiceServer) GetStudentCourseSelections(context.Context, *GetStudentCourseSelectionsRequest) (*GetStudentCourseSelectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentCourseSelections not implemented")
}
func (UnimplementedCampusServiceServer) GetStudentGrade(context.Context, *GetStudentGradeRequest) (*GetStudentGradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentGrade not implemented")
}
func (UnimplementedCampusServiceServer) GetGlobalStudentGrade(context.Context, *GetGlobalStudentGradeRequest) (*GetGlobalStudentGradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobalStudentGrade not implemented")
}
func (UnimplementedCampusServiceServer) GetStudentExam(context.Context, *GetStudentExamRequest) (*GetStudentExamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentExam not implemented")
}
func (UnimplementedCampusServiceServer) GetStudentGPA(context.Context, *emptypb.Empty) (*GetStudentGPAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentGPA not implemented")
}
func (UnimplementedCampusServiceServer) PostStudentGateAccess(context.Context, *PostStudentGateAccessRequest) (*PostStudentGateAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostStudentGateAccess not implemented")
}
func (UnimplementedCampusServiceServer) GetStudentStaySchoolInfo(context.Context, *emptypb.Empty) (*GetStudentStaySchoolInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentStaySchoolInfo not implemented")
}
func (UnimplementedCampusServiceServer) GetFreshmanBaseInfo(context.Context, *GetFreshmanBaseInfoRequest) (*GetFreshmanBaseInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreshmanBaseInfo not implemented")
}
func (UnimplementedCampusServiceServer) GetFreshmanDetail(context.Context, *GetFreshmanDetailRequest) (*GetFreshmanDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreshmanDetail not implemented")
}
func (UnimplementedCampusServiceServer) GetFreshmanRoommates(context.Context, *GetFreshmanRoommatesRequest) (*GetFreshmanRoommatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreshmanRoommates not implemented")
}
func (UnimplementedCampusServiceServer) mustEmbedUnimplementedCampusServiceServer() {}

// UnsafeCampusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampusServiceServer will
// result in compilation errors.
type UnsafeCampusServiceServer interface {
	mustEmbedUnimplementedCampusServiceServer()
}

func RegisterCampusServiceServer(s grpc.ServiceRegistrar, srv CampusServiceServer) {
	s.RegisterService(&CampusService_ServiceDesc, srv)
}

func _CampusService_GetPersonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetPersonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetPersonInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetPersonInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetStudentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetStudentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetStudentInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetStudentInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetStudentSchoolRollStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetStudentSchoolRollStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetStudentSchoolRollStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetStudentSchoolRollStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetStudentDormInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetStudentDormInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetStudentDormInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetStudentDormInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetStudentBirthdayInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetStudentBirthdayInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetStudentBirthdayInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetStudentBirthdayInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetStudentBirthdaysIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentBirthdaysInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetStudentBirthdaysIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetStudentBirthdaysIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetStudentBirthdaysIn(ctx, req.(*GetStudentBirthdaysInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetStudentRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetStudentRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetStudentRewards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetStudentRewards(ctx, req.(*GetStudentRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetStudentCourseSelections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentCourseSelectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetStudentCourseSelections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetStudentCourseSelections_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetStudentCourseSelections(ctx, req.(*GetStudentCourseSelectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetStudentGrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentGradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetStudentGrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetStudentGrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetStudentGrade(ctx, req.(*GetStudentGradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetGlobalStudentGrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGlobalStudentGradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetGlobalStudentGrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetGlobalStudentGrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetGlobalStudentGrade(ctx, req.(*GetGlobalStudentGradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetStudentExam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentExamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetStudentExam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetStudentExam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetStudentExam(ctx, req.(*GetStudentExamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetStudentGPA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetStudentGPA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetStudentGPA_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetStudentGPA(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_PostStudentGateAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostStudentGateAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).PostStudentGateAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_PostStudentGateAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).PostStudentGateAccess(ctx, req.(*PostStudentGateAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetStudentStaySchoolInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetStudentStaySchoolInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetStudentStaySchoolInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetStudentStaySchoolInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetFreshmanBaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFreshmanBaseInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetFreshmanBaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetFreshmanBaseInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetFreshmanBaseInfo(ctx, req.(*GetFreshmanBaseInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetFreshmanDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFreshmanDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetFreshmanDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetFreshmanDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetFreshmanDetail(ctx, req.(*GetFreshmanDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampusService_GetFreshmanRoommates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFreshmanRoommatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServiceServer).GetFreshmanRoommates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampusService_GetFreshmanRoommates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServiceServer).GetFreshmanRoommates(ctx, req.(*GetFreshmanRoommatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampusService_ServiceDesc is the grpc.ServiceDesc for CampusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "campusapis.staff.v1.CampusService",
	HandlerType: (*CampusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPersonInfo",
			Handler:    _CampusService_GetPersonInfo_Handler,
		},
		{
			MethodName: "GetStudentInfo",
			Handler:    _CampusService_GetStudentInfo_Handler,
		},
		{
			MethodName: "GetStudentSchoolRollStatus",
			Handler:    _CampusService_GetStudentSchoolRollStatus_Handler,
		},
		{
			MethodName: "GetStudentDormInfo",
			Handler:    _CampusService_GetStudentDormInfo_Handler,
		},
		{
			MethodName: "GetStudentBirthdayInfo",
			Handler:    _CampusService_GetStudentBirthdayInfo_Handler,
		},
		{
			MethodName: "GetStudentBirthdaysIn",
			Handler:    _CampusService_GetStudentBirthdaysIn_Handler,
		},
		{
			MethodName: "GetStudentRewards",
			Handler:    _CampusService_GetStudentRewards_Handler,
		},
		{
			MethodName: "GetStudentCourseSelections",
			Handler:    _CampusService_GetStudentCourseSelections_Handler,
		},
		{
			MethodName: "GetStudentGrade",
			Handler:    _CampusService_GetStudentGrade_Handler,
		},
		{
			MethodName: "GetGlobalStudentGrade",
			Handler:    _CampusService_GetGlobalStudentGrade_Handler,
		},
		{
			MethodName: "GetStudentExam",
			Handler:    _CampusService_GetStudentExam_Handler,
		},
		{
			MethodName: "GetStudentGPA",
			Handler:    _CampusService_GetStudentGPA_Handler,
		},
		{
			MethodName: "PostStudentGateAccess",
			Handler:    _CampusService_PostStudentGateAccess_Handler,
		},
		{
			MethodName: "GetStudentStaySchoolInfo",
			Handler:    _CampusService_GetStudentStaySchoolInfo_Handler,
		},
		{
			MethodName: "GetFreshmanBaseInfo",
			Handler:    _CampusService_GetFreshmanBaseInfo_Handler,
		},
		{
			MethodName: "GetFreshmanDetail",
			Handler:    _CampusService_GetFreshmanDetail_Handler,
		},
		{
			MethodName: "GetFreshmanRoommates",
			Handler:    _CampusService_GetFreshmanRoommates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "campusapis/staff/v1/campus.proto",
}
