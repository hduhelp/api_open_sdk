package schoolTime

import (
	"github.com/hduhelp/api_open_sdk/types"
	"time"
)

type SchoolDate struct {
	types.Date
	*SchoolYear
	*Semester
}

func (d SchoolDate) Valid() bool {
	return d.SchoolYear.Year != 0 && d.Semester.Num != 0
}

// SchoolDater 学校日期接口，从时间获得学年、学期
type SchoolDater interface {
	GetSchoolDateFromTime(t time.Time) *SchoolDate
}
