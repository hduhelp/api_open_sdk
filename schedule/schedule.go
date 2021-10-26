package schedule

import (
	"encoding/json"
	"github.com/hduhelp/api_open_sdk/baseStaff"
	"github.com/hduhelp/api_open_sdk/schoolTime"
)

// AddMember 添加课程人员
func (m *ScheduleItem) AddMember(r CourseReader, t baseStaff.Type) {
	switch t {
	case baseStaff.Type_Teacher:
		//教师展示上课学生
		m.AddStudent(r)
	case baseStaff.Type_Undergraduate, baseStaff.Type_Postgraduate:
		//学生添加展示授课教师
		m.AddTeacher(r)
	}
}

func (m *ScheduleItem) AddTeacher(r CourseReader) {
	if m.Teachers == nil {
		m.Teachers = &baseStaff.InfoMapList{
			Items: map[string]*baseStaff.Info{},
		}
	}
	m.Teachers.Append(r.TeacherInfo())
}

func (m *ScheduleItem) AddStudent(r CourseReader) {
	if m.Students == nil {
		m.Students = &baseStaff.InfoMapList{
			Items: map[string]*baseStaff.Info{},
		}
	}
	m.Students.Append(r.StudentInfo())
}

type CourseReader interface {
	ScheduleID() string
	CourseInfo(schoolDate *schoolTime.SchoolDate) *ScheduleItem
	TeacherInfo() *baseStaff.Info
	StudentInfo() *baseStaff.Info
}

type Teachers struct {
	baseStaff.InfoMapList
}

type Students struct {
	baseStaff.InfoMapList
}

func (m Schedule) MarshalJSON() ([]byte, error) {
	list := make([]*ScheduleItem, 0)
	for _, v := range m.Items {
		list = append(list, v)
	}
	return json.Marshal(list)
}
