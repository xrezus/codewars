package find_the_missing_letter

func FindMissingLetter(chars []rune) rune {
	var desired rune
	for i, char := range chars {
		if char+1 != chars[i+1] {
			desired = char + 1
			break
		}
	}
	return desired
}
