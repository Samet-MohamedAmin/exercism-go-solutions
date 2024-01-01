package isbn

import (
	"errors"
	"unicode"
)

func mapValue(c rune) (v rune, err error) {
	switch {
	case c == '-':
		v = -1
	case c == 'X':
		v = 10
	case unicode.IsDigit(c):
		v = c - '0'
	default:
		err = errors.New("cannat have characters other than digits, X and hyphen")
	}
	return
}

// IsValidISBN verifies if a given ISBN is valid
func IsValidISBN(in string) bool {
	var sum, count rune = 0, 0
	for _, c := range in {
		if count >= 10 || c == 'X' && count < 8 {
			return false
		}
		v, err := mapValue(c)
		if err != nil {
			return false
		}
		if v >= 0 {
			sum += v * rune(10-count)
			count++
		}
	}
	if count < 10 {
		return false
	}
	return sum%11 == 0
}
