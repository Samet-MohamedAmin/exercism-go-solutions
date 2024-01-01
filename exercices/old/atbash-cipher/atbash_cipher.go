package atbash

import (
	"strings"
)

// Atbash encrypt given input using the Atbash encryption system
func Atbash(s string) string {
	var words []string
	for i, char := range strings.ToLower(s) {
		var encrypedChar rune
		switch {
		case char >= '0' && char <= '9':
			encrypedChar = char
		case char < 'a' || char > 'z':
			continue
		default:
			encrypedChar = 'z' - (char - 'a')
		}
		if i == 0 || len(words[len(words)-1])%5 == 0 {
			words = append(words, "")
		}
		words[len(words)-1] += string(encrypedChar)
	}
	return strings.Join(words, " ")
}
