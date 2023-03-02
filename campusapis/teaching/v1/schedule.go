package teachingv1

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/hduhelp/api_open_sdk/campusapis/schoolTime"
	"github.com/hduhelp/api_open_sdk/campusapis/staff"
)

// AddSchedule 添加课表
func (x *CourseItem) AddSchedule(timeable schoolTime.Timeable, optionable ShowMemberOptionable, r ScheduleReader) {
	if x.Schedule == nil {
		x.Schedule = &Schedule{Items: map[string]*ScheduleItem{}}
	}
	if x.Schedule.Items[r.ScheduleID()] == nil {
		x.Schedule.Items[r.ScheduleID()] = r.ScheduleInfo()
	}

	x.Schedule.Items[r.ScheduleID()].AddMember(optionable, r)
}

//func (x *ScheduleItem) SetTime(timeable schoolTime.Timeable) {
//	//非当前周课程 跳过时间赋值
//	if _, ok := InArray(x.Week, timeable.GetSchoolDate().Week); !ok {
//		return
//	}
//	d, err := timeable.GetSchoolDateToDater().GetSchoolDateFrom(timeable.GetSchoolDate(), x.WeekDay)
//	if err != nil {
//		return
//	}
//	x.StartTime = timeable.GetSectionReader().StartTime(d, FirstOfArray(x.Section)).Unix()
//	x.EndTime = timeable.GetSectionReader().EndTime(d, LastOfArray(x.Section)).Unix()
//	x.IsThisWeek = true
//}

type ShowMemberOptionable interface {
	GetOptionShowMemberOption() OptionShowMember
}

type OptionShowMember int32

const (
	OptionShowMemberNone OptionShowMember = iota
	OptionShowMemberStudent
	OptionShowMemberTeacher
	OptionShowMemberAll
)

// AddMember 添加课程人员
func (x *ScheduleItem) AddMember(optionable ShowMemberOptionable, r ScheduleReader) {
	// 初始化 Teachers 和 Students ，保证非空，防止调用方出现空指针错误
	x.InitMember()
	switch optionable.GetOptionShowMemberOption() {
	case OptionShowMemberAll:
		x.AddStudent(r)
		x.AddTeacher(r)
	case OptionShowMemberStudent:
		x.AddStudent(r)
	case OptionShowMemberTeacher:
		x.AddTeacher(r)
	}
}

func (x *ScheduleItem) InitMember() {
	if x.Teachers == nil {
		x.Teachers = &staff.InfoMapList{
			Items: map[string]*staff.Info{},
		}
	}
	if x.Students == nil {
		x.Students = &staff.InfoMapList{
			Items: map[string]*staff.Info{},
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
	GlobalSchedule() *GlobalSchedule
}

type CourseReaders []CourseReader

func (readers CourseReaders) ToTeachingV1Courses(timeable schoolTime.Timeable, optionable ShowMemberOptionable) (courses *Courses) {
	courses = &Courses{
		Items: make(map[string]*CourseItem),
	}
	//合并课程信息内容到标准课程信息中
	for _, v := range readers {
		if courses.Items[v.CourseID()] == nil {
			courses.Items[v.CourseID()] = v.CourseInfo()
		}
		courses.Items[v.CourseID()].AddSchedule(timeable, optionable, v.ScheduleReader())
	}
	return courses
}

func (readers CourseReaders) ToTeachingV1GlobalSchedule() (courses []*GlobalSchedule) {
	courses = make([]*GlobalSchedule, 0, len(readers))
	//合并课程信息内容到标准课程信息中
	for _, v := range readers {
		courses = append(courses, v.GlobalSchedule())
	}
	return courses
}

type ScheduleReader interface {
	ScheduleID() string
	ScheduleInfo() *ScheduleItem
	TeacherInfo() *staff.Info
	StudentInfo() *staff.Info
}

type Teachers struct {
	staff.InfoMapList
}

type Students struct {
	staff.InfoMapList
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
