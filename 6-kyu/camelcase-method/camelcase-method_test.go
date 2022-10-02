package camelcase_method

import "testing"

func TestCamelCase(t *testing.T) {
	type args struct {
		n string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"test case"}, "TestCase"},
		{"", args{"camel case method"}, "CamelCaseMethod"},
		{"", args{"say hello "}, "SayHello"},
		{"", args{" camel case word"}, "CamelCaseWord"},
		{"", args{""}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.args.n); got != tt.want {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
