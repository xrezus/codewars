package valid_braces

import "testing"

func TestValidBraces(t *testing.T) {
	type args struct {
		n string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"(){}[]"}, true},
		{"2", args{"([{}])"}, true},
		{"3", args{"(}"}, false},
		{"4", args{"[(])"}, false},
		{"5", args{"[({)](]"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidBraces(tt.args.n); got != tt.want {
				t.Errorf("ValidBraces() = %v, want %v", got, tt.want)
			}
		})
	}
}
