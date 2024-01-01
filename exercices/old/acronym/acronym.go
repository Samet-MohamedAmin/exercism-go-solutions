package acronym

import (
	"strings"
	"unicode"
)

func isSeparator(c rune) bool {
	if unicode.IsLetter(c) || c == '\'' {
		return false
	}
	return true
}

// Abbreviate returns abbreviation form given string
func Abbreviate(s string) (out string) {
	sep := true
	for _, c := range strings.ToUpper(s) {
		if isSeparator(c) {
			sep = true
		} else if sep {
			out += string(c)
			sep = false
		}
	}
	return
}
