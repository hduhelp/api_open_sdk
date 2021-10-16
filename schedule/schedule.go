package schedule

import (
	"encoding/json"
	"github.com/hduhelp/api_open_sdk/baseStaff"
	"time"
)

func (m *ScheduleItem) AddMember(r CourseReader, t baseStaff.Type) {
	switch t {
	case baseStaff.Type_Teacher:
		m.AddStudent(r)
	case baseStaff.Type_Undergraduate, baseStaff.Type_Postgraduate:
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
	CourseInfo(t time.Time) *ScheduleItem
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
