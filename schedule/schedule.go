package schedule

import (
	"encoding/json"
	"github.com/hduhelp/api_open_sdk/baseStaff"
	"time"
)

func (m *Schedule) AddMember(r CourseReader, t baseStaff.Type) {
	switch t {
	case baseStaff.TypeTeacher:
		m.AddStudent(r)
	case baseStaff.TypeStudent:
		m.AddTeacher(r)
	}
}

func (m *Schedule) AddTeacher(r CourseReader) {
	if m.Teachers == nil {
		m.Teachers = &baseStaff.InfoMapList{
			Items: map[string]*baseStaff.Info{},
		}
	}
	m.Teachers.Append(r.TeacherInfo())
}

func (m *Schedule) AddStudent(r CourseReader) {
	if m.Students == nil {
		m.Students = &baseStaff.InfoMapList{
			Items: map[string]*baseStaff.Info{},
		}
	}
	m.Students.Append(r.StudentInfo())
}

type CourseReader interface {
	ScheduleID() string
	CourseInfo(t time.Time) *Schedule
	TeacherInfo() *baseStaff.Info
	StudentInfo() *baseStaff.Info
}

type Teachers struct {
	baseStaff.InfoMapList
}

type Students struct {
	baseStaff.InfoMapList
}

type Courses map[string]*Schedule

func (c Courses) MarshalJSON() ([]byte, error) {
	list := make([]*Schedule, 0)
	for _, v := range c {
		list = append(list, v)
	}
	return json.Marshal(list)
}
