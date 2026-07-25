package luhn

import (
	"regexp"
	"slices"
	"strings"
)

func Reverse(id string) string {
	byteSlice := []byte(id)
	slices.Reverse(byteSlice)
	return string(byteSlice)
}

func ConvertDigit(digit rune) int {
	num := int(digit - '0')
	num *= 2
	if num > 9 {
		num = num - 9
	}
	return num
}

func Valid(id string) bool {
	match, _ := regexp.Match("^[0-9 ]*$", []byte(id))
	if !match {
		return false
	}
	withoutSpaces := strings.ReplaceAll(id, " ", "")

	// the next step works on every other digit starting
	// from the right and skipping the first position
	// so it is easier to work on a reversed string
	reversed := Reverse(withoutSpaces)
	sum := 0
	for i, char := range reversed {
		if i%2 == 1 {
			convertedChar := ConvertDigit(char)
			sum += convertedChar
		} else {
			sum += int(char - '0')
		}
	}

	return sum%10 == 0 && len(reversed) > 1
}
