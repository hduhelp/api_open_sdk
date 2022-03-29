package teachingv1

import (
	"errors"

	"github.com/hduhelp/api_open_sdk/baseStaff"
	schooltime "github.com/hduhelp/api_open_sdk/campusapis/schooltime"
)

type CourseQuery struct {
	QueryStaff *baseStaff.Staff
	Queries    []Queryable

	*schooltime.SchoolDate

	schooltime.SchoolDateToDater
	schooltime.SectionReader

	*Courses
}

type Queryable interface {
	GetCourses(staff *baseStaff.Staff, semester *schooltime.Semester, schoolYear *schooltime.SchoolYear) ([]CourseReader, error)
}

func NewCourseQuery(staff *baseStaff.Staff,
	st *schooltime.SchoolDate,
	dater schooltime.SchoolDateToDater,
	sectionReader schooltime.SectionReader,
	q ...Queryable) *CourseQuery {
	return &CourseQuery{
		QueryStaff:        staff,
		SchoolDate:        st,
		SchoolDateToDater: dater,
		SectionReader:     sectionReader,
		Courses:           &Courses{Items: map[string]*CourseItem{}},
		Queries:           q,
	}
}

func (q *CourseQuery) GetCourses() (*CourseQuery, error) {
	if q.QueryStaff.Id == "" {
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
	q.Courses.Items = make(map[string]*CourseItem)
	//合并课程信息内容到标准课程信息中
	for _, v := range courseReaders {
		if q.Courses.Items[v.CourseID()] == nil {
			q.Courses.Items[v.CourseID()] = v.CourseInfo()
		}
		q.Courses.Items[v.CourseID()].AddSchedule(q, v.ScheduleReader())
	}
	return q, nil
}
