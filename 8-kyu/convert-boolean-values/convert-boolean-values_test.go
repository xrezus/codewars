package convert_boolean_values

import "testing"

func TestBoolToWord(t *testing.T) {
	type args struct {
		n bool
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{true}, "Yes"},
		{"", args{false}, "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolToWord(tt.args.n); got != tt.want {
				t.Errorf("BoolToWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
