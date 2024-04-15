package ip_validation

import "testing"

func TestFindUniq(t *testing.T) {
	type args struct {
		n string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"12.255.56.1"}, true},
		{"", args{""}, false},
		{"", args{"abc.def.ghi.jkl"}, false},
		{"", args{"123.456.789.0"}, false},
		{"", args{"12.34.56"}, false},
		{"", args{"127.1.1.0"}, true},
		{"", args{"0.0.0.0"}, true},
		{"", args{"0.34.82.53"}, true},
		{"", args{"192.168.1.3003"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is_valid_ip(tt.args.n); got != tt.want {
				t.Errorf("FindUniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
