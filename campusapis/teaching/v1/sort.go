package teachingv1

import (
	"encoding/json"
	"sort"
)

type CoursesCustomLess struct {
	*Courses
	Less func(c CourseItems, i, j int) bool
}

func (x *CoursesCustomLess) MarshalJSON() ([]byte, error) {
	list := CourseItems{}
	for _, v := range x.Items {
		list = append(list, v)
	}
	sort.Slice(list, func(i, j int) bool {
		return x.Less(list, i, j)
	})
	return json.Marshal(list)
}

var courseItemsLess = DefaultCourseItemsLess

func SetCourseItemLess(less func(c CourseItems, i, j int) bool) {
	courseItemsLess = less
}

var DefaultCourseItemsLess = func(c CourseItems, i, j int) bool {
	return c[i].CourseID < c[j].CourseID
}

var CourseItemsLessByStartTime = func(c CourseItems, i, j int) bool {
	return c[i].minStartTime() < c[j].minStartTime()
}

func (x *CourseItem) minStartTime() int64 {
	if x.Schedule == nil || x.Schedule.Items == nil || len(x.Schedule.Items) == 0 {
		return 0
	}
	min := int32(0)
	for _, v := range x.Schedule.Items {
		if len(v.Section) > 0 && (min == 0 || v.Section[0] < min) {
			min = v.Section[0]
		}
	}
	return int64(min)
}

var scheduleItemsLess = DefaultScheduleItemsLess

func SetScheduleItemLess(less func(c ScheduleItems, i, j int) bool) {
	scheduleItemsLess = less
}

var DefaultScheduleItemsLess = func(c ScheduleItems, i, j int) bool {
	return c[i].WeekDay < c[j].WeekDay
}

var ScheduleItemsLessByStartTime = func(c ScheduleItems, i, j int) bool {
	return c[i].StartTime < c[j].StartTime
}
