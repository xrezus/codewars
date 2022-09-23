package digital_root

import "strconv"

func DigitalRoot(n int) int {
	if len(strconv.Itoa(n)) == 1 {
		return n
	}
	var sum int
	for _, num := range strconv.Itoa(n) {
		chr, err := strconv.Atoi(string(num))
		if err != nil {
			panic(err)
		}
		sum += chr
	}

	return DigitalRoot(sum)
}
