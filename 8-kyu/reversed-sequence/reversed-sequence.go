package reversed_sequence

func ReverseSeq(n int) []int {
	var res []int

	for i := 0; i < n; n-- {
		res = append(res, n)
	}

	return res
}
