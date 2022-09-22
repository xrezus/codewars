package grasshopper_summation

import "testing"

func TestSummation(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{2}, 3},
		{"", args{4}, 10},
		{"", args{8}, 36},
		{"", args{22}, 253},
		{"", args{100}, 5050},
		{"", args{213}, 22791},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Summation(tt.args.n); got != tt.want {
				t.Errorf("Summation() = %v, want %v", got, tt.want)
			}
		})
	}
}
