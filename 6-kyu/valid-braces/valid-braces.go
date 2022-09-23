package valid_braces

func ValidBraces(str string) bool {
	if len(str)%2 != 0 {
		return false
	}

	var st []rune
	open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, r := range str {
		if closer, ok := open[r]; ok {
			st = append(st, closer)
			continue
		}

		l := len(st) - 1
		if l < 0 || r != st[l] {
			return false
		}

		st = st[:l]
	}

	return len(st) == 0
}
