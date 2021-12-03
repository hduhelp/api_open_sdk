package teaching

import "sort"

// FilterThisWeek 过滤出本周课程, 会原地修改原课表
// 注意：只是根据接口返回的 IsThisWeek 字段来判断是否是本周课程
// 如果采用的是salmon base 优先级 3 的请求方式，则需要自己计算是否是本周课程
func (courses *Courses) FilterThisWeek() *Courses {
	if len(courses.Items) == 0 {
		return courses
	}
	for _, course := range courses.Items {
		if course.Schedule == nil {
			continue
		}
		for scheduleID, schedule := range course.Schedule.Items {
			if !schedule.IsThisWeek {
				delete(course.Schedule.Items, scheduleID)
			}
		}
	}
	return courses
}

// FilterByWeekdays 过滤出星期几的课程，会原地修改原课表
func (courses *Courses) FilterByWeekdays(weekdays ...int32) *Courses {
	in := func(i int32, list []int32) bool {
		for _, v := range list {
			if i == v {
				return true
			}
		}
		return false
	}
	if len(courses.Items) == 0 {
		return courses
	}
	for _, course := range courses.Items {
		if course.Schedule == nil {
			continue
		}
		for scheduleID, schedule := range course.Schedule.Items {
			if !in(schedule.WeekDay, weekdays) {
				delete(course.Schedule.Items, scheduleID)
			}
		}
	}
	return courses
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

type scheduleWithCourseSlice struct {
	schedules []*scheduleWithCourse
	less      ScheduleLess
}

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
	schedules := scheduleWithCourseSlice{less: less}
	for _, course := range courses.Items {
		if course.Schedule == nil {
			continue
		}
		for _, schedule := range course.Schedule.Items {
			schedules.append(&scheduleWithCourse{CourseName: course.CourseName, ScheduleItem: schedule})
		}
	}

	length := schedules.Len()
	sort.Sort(schedules)

	return func() (courseName string, schedule *ScheduleItem, ok bool) {
		if index >= length {
			return
		}
		courseName, schedule, ok = schedules.schedules[index].CourseName, schedules.schedules[index].ScheduleItem, true
		index++
		return
	}
}

func (s *scheduleWithCourseSlice) append(item *scheduleWithCourse) {
	s.schedules = append(s.schedules, item)
}

func (s scheduleWithCourseSlice) Len() int {
	return len(s.schedules)
}

func (s scheduleWithCourseSlice) Less(i, j int) bool {
	return s.less(s.schedules[i].ScheduleItem, s.schedules[j].ScheduleItem)
}

func (s scheduleWithCourseSlice) Swap(i, j int) {
	s.schedules[i], s.schedules[j] = s.schedules[j], s.schedules[i]
}

func DefaultSort(s1, s2 *ScheduleItem) bool {
	// 只考虑本周课程
	// 先按星期升序，再按开始节次升序
	switch {
	case s1.WeekDay < s2.WeekDay:
		return true
	case s1.WeekDay > s2.WeekDay:
		return false
	case s1.WeekDay == s2.WeekDay:
		switch {
		case s1.StartSection() < s2.StartSection():
			return true
		case s1.StartSection() > s2.StartSection():
			return false
		default:
			return true
		}
	}
	return true
}

func Reverse(less ScheduleLess) ScheduleLess {
	return func(s1, s2 *ScheduleItem) bool {
		return !less(s1, s2)
	}
}
