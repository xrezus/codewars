package you_can_code_under_pressure

import "testing"

func TestDoubleInteger(t *testing.T) {
	type args struct {
		a int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{1}, 2},
		{"", args{5}, 10},
		{"", args{15}, 30},
		{"", args{2000}, 4000},
		{"", args{4}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoubleInteger(tt.args.a); got != tt.want {
				t.Errorf("DoubleInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
