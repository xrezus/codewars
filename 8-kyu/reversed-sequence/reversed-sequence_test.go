package reversed_sequence

import (
	"reflect"
	"testing"
)

func TestReverseSeq(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{5}, []int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseSeq(tt.args.n); reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}
