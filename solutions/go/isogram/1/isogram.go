package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	letters := make(map[rune]bool)
	for _, letter := range word {
		letter := unicode.ToLower(letter)
		if letters[letter] {
			return false
		} else if letter != ' ' && letter != '-' {
			letters[letter] = true
		}
	}
	return true
}
