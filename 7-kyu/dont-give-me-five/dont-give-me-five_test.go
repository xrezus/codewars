package dont_give_me_five

import "testing"

func TestDontGiveMeFive(t *testing.T) {
	type args struct {
		n int
		m int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{1, 9}, 8},
		{"", args{4, 17}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DontGiveMeFive(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("DontGiveMeFive() = %v, want %v", got, tt.want)
			}
		})
	}
}
