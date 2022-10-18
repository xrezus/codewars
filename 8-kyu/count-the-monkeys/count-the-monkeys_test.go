package count_the_monkeys

import (
	"reflect"
	"testing"
)

func TestMonkeyCount(t *testing.T) {
	type args struct {
		a int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{1}, []int{1}},
		{"", args{2}, []int{1, 2}},
		{"", args{3}, []int{1, 2, 3}},
		{"", args{4}, []int{1, 2, 3, 4}},
		{"", args{5}, []int{1, 2, 3, 4, 5}},
		{"", args{6}, []int{1, 2, 3, 4, 5, 6}},
		{"", args{7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"", args{8}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"", args{9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"", args{10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonkeyCount(tt.args.a); reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
