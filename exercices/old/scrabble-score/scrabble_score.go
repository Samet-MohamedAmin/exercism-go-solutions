package scrabble

import (
	"errors"
	"strings"
)

//SingleScore returns the score of a single character
func SingleScore(char rune) int {

	if char < 'A' || char > 'Z' {
		errors.New("wrong character")
		return -1
	}

	switch char {
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 1
	}
}

//Score returns the score of a string
func Score(input string) int {
	result := 0

	for _, char := range strings.ToUpper(input) {
		result += SingleScore(char)
	}

	return result
}
