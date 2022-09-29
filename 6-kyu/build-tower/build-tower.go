package build_tower

import (
	"strings"
)

func TowerBuilder(nFloors int) []string {
	var build []string

	for i := 0; i < nFloors; i++ {
		stars := strings.Repeat("*", i)
		blankOffset := strings.Repeat(" ", nFloors-i-1)
		build = append(build, blankOffset+stars+"*"+stars+blankOffset)
	}

	return build
}
