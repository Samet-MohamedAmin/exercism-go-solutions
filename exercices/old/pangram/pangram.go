package pangram

import (
	"strings"
	"unicode"
)

const max = 'z' - 'a' + 1

// IsPangram checks if given string is pangram. i.e it contains all letters
func IsPangram(in string) bool {
	used := [max]bool{}
	count := 0
	for _, c := range strings.ToLower(in) {
		if !unicode.IsLower(c) {
			continue
		}
		index := c - 'a'
		if !used[index] {
			used[index] = true
			count++
		}
	}
	return count == max
}
