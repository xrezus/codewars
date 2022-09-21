package find_the_smallest_integer_in_the_array

import "sort"

func SmallestIntegerFinder(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}
