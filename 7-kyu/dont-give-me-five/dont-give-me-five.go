package dont_give_me_five

func DontGiveMeFive(start int, end int) int {
	var res []int

	for i := start; i <= end; i++ {
		if i%10 == 5 {
			continue
		}
		res = append(res, i)
	}

	return len(res)
}
