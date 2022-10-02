package shortest_word

import "testing"

func TestFindShort(t *testing.T) {
	type args struct {
		n string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"bitcoin take over the world maybe who knows perhaps"}, 3},
		{"", args{"turns out random test cases are easier than writing out basic ones"}, 3},
		{"", args{"Let's travel abroad shall we"}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindShort(tt.args.n); got != tt.want {
				t.Errorf("FindShort() = %v, want %v", got, tt.want)
			}
		})
	}
}
