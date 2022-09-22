package convert_a_string_to_a_number

import (
	"strconv"
)

func StringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}
