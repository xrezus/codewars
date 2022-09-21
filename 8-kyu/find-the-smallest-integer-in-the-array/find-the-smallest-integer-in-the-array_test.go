package find_the_smallest_integer_in_the_array

import "testing"

func TestSmallestIntegerFinder(t *testing.T) {
	type args struct {
		n []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{34, 15, 88, 2}}, 2},
		{"", args{[]int{34, -345, -1, 100}}, -345},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestIntegerFinder(tt.args.n); got != tt.want {
				t.Errorf("SmallestIntegerFinder() = %v, want %v", got, tt.want)
			}
		})
	}
}
