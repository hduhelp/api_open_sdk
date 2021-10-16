package schedule

import (
	"errors"
	"github.com/hduhelp/api_open_sdk/baseStaff"
	"github.com/hduhelp/api_open_sdk/schoolTime"
	"time"
)

type CourseQuery struct {
	QueryStaff *baseStaff.Staff
	Queries    []Queryable
	time.Time
	dater schoolTime.SchoolDater

	*Schedule
}

type Queryable interface {
	GetCourses(staff *baseStaff.Staff, semester *schoolTime.Semester, schoolYear *schoolTime.SchoolYear) ([]CourseReader, error)
}

func NewCourseQuery(staff *baseStaff.Staff, d schoolTime.SchoolDater, t time.Time, q ...Queryable) *CourseQuery {
	return &CourseQuery{
		QueryStaff: staff,
		Time:       t,
		Schedule:   &Schedule{Items: map[string]*ScheduleItem{}},
		Queries:    q,

		dater: d,
	}
}

func (q CourseQuery) GetCourses() (CourseQuery, error) {
	if q.QueryStaff.ID == "" {
		return q, errors.New("staffID cannot be nil")
	}
	if q.dater == nil {
		return q, errors.New("dater cannot be nil")
	}
	schoolDate := q.dater.GetSchoolDateFromTime(q.Time)
	if !schoolDate.Valid() {
		return q, errors.New("cannot get school date from dater")
	}
	//从课程信息接口查询对应学期课程信息
	readers := make([]CourseReader, 0)
	for _, v := range q.Queries {
		courses, err := v.GetCourses(q.QueryStaff, schoolDate.Semester, schoolDate.SchoolYear)
		if err != nil {
			return q, err
		}
		readers = append(readers, courses...)
	}
	//合并课程信息内容到标准
	for _, v := range readers {
		scheduleID := v.ScheduleID()
		if q.Schedule.Items[scheduleID] == nil {
			q.Schedule.Items[scheduleID] = v.CourseInfo(q.Time)
		}
		q.Schedule.Items[scheduleID].AddMember(v, q.QueryStaff.Type)
	}
	return q, nil
}
