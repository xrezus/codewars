package convert_a_string_to_a_number

import "testing"

func TestStringToNumber(t *testing.T) {
	type args struct {
		n string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"1234"}, 1234},
		{"", args{"605"}, 605},
		{"", args{"1405"}, 1405},
		{"", args{"-7"}, -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToNumber(tt.args.n); got != tt.want {
				t.Errorf("StringToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
