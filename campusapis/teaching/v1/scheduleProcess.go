package teachingv1

import (
	"math"
	"sort"
	"time"
)

type RemainFunc func(s *ScheduleItem) bool

// FilterBySchedule 根据 Schedule 的属性判断函数过滤符合要求的课表
func (courses *Courses) FilterBySchedule(remains ...RemainFunc) *Courses {
	// 复制一份Courses
	var newCourses Courses
	newCourses.Items = make(map[string]*CourseItem)
	for key, course := range courses.Items {
		var newSchedule Schedule
		newSchedule.Items = make(map[string]*ScheduleItem)
		for scheduleID, schedule := range course.Schedule.Items {
			flag := true
			for _, remain := range remains {
				if !remain(schedule) { // 如果有任何一个不满足，则不保留
					flag = false
					break
				}
			}
			if flag {
				newSchedule.Items[scheduleID] = schedule
			}
		}
		if len(newSchedule.Items) > 0 {
			newCourses.Items[key] = &CourseItem{
				ClassID:    course.ClassID,
				ClassName:  course.ClassName,
				CourseID:   course.CourseID,
				CourseName: course.CourseName,
				ClassTime:  course.ClassTime,
				Credit:     course.Credit,
				SchoolYear: course.SchoolYear,
				Semester:   course.Semester,
				Schedule:   &newSchedule,
			}
		}
	}

	return &newCourses
}

//// FilterThisWeek 过滤出本周课程, 会原地修改原课表
//// 注意：只是根据接口返回的 IsThisWeek 字段来判断是否是本周课程
//// 如果采用的是salmon base 优先级 3 的请求方式，则需要自己计算是否是本周课程
//func (courses *Courses) FilterThisWeek() *Courses {
//	return courses.FilterBySchedule(RemainThisWeek)
//}
//
//// FilterByWeekdays 过滤出星期几的课程，会原地修改原课表
//func (courses *Courses) FilterByWeekdays(weekdays ...int32) *Courses {
//	return courses.FilterBySchedule(RemainByWeekdays(weekdays...))
//}

const timeStart = 1709481600 // 2023-03-04 00:00:00

// RemainThisWeek 留下本星期的课程
func RemainThisWeek(day time.Time) RemainFunc {
	return func(s *ScheduleItem) bool {
		weekNow := int32(math.Floor(float64(day.Unix()-timeStart)/(86400*7))) + 1
		for _, v := range s.Week {
			if v == weekNow {
				return true
			}
		}
		return false
	}
}

//// RemainByWeekdays 留下指定星期几的课程
//func RemainByWeekdays(weekdays ...int32) RemainFunc {
//	in := func(i int32, list []int32) bool {
//		for _, v := range list {
//			if i == v {
//				return true
//			}
//		}
//		return false
//	}
//	return func(s *ScheduleItem) bool {
//		return in(s.WeekDay, weekdays)
//	}
//}

// RemainByWeekDay 留下当天的课程
func RemainByWeekDay(day time.Time) RemainFunc {
	return func(s *ScheduleItem) bool {
		return s.WeekDay == int32(day.Weekday())
	}
}

func (schedule *ScheduleItem) StartSection() int {
	return int(schedule.Section[0])
}

func (schedule *ScheduleItem) EndSection() int {
	return int(schedule.Section[len(schedule.Section)-1])
}

type scheduleWithCourse struct {
	*ScheduleItem
	CourseName string
}

type ScheduleLess func(s1, s2 *ScheduleItem) bool

// Iterator 遍历课表，返回课程名，ScheduleItem，可自定义排序方法
// 用法详见 test
func (courses *Courses) Iterator(less ScheduleLess) func() (courseName string, schedule *ScheduleItem, ok bool) {
	if len(courses.Items) == 0 {
		return func() (courseName string, schedule *ScheduleItem, ok bool) {
			return
		}
	}
	index := 0
	// 把课表拉平
	schedules := make([]*scheduleWithCourse, 0, len(courses.Items))
	for _, course := range courses.Items {
		if course.Schedule == nil {
			continue
		}
		for _, schedule := range course.Schedule.Items {
			schedules = append(schedules, &scheduleWithCourse{CourseName: course.CourseName, ScheduleItem: schedule})
		}
	}

	length := len(schedules)
	sort.Slice(schedules, func(i, j int) bool {
		return less(schedules[i].ScheduleItem, schedules[j].ScheduleItem)
	})

	return func() (courseName string, schedule *ScheduleItem, ok bool) {
		if index >= length {
			return
		}
		courseName, schedule, ok = schedules[index].CourseName, schedules[index].ScheduleItem, true
		index++
		return
	}
}

func DefaultSort(s1, s2 *ScheduleItem) bool {
	// 只考虑本周课程
	// 依次按星期升序，开始节次升序，结束节次升序
	switch {
	case s1.WeekDay < s2.WeekDay:
		return true
	case s1.WeekDay > s2.WeekDay:
		return false
	}
	switch {
	case s1.StartSection() < s2.StartSection():
		return true
	case s1.StartSection() > s2.StartSection():
		return false
	}
	switch {
	case s1.EndSection() < s2.EndSection():
		return true
	case s1.EndSection() < s2.EndSection():
		return false
	}
	return true
}

func Reverse(less ScheduleLess) ScheduleLess {
	return func(s1, s2 *ScheduleItem) bool {
		return !less(s1, s2)
	}
}
