package pete_the_baker

func Cakes(recipe, available map[string]int) int {
	var count, iter int

	for name, value := range recipe {
		v, ok := available[name]
		if !ok {
			return 0
		}

		i := v / value

		if iter == 0 {
			count = i
		}

		if i < count {
			count = i
		}

		iter++
	}

	return count
}
