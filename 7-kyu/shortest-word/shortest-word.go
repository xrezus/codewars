package shortest_word

import (
	"sort"
	"strings"
)

func FindShort(s string) int {
	sl := strings.Split(s, " ")
	sort.Slice(sl, func(i, j int) bool {
		return len(sl[i]) < len(sl[j])
	})

	return len(sl[0])
}
