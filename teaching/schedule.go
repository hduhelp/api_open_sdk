package teaching

import (
	"encoding/json"
	"github.com/hduhelp/api_open_sdk/baseStaff"
	"github.com/hduhelp/api_open_sdk/schoolTime"
)

// AddSchedule 添加课表
func (x *CourseItem) AddSchedule(r ScheduleReader, q *CourseQuery) {
	if x.Schedule == nil {
		x.Schedule = &Schedule{Items: map[string]*ScheduleItem{}}
	}
	if x.Schedule.Items[r.ScheduleID()] == nil {
		x.Schedule.Items[r.ScheduleID()] = r.ScheduleInfo(q.SchoolDate)
	} else {
		x.Schedule.Items[r.ScheduleID()].AddMember(r, q.QueryStaff.Type)
	}
}

// AddMember 添加课程人员
func (x *ScheduleItem) AddMember(r ScheduleReader, t baseStaff.Type) {
	switch t {
	case baseStaff.Type_Teacher:
		//教师展示上课学生
		x.AddStudent(r)
	case baseStaff.Type_Undergraduate, baseStaff.Type_Postgraduate:
		//学生添加展示授课教师
		x.AddTeacher(r)
	}
}

func (x *ScheduleItem) AddTeacher(r ScheduleReader) {
	if x.Teachers == nil {
		x.Teachers = &baseStaff.InfoMapList{
			Items: map[string]*baseStaff.Info{},
		}
	}
	x.Teachers.Append(r.TeacherInfo())
}

func (x *ScheduleItem) AddStudent(r ScheduleReader) {
	if x.Students == nil {
		x.Students = &baseStaff.InfoMapList{
			Items: map[string]*baseStaff.Info{},
		}
	}
	x.Students.Append(r.StudentInfo())
}

type CourseReader interface {
	CourseID() string
	CourseInfo() *CourseItem
	ScheduleInfo() ScheduleReader
}

type ScheduleReader interface {
	ScheduleID() string
	ScheduleInfo(schoolDate *schoolTime.SchoolDate) *ScheduleItem
	TeacherInfo() *baseStaff.Info
	StudentInfo() *baseStaff.Info
}

type Teachers struct {
	baseStaff.InfoMapList
}

type Students struct {
	baseStaff.InfoMapList
}

func (x *Course) MarshalJSON() ([]byte, error) {
	list := make([]*CourseItem, 0)
	for _, v := range x.Items {
		list = append(list, v)
	}
	return json.Marshal(list)
}

func (x *Schedule) MarshalJSON() ([]byte, error) {
	list := make([]*ScheduleItem, 0)
	for _, v := range x.Items {
		list = append(list, v)
	}
	return json.Marshal(list)
}
