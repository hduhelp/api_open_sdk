package teachingv1

import (
	"errors"

	"github.com/hduhelp/api_open_sdk/campusapis/schoolTime"
	"github.com/hduhelp/api_open_sdk/campusapis/staff"
)

type CourseQuery struct {
	QueryStaff *staff.Staff
	Queries    []Queryable

	*schoolTime.SchoolDate

	schoolTime.SchoolDateToDater
	schoolTime.SectionReader

	Courses CourseReaders
}

type Queryable interface {
	GetCourses(staff *staff.Staff, semester *schoolTime.Semester, schoolYear *schoolTime.SchoolYear) ([]CourseReader, error)
}

func NewCourseQuery(staff *staff.Staff, st *schoolTime.SchoolDate, dater schoolTime.SchoolDateToDater, sectionReader schoolTime.SectionReader, q ...Queryable) *CourseQuery {
	return &CourseQuery{
		QueryStaff:        staff,
		SchoolDate:        st,
		SchoolDateToDater: dater,
		SectionReader:     sectionReader,
		Queries:           q,
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
	q.Courses = make([]CourseReader, 0)
	for _, v := range q.Queries {
		courses, err := v.GetCourses(q.QueryStaff, q.SchoolDate.Semester, q.SchoolDate.SchoolYear)
		if err != nil {
			return q, err
		}
		q.Courses = append(q.Courses, courses...)
	}
	return q, nil
}
