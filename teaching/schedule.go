package teaching

import (
	"encoding/json"
	"fmt"
	"github.com/hduhelp/api_open_sdk/baseStaff"
	"sort"
)

// AddSchedule 添加课表
func (x *CourseItem) AddSchedule(q *CourseQuery, r ScheduleReader) {
	if x.Schedule == nil {
		x.Schedule = &Schedule{Items: map[string]*ScheduleItem{}}
	}
	if x.Schedule.Items[r.ScheduleID()] == nil {
		x.Schedule.Items[r.ScheduleID()] = r.ScheduleInfo()
		x.Schedule.Items[r.ScheduleID()].SetTime(q)
	}

	x.Schedule.Items[r.ScheduleID()].AddMember(q, r)
}

func (x *ScheduleItem) SetTime(q *CourseQuery) {
	//非当前周课程 跳过时间赋值
	if _, ok := InArray(x.Week, q.SchoolDate.Week); !ok {
		return
	}
	d, err := q.SchoolDateToDater.GetSchoolDateFrom(q.SchoolDate, x.WeekDay)
	if err != nil {
		return
	}
	x.StartTime = q.SectionReader.StartTime(d, FirstOfArray(x.Section)).Unix()
	x.EndTime = q.SectionReader.EndTime(d, LastOfArray(x.Section)).Unix()
	x.IsThisWeek = true
}

// AddMember 添加课程人员
func (x *ScheduleItem) AddMember(q *CourseQuery, r ScheduleReader) {
	switch q.QueryStaff.Type {
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
	// CourseID 课程ID
	CourseID() string
	CourseInfo() *CourseItem
	ScheduleReader() ScheduleReader
}

type ScheduleReader interface {
	ScheduleID() string
	ScheduleInfo() *ScheduleItem
	TeacherInfo() *baseStaff.Info
	StudentInfo() *baseStaff.Info
}

type Teachers struct {
	baseStaff.InfoMapList
}

type Students struct {
	baseStaff.InfoMapList
}

type CourseItems []*CourseItem

func (c CourseItems) Len() int {
	return len(c)
}

func (c CourseItems) Less(i, j int) bool {
	return c[i].CourseID < c[j].CourseID
}

func (c CourseItems) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (x *Courses) MarshalJSON() ([]byte, error) {
	list := CourseItems{}
	for _, v := range x.Items {
		list = append(list, v)
	}
	sort.Sort(list)
	return json.Marshal(list)
}

func (x *Courses) UnmarshalJSON(bytes []byte) error {
	*x = Courses{Items: map[string]*CourseItem{}}
	list := CourseItems{}
	err := json.Unmarshal(bytes, &list)
	if err != nil {
		return err
	}
	for _, v := range list {
		x.Items[v.CourseID] = v
	}
	return nil
}

type ScheduleItems []*ScheduleItem

func (c ScheduleItems) Len() int {
	return len(c)
}

func (c ScheduleItems) Less(i, j int) bool {
	return c[i].WeekDay < c[j].WeekDay
}

func (c ScheduleItems) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (x *Schedule) MarshalJSON() ([]byte, error) {
	list := ScheduleItems{}
	for _, v := range x.Items {
		list = append(list, v)
	}
	sort.Sort(list)
	return json.Marshal(list)
}

func (x *Schedule) UnmarshalJSON(bytes []byte) error {
	*x = Schedule{Items: map[string]*ScheduleItem{}}
	list := ScheduleItems{}
	err := json.Unmarshal(bytes, &list)
	if err != nil {
		return err
	}
	for _, v := range list {
		x.Items[fmt.Sprint(v.Week, v.WeekDay, v.Section)] = v
	}
	return nil
}
