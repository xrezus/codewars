package supermarket_queue

import "sort"

func QueueTime(customers []int, n int) int {
	t := make([]int, n)

	for i := 0; i < len(customers); i++ {
		idx := indexOf(getMin(t), t)
		t[idx] += customers[i]
	}

	return getMax(t)
}

func getMax(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}

func getMin(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
