package which_are_in

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	var r []string
	if len(array1) == 0 || len(array2) == 0 {
		return r
	}

	for _, str1 := range array1 {
		for _, str2 := range array2 {

			if strings.ContainsAny(str1, str2) {
				r = append(r, str1)
			}
		}
	}

	sort.Strings(r)

	return r
}
