package which_are_in

import (
	"reflect"
	"testing"
)

func TestInArray(t *testing.T) {
	type args struct {
		array1 []string
		array2 []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"live", "arp", "strong"}, []string{"lively", "alive", "harp", "sharp", "armstrong"}}, []string{"arp", "live", "strong"}},
		{"2", args{[]string{"cod", "code", "wars", "ewar", "ar"}, []string{}}, []string{}},
		{"3", args{[]string{"tarp", "mice", "bull"}, []string{"lively", "alive", "harp", "sharp", "armstrong"}}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.args.array1, tt.args.array2); reflect.DeepEqual(got, tt.want) {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
