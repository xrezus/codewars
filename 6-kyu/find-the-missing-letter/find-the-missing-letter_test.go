package find_the_missing_letter

import "testing"

func TestFindMissingLetter(t *testing.T) {
	type args struct {
		n []rune
	}

	tests := []struct {
		name string
		args args
		want rune
	}{
		{"", args{[]rune{'a', 'b', 'c', 'd', 'f'}}, 'e'},
		{"", args{[]rune{'O', 'Q', 'R', 'S'}}, 'P'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMissingLetter(tt.args.n); got != tt.want {
				t.Errorf("TestFindMissingLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
