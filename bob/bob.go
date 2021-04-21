package bob

import (
	"regexp"
	"strings"
)

func isQuestion(in string) bool {
	return strings.HasSuffix(in, "?")
}

func hasLower(in string) bool {
	return regexp.MustCompile("[a-z]").MatchString(in)
}

func hasUpper(in string) bool {
	return regexp.MustCompile("[A-Z]").MatchString(in)
}

// Hey returns Bob response.
func Hey(remark string) string {
	remark = regexp.MustCompile("\\s").ReplaceAllString(remark, "")
	if remark == "" {
		return "Fine. Be that way!"
	}
	if hasUpper(remark) && !hasLower(remark) {
		if isQuestion(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}
	if isQuestion(remark) {
		return "Sure."
	}
	return "Whatever."
}
