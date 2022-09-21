package where_my_anagrams_at

import (
	"reflect"
	"sort"
	"strings"
)

func Anagrams(word string, words []string) []string {
	var res []string

	for _, w := range words {
		if IsAnagram(word, w) {
			res = append(res, w)
		}
	}

	if len(res) > 0 {
		return res
	}

	return nil
}

func IsAnagram(w1, w2 string) bool {
	sl1 := strings.Split(w1, "")
	sl2 := strings.Split(w2, "")

	sort.Strings(sl1)
	sort.Strings(sl2)

	return reflect.DeepEqual(sl1, sl2)
}
