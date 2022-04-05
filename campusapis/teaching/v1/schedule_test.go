package teachingv1

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hduhelp/api_open_sdk/campusapis/staff"
)

func TestCourses_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Items map[string]*CourseItem
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				bytes: []byte(`[{"ClassName":"12班（计算机）","CourseID":"G123002","CourseName":"新时代中国特色社会主义理论与实践研究","Credit":"2","Schedule":[{"EndTime":1638358500,"IsThisWeek":true,"Location":"6-中227","Section":[3,4],"StartTime":1638352800,"Teachers":[{"StaffID":"41321","StaffName":"戚陈炯","StaffType":3}],"Week":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17],"WeekDay":3}],"SchoolYear":"2021-2022","Semester":1}]`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Courses{
				Items: tt.fields.Items,
			}
			if err := x.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSchedule_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Items map[string]*ScheduleItem
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				bytes: []byte(`[{"EndTime":1638358500,"IsThisWeek":true,"Location":"6-中227","Section":[3,4],"StartTime":1638352800,"Teachers":[{"StaffID":"41321","StaffName":"戚陈炯","StaffType":3}],"Week":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17],"WeekDay":3}]`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Schedule{
				Items: tt.fields.Items,
			}
			if err := x.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSchedule_MarshalJSON(t *testing.T) {
	type fields struct {
		Items map[string]*ScheduleItem
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				Items: map[string]*ScheduleItem{
					"1": {
						EndTime:    1638358500,
						IsThisWeek: true,
						Location:   "6-中227",
						Section:    []int32{3, 4},
						StartTime:  1638352800,
						Teachers:   &staff.InfoMapList{},
						Week:       []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
						WeekDay:    3,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Schedule{
				Items: tt.fields.Items,
			}
			got, err := x.MarshalJSON()
			fmt.Println(string(got))
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
