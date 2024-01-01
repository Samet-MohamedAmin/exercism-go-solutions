package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram verify if a string is isogram
func IsIsogram(input string) bool {

	input = strings.ToUpper(input)
	for i, char := range input {
		if !unicode.IsLetter(char) {
			continue
		}
		if strings.Contains(input[:i], string(char)) {
			return false
		}
	}
	return true
}
