package types

import (
	"reflect"
	"testing"
)

func TestBitToList(t *testing.T) {
	type args struct {
		bits int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{bits: 1536 * 2},
			want: []int{10, 11},
		},
		{
			name: "test2",
			args: args{bits: 131069 * 2},
			want: []int{1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
		},
		{
			name: "test3",
			args: args{bits: 16 * 2},
			want: []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BitToList(tt.args.bits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BitToList() = %v, want %v", got, tt.want)
			}
		})
	}
}
