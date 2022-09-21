package your_order_please

import (
	"strconv"
	"strings"
)

func Order(sentence string) string {
	if len(sentence) == 0 {
		return ""
	}

	sw := strings.Split(sentence, " ")
	res := make([]string, len(sw))

	for _, word := range sw {
		for _, char := range word {
			ind, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}

			res[ind-1] = word
			continue
		}
	}

	return strings.Join(res, " ")
}
