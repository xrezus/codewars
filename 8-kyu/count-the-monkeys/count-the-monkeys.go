package count_the_monkeys

func MonkeyCount(n int) []int {
	var count []int

	for i := 1; i <= n; i++ {
		count = append(count, i)
	}

	return count
}
