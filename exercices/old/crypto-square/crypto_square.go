package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func normalize(in string) string {
	out := ""
	for _, c := range in {
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			out += strings.ToLower(string(c))
		}
	}
	return out
}

// Encode encode input
func Encode(in string) string {
	in = normalize(in)
	l := int(math.Ceil(math.Sqrt(float64(len(in)))))
	lines := make([]string, l)
	// iterate input string and add chars to appropriate line
	for i, char := range in {
		lines[i%l] += string(char)
	}
	// append white spaces if needed
	for i := l - 1; i > 0; i-- {
		if len(lines[i]) >= len(lines[0]) {
			break
		}
		lines[i] += " "
	}

	return strings.Join(lines, " ")
}
