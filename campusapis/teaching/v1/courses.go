package teachingv1

import "github.com/hduhelp/api_open_sdk/campusapis/staff"

func (x *Courses) ToGetScheduleResponse() *GetScheduleResponse {
	return &GetScheduleResponse{
		Items: x.GetScheduleResponseCourse(),
	}
}

func (x *Courses) GetScheduleResponseCourse() (list []*ScheduleResponseCourse) {
	list = make([]*ScheduleResponseCourse, 0)
	for _, v := range x.Items {
		list = append(list, v.ToScheduleResponseCourse())
	}
	return list
}

func (i CourseItems) ToGetScheduleResponse() (list []*ScheduleResponseCourse) {
	list = make([]*ScheduleResponseCourse, 0)
	for _, v := range i {
		list = append(list, v.ToScheduleResponseCourse())
	}
	return list
}

func (x *CourseItem) ToScheduleResponseCourse() *ScheduleResponseCourse {
	return &ScheduleResponseCourse{
		ClassID:    x.ClassID,
		ClassName:  x.ClassName,
		ClassTime:  x.ClassTime,
		CourseID:   x.CourseID,
		CourseName: x.CourseName,
		Credit:     x.Credit,
		Schedules:  x.GetScheduleResponseCourse_Schedule(),
		SchoolYear: x.SchoolYear.FullName(),
		Semester:   x.Semester.Num,
	}
}

func (x *CourseItem) GetScheduleResponseCourse_Schedule() (list []*ScheduleResponseCourse_Schedule) {
	list = make([]*ScheduleResponseCourse_Schedule, 0)
	if x.Schedule == nil {
		return list
	}
	for _, v := range x.Schedule.Items {
		list = append(list, v.ToScheduleResponseCourse_Schedule())
	}
	return list
}

func (x *ScheduleItem) ToScheduleResponseCourse_Schedule() *ScheduleResponseCourse_Schedule {
	return &ScheduleResponseCourse_Schedule{
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

func convertStaffInfoFromInfoMapList(m *staff.InfoMapList) (list []*ScheduleResponseCourse_StaffInfo) {
	list = make([]*ScheduleResponseCourse_StaffInfo, 0)
	if len(m.Items) == 0 {
		return list
	}
	for _, v := range m.Items {
		list = append(list, &ScheduleResponseCourse_StaffInfo{
			StaffID:   v.StaffID,
			StaffName: v.StaffName,
			StaffType: int32(v.StaffType),
		})
	}
	return list
}
