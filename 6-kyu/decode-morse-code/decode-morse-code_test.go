package decode_morse_code

import "testing"

func TestDecodeMorse(t *testing.T) {
	type args struct {
		n string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{".... . -.--   .--- ..- -.. ."}, "HEY JUDE"},
		{"", args{"... --- ..."}, "SOS"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeMorse(tt.args.n); got != tt.want {
				t.Errorf("DecodeMorse() = %v, want %v", got, tt.want)
			}
		})
	}
}
