package mexican_wave

import (
	"fmt"
	"testing"
)

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestWave(t *testing.T) {
	type args struct {
		a string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{" x yz"}, []string{" X yz", " x Yz", " x yZ"}},
		{"2", args{"abc"}, []string{"Abc", "aBc", "abC"}},
		{"3", args{" ab  c"}, []string{" Ab  c", " aB  c", " ab  C"}},
		{"4", args{""}, []string{}},
		{"5", args{"z"}, []string{"Z"}},
		{"6", args{"a a a a a"}, []string{"A a a a a", "a A a a a", "a a A a a", "a a a A a", "a a a a A"}},
		{"7", args{"aaaaa"}, []string{"Aaaaa", "aAaaa", "aaAaa", "aaaAa", "aaaaA"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wave(tt.args.a); Equal(got, tt.want) {

				fmt.Printf("%#v, %#v", got, tt.want)

				t.Errorf("Wave() = %v, want %v", got, tt.want)
			}
		})
	}
}
