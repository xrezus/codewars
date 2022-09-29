package build_tower

import (
	"reflect"
	"testing"
)

func TestTowerBuilder(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{1}, []string{"*"}},
		{"", args{2}, []string{" * ", "***"}},
		{"", args{3}, []string{"  *  ", " *** ", "*****"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TowerBuilder(tt.args.n); reflect.DeepEqual(got, tt.want) {
				t.Errorf("TowerBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
