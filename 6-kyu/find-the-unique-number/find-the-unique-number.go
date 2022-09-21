package find_the_unique_number

import (
	"sort"
)

func FindUniq(arr []float32) float32 {
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })

	if arr[0] == arr[1] {
		return arr[len(arr)-1]
	}

	return arr[0]
}
