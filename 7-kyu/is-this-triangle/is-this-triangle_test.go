package is_this_triangle

import "testing"

func TestIsTriangle(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{5, 1, 2}, false},
		{"", args{1, 2, 5}, false},
		{"", args{2, 5, 1}, false},
		{"", args{4, 2, 3}, true},
		{"", args{5, 1, 5}, true},
		{"", args{2, 2, 2}, true},
		{"", args{-1, 2, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTriangle(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("IsTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
