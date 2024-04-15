package first_non_repeating_character

import "testing"

func TestLastDigit(t *testing.T) {
	type args struct {
		a string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{""}, ""},
		{"", args{"stress"}, "t"},
		{"", args{"abba"}, ""},
		{"", args{"hello world, eh?"}, "w"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstNonRepeating(tt.args.a); got != tt.want {
				t.Errorf("LastDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
