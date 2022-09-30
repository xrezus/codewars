package decode_morse_code

import (
	"strings"
)

func DecodeMorse(morseCode string) string {
	morseTable := map[string]string{
		".-":    "A",
		"-...":  "B",
		"-.-.":  "C",
		"-..":   "D",
		".":     "E",
		"..-.":  "F",
		"--.":   "G",
		"....":  "H",
		"..":    "I",
		".---":  "J",
		"-.-":   "K",
		".-..":  "L",
		"--":    "M",
		"-.":    "N",
		"---":   "O",
		".--.":  "P",
		"--.-":  "Q",
		".-.":   "R",
		"...":   "S",
		"-":     "T",
		"..-":   "U",
		"...-":  "V",
		".--":   "W",
		"-..-":  "X",
		"-.--":  "Y",
		"--..":  "Z",
		".----": "1",
		"..---": "2",
		"...--": "3",
		"....-": "4",
		".....": "5",
		"-....": "6",
		"--...": "7",
		"---..": "8",
		"----.": "9",
		"-----": "0",
	}

	var decodeWords []string

	for _, word := range strings.Split(morseCode, "   ") {
		var decodeChars []string

		for _, char := range strings.Split(word, " ") {
			decodeChars = append(decodeChars, morseTable[char])
		}

		decodeWords = append(decodeWords, strings.Join(decodeChars, ""))
	}

	return strings.Join(decodeWords, " ")
}
