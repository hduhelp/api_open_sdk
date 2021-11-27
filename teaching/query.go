package teaching

import (
	"errors"
	"github.com/hduhelp/api_open_sdk/baseStaff"
	"github.com/hduhelp/api_open_sdk/schoolTime"
)

type CourseQuery struct {
	QueryStaff *baseStaff.Staff
	Queries    []Queryable

	*schoolTime.SchoolDate

	*Course
}

type Queryable interface {
	GetCourses(staff *baseStaff.Staff, semester *schoolTime.Semester, schoolYear *schoolTime.SchoolYear) ([]CourseReader, error)
}

func NewCourseQuery(staff *baseStaff.Staff, st *schoolTime.SchoolDate, q ...Queryable) *CourseQuery {
	return &CourseQuery{
		QueryStaff: staff,
		SchoolDate: st,
		Course:     &Course{Items: map[string]*CourseItem{}},
		Queries:    q,
	}
}

func (q *CourseQuery) GetCourses() (*CourseQuery, error) {
	if q.QueryStaff.ID == "" {
		return q, errors.New("staffID cannot be nil")
	}
	if q.SchoolDate == nil {
		return q, errors.New("schoolDate cannot be nil")
	}

	if !q.SchoolDate.Valid() {
		return q, errors.New("cannot get school date from dater")
	}
	//从课程信息接口查询对应学期课程信息
	courseReaders := make([]CourseReader, 0)
	for _, v := range q.Queries {
		courses, err := v.GetCourses(q.QueryStaff, q.SchoolDate.Semester, q.SchoolDate.SchoolYear)
		if err != nil {
			return q, err
		}
		courseReaders = append(courseReaders, courses...)
	}
	q.Course.Items = make(map[string]*CourseItem)
	//合并课程信息内容到标准课程信息中
	for _, v := range courseReaders {
		if q.Course.Items[v.CourseID()] == nil {
			q.Course.Items[v.CourseID()] = v.CourseInfo()
		} else {
			q.Course.Items[v.CourseID()].AddSchedule(v.ScheduleReader(), q)
		}
	}
	return q, nil
}
