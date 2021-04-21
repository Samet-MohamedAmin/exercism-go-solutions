package etl

import "strings"

// Transform returns
func Transform(input map[int][]string) map[string]int {
	res := map[string]int{}
	for k, as := range input {
		for _, s := range as {
			res[strings.ToLower(s)] = k
		}
	}
	return res
}
