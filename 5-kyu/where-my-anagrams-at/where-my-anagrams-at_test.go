package where_my_anagrams_at

import (
	"reflect"
	"testing"
)

func TestAnagrams(t *testing.T) {
	type args struct {
		word  string
		words []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{"abba", []string{"aabb", "abcd", "bbaa", "dada"}}, []string{"aabb", "bbaa"}},
		{"", args{"laser", []string{"lazing", "lazy", "lacer"}}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Anagrams(tt.args.word, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Anagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
