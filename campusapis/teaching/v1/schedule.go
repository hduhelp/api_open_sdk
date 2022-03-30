package teachingv1

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/hduhelp/api_open_sdk/baseStaff"
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
	// 初始化 Teachers 和 Students ，保证非空，防止调用方出现空指针错误
	x.InitMember()
	switch q.QueryStaff.Type {
	case baseStaff.Type_Teacher:
		//教师展示上课学生
		x.AddStudent(r)
	case baseStaff.Type_Undergraduate, baseStaff.Type_Postgraduate:
		//学生添加展示授课教师
		x.AddTeacher(r)
	}
}

func (x *ScheduleItem) InitMember() {
	if x.Teachers == nil {
		x.Teachers = &baseStaff.InfoMapList{
			Items: map[string]*baseStaff.Info{},
		}
	}
	if x.Students == nil {
		x.Students = &baseStaff.InfoMapList{
			Items: map[string]*baseStaff.Info{},
		}
	}
}

func (x *ScheduleItem) AddTeacher(r ScheduleReader) {
	x.Teachers.Append(r.TeacherInfo())
}

func (x *ScheduleItem) AddStudent(r ScheduleReader) {
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

func (x *Courses) MarshalJSON() ([]byte, error) {
	list := CourseItems{}
	for _, v := range x.Items {
		list = append(list, v)
	}
	sort.Slice(list, func(i, j int) bool {
		return courseItemsLess(list, i, j)
	})
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

func (x *Schedule) MarshalJSON() ([]byte, error) {
	list := ScheduleItems{}
	for _, v := range x.Items {
		list = append(list, v)
	}
	sort.SliceStable(list, func(i, j int) bool {
		return scheduleItemsLess(list, i, j)
	})
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