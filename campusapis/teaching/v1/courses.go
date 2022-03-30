package teachingv1

import "github.com/hduhelp/api_open_sdk/baseStaff"

func (x *Courses) ToGetScheduleResponse() *GetScheduleResponse {
	return &GetScheduleResponse{
		Items: x.GetScheduleResponse_Course(),
	}
}

func (x *Courses) GetScheduleResponse_Course() (list []*GetScheduleResponse_Course) {
	list = make([]*GetScheduleResponse_Course, 0)
	for _, v := range x.Items {
		list = append(list, v.ToGetScheduleResponse_Course())
	}
	return
}

func (x *CourseItem) ToGetScheduleResponse_Course() *GetScheduleResponse_Course {
	return &GetScheduleResponse_Course{
		ClassID:    x.ClassID,
		ClassName:  x.ClassName,
		ClassTime:  x.ClassTime,
		CourseID:   x.CourseID,
		CourseName: x.CourseName,
		Credit:     x.Credit,
		Schedules:  x.GetScheduleResponse_Course_Schedule(),
		SchoolYear: x.SchoolYear.FullName(),
		Semester:   x.Semester.Num,
	}
}

func (x *CourseItem) GetScheduleResponse_Course_Schedule() (list []*GetScheduleResponse_Course_Schedule) {
	list = make([]*GetScheduleResponse_Course_Schedule, 0)
	if x.Schedule == nil {
		return list
	}
	for _, v := range x.Schedule.Items {
		list = append(list, v.ToGetScheduleResponse_Course_Schedule())
	}
	return list
}

func (x *ScheduleItem) ToGetScheduleResponse_Course_Schedule() *GetScheduleResponse_Course_Schedule {
	return &GetScheduleResponse_Course_Schedule{
		Location:   x.Location,
		SeatsNum:   x.SeatsNum,
		Section:    x.Section,
		Teachers:   convertStaffInfoFromInfoMapList(x.Teachers),
		Students:   convertStaffInfoFromInfoMapList(x.Students),
		Week:       x.Week,
		WeekDay:    x.WeekDay,
		IsThisWeek: x.IsThisWeek,
		StartTime:  x.StartTime,
		EndTime:    x.EndTime,
	}
}

func convertStaffInfoFromInfoMapList(m *baseStaff.InfoMapList) (list []*GetScheduleResponse_Course_StaffInfo) {
	list = make([]*GetScheduleResponse_Course_StaffInfo, 0)
	if len(m.Items) == 0 {
		return list
	}
	for _, v := range m.Items {
		list = append(list, &GetScheduleResponse_Course_StaffInfo{
			StaffID:   v.StaffID,
			StaffName: v.StaffName,
			StaffType: int32(v.StaffType),
		})
	}
	return list
}
