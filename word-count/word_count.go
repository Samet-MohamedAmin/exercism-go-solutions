package wordcount

import (
	"regexp"
	"strings"
)

// Frequency holds the frequency of given word
type Frequency map[string]int

var r = regexp.MustCompile(`(?P<word>\w+(?:'\w+)*)\b`)

// WordCount returns words frequencies
func WordCount(in string) Frequency {
	out := Frequency{}
	for _, w := range r.FindAllString(in, -1) {
		out[strings.ToLower(w)]++
	}
	return out
}
