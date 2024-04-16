package pete_the_baker

import "testing"

func TestCakes(t *testing.T) {
	type args struct {
		recipe    map[string]int
		available map[string]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{map[string]int{"flour": 500, "sugar": 200, "eggs": 1}, map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200}}, 2},
		{"2", args{map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100}, map[string]int{"sugar": 500, "flour": 2000, "milk": 2000}}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cakes(tt.args.recipe, tt.args.available); got != tt.want {
				t.Errorf("Cakes() = %v, want %v", got, tt.want)
			}
		})
	}
}
