package schoolConst

import (
	"math"
	"time"
)

const (
	SchoolYear             = 2025
	Semester               = 1
	SemesterStartTimestamp = int64(1757865600) // 2025-9-15 00:00:00
)

func getWeekNum() int {
	timeStart := SemesterStartTimestamp
	return int(math.Floor(float64(time.Now().Unix()-timeStart)/(86400*7))) + 1
}

func getWeekDay() int {
	day := int(time.Now().Weekday())
	return day
}

type TimeData struct {
	SchoolYear             int   `json:"schoolYear"`
	Semester               int   `json:"semester"`
	SemesterStartTimestamp int64 `json:"semester_start_timestamp"`
	WeekNow                int   `json:"weekNow"`
	WeekDayNow             int   `json:"weekDayNow"`
	Timestamp              int64 `json:"timeStamp"`
	Section                int   `json:"section"`
}

func GetTimeData() *TimeData {
	timeNow := time.Now()
	hour := timeNow.Hour()
	minute := time.Now().Minute() + hour*60
	tMap := []int{0, 530, 580, 645, 695, 745, 855, 905, 960, 1010, 1155, 1205, 1255}
	section := -1
	for i := 0; i < len(tMap); i++ {
		if minute < tMap[i] {
			section = i
			break
		}
	}

	timed := &TimeData{
		SchoolYear:             SchoolYear,
		Semester:               Semester,
		SemesterStartTimestamp: SemesterStartTimestamp,
		WeekNow:                getWeekNum(),
		WeekDayNow:             getWeekDay(),
		Timestamp:              time.Now().Unix(),
		Section:                section,
	}

	return timed
}
