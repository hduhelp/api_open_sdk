package teachingv1

import (
	"fmt"
	"testing"
)

func getCourses() *Courses {
	return &Courses{Items: map[string]*CourseItem{
		"课程id1": &CourseItem{
			CourseName: "课程名称1",
			Schedule: &Schedule{Items: map[string]*ScheduleItem{
				// 此 schedule 在本测试中应当被过滤
				"课程1-a-应当被过滤": &ScheduleItem{
					IsThisWeek: true,
					Teachers:   nil,
					//Week:       []int32{1, 2, 3, 4},
					WeekDay: 3,
					Section: []int32{6, 7, 8, 9},
					// 借用一下 Location 字段用来更明显地测试，其实这个是教室字段
					Location: "应当被过滤",
				},
				// 此 schedule 在本测试中应当排在第四
				"课程1-b-应当第四": &ScheduleItem{
					IsThisWeek: true,
					Teachers:   nil,
					//Week:       []int32{1, 2, 3, 4},
					WeekDay:  2,
					Section:  []int32{3, 4},
					Location: "应当第四",
				},
			},
			},
		},
		"课程id2": &CourseItem{
			CourseName: "课程名称2",
			Schedule: &Schedule{Items: map[string]*ScheduleItem{
				// 此 schedule 在本测试中应当排在第五
				"课程2-a-应当第五": &ScheduleItem{
					IsThisWeek: true,
					Teachers:   nil,
					WeekDay:    2,
					Section:    []int32{6, 7, 8, 9},
					Location:   "应当第五",
				},
				// 此 schedule 在本测试中应当排在第二
				"课程2-b-应当第二": &ScheduleItem{
					IsThisWeek: true,
					Teachers:   nil,
					WeekDay:    1,
					Section:    []int32{3, 4},
					Location:   "应当第二",
				},
			},
			},
		},
		"课程id3": &CourseItem{
			CourseName: "课程名称3",
			Schedule: &Schedule{Items: map[string]*ScheduleItem{
				// 此 schedule 在本测试中应当排在第三
				"课程3-a-应当第三": &ScheduleItem{
					IsThisWeek: true,
					Teachers:   nil,
					WeekDay:    1,
					Section:    []int32{6, 7, 8, 9},
					Location:   "应当第三",
				},
				// 此 schedule 在本测试中应当排在第一
				"课程3-b-应当第一": &ScheduleItem{
					IsThisWeek: true,
					Teachers:   nil,
					WeekDay:    1,
					Section:    []int32{1, 2},
					Location:   "应当第一",
				},
			},
			},
		},
	}}
}

func TestScheduleProcess(t *testing.T) {
	courses := getCourses()
	courses.FilterThisWeek().FilterByWeekdays(1, 2)
	// 这里测试一下反序函数，所以实际输出一二三四五会反序，懒得改中文表述了
	it := courses.Iterator(Reverse(DefaultSort))
	for {
		courseName, schedule, ok := it()
		if !ok {
			break
		}
		fmt.Println(courseName, schedule)
	}
}
