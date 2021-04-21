package anagram

import (
	"sort"
	"strings"
)

// Detect returns possible anagram for given subject from given string condidates
func Detect(subject string, condidates []string) (out []string) {
	a1 := []rune(strings.ToLower(subject))
	sort.SliceStable(a1, func(i, j int) bool { return a1[i] < a1[j] })
	for _, c := range condidates {
		if len(c) != len(subject) || strings.EqualFold(subject, c) {
			continue
		}
		a2 := []rune(strings.ToLower(c))
		sort.SliceStable(a2, func(i, j int) bool { return a2[i] < a2[j] })
		if string(a1) == string(a2) {
			out = append(out, c)
		}
	}
	return
}
