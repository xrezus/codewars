package grasshopper_grade_book

import "testing"

func TestGetGrade(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}

	tests := []struct {
		name string
		args args
		want rune
	}{
		{"", args{95, 90, 93}, 'A'},
		{"", args{82, 85, 87}, 'B'},
		{"", args{60, 82, 76}, 'C'},
		{"", args{58, 62, 70}, 'D'},
		{"", args{48, 55, 52}, 'F'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGrade(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("GetGrade() = %v, want %v", got, tt.want)
			}
		})
	}
}
