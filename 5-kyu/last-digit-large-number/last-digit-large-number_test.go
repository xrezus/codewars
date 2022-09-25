package last_digit_large_number

import "testing"

func TestLastDigit(t *testing.T) {
	type args struct {
		a string
		b string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"4", "1"}, 4},
		{"2", args{"4", "2"}, 6},
		{"3", args{"9", "7"}, 9},
		{"4", args{"10", "10000000000"}, 0},
		{"5", args{"1606938044258990275541962092341162602522202993782792835301376", "2037035976334486086268445688409378161051468393665936250636140449354381299763336706183397376"}, 6},
		{"6", args{"3715290469715693021198967285016729344580685479654510946723", "68819615221552997273737174557165657483427362207517952651"}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastDigit(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("LastDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
