package strand

import (
	"strings"
)

var RNAtoDNA = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var result strings.Builder
	for _, dnaRune := range dna {
		if dnaRune, ok := RNAtoDNA[dnaRune]; ok {
			result.WriteRune(dnaRune)
		}
	}

	return result.String()
}
