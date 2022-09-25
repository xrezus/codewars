package last_digit_large_number

import (
	"strconv"
)

func LastDigit(n1, n2 string) int {
	var lastDigitCycle = map[int][]int{
		0: []int{0, 0, 0, 0},
		1: []int{1, 1, 1, 1},
		2: []int{2, 4, 8, 6},
		3: []int{9, 7, 1, 3},
		4: []int{6, 4, 6, 4},
		5: []int{5, 5, 5, 5},
		6: []int{6, 6, 6, 6},
		7: []int{7, 9, 3, 1},
		8: []int{8, 4, 2, 6},
		9: []int{1, 9, 1, 9},
	}

	lastDigitOfA, _ := strconv.Atoi(string(n1[len(n1)-1]))
	lastDigitOfB, _ := strconv.Atoi(string(n2[len(n2)-1]))

	reqIndex := lastDigitOfB % 4

	return lastDigitCycle[lastDigitOfA][reqIndex]
}
