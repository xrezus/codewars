package your_order_please

import "testing"

func TestOrder(t *testing.T) {
	type args struct {
		n string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"is2 Thi1s T4est 3a"}, "Thi1s is2 3a T4est"},
		{"", args{"4of Fo1r pe6ople g3ood th5e the2"}, "Fo1r the2 g3ood 4of th5e pe6ople"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Order(tt.args.n); got != tt.want {
				t.Errorf("Order() = %v, want %v", got, tt.want)
			}
		})
	}
}
