package first_non_repeating_character

import (
	"strings"
)

func FirstNonRepeating(str string) string {
	for _, c := range str {
		if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
			return string(c)
		}
	}
	return ""
}
