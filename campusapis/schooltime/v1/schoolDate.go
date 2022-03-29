package schooltimev1

import (
	"errors"
	"time"

	"github.com/hduhelp/api_open_sdk/types"
)

type SchoolDate struct {
	d          *types.Date
	SchoolYear *SchoolYear
	Semester   *Semester
	Week       int32
}

func (d SchoolDate) Valid() bool {
	return d.SchoolYear.Year != 0 && d.Semester.Num != 0
}

func (d SchoolDate) Time() time.Time {
	return d.d.Time
}

func (d SchoolDate) Date() *types.Date {
	return d.d
}

func (SchoolDate) FromTime(t time.Time, handler TimeToSchoolDater) (*SchoolDate, error) {
	d := handler.GetSchoolDateFromTime(t)
	d.d = types.DateFromTime(t)
	if !d.Valid() {
		return nil, errors.New("invalid school date")
	}
	return d, nil
}

// TimeToSchoolDater 学校日期接口，从时间获得学年、学期
type TimeToSchoolDater interface {
	GetSchoolDateFromTime(t time.Time) *SchoolDate
}

// SchoolDateToDater 从学校日期获得日期
type SchoolDateToDater interface {
	GetSchoolDateFrom(schoolDate *SchoolDate, weekDay int32) (*types.Date, error)
}
