package supermarket_queue

import "testing"

func TestQueueTime(t *testing.T) {
	type args struct {
		n []int
		m int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{}, 1}, 0},
		{"", args{[]int{1, 2, 3, 4}, 1}, 10},
		{"", args{[]int{2, 2, 3, 3, 4, 4}, 2}, 9},
		{"", args{[]int{1, 2, 3, 4, 5}, 100}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueueTime(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("QueueTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
