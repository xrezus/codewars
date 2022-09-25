package mexican_wave

import (
	"strings"
)

func Wave(words string) []string {
	if len(words) == 0 {
		return []string{}
	}

	var wave []string

	for ind, char := range words {
		if string(char) == " " {
			continue
		}

		upperChar := []rune(strings.ToUpper(string(words[ind])))

		sr := []rune(words)

		sr[ind] = upperChar[0]

		waveString := string(sr)

		wave = append(wave, waveString)
	}

	return wave
}
