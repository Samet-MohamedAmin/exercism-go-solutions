package wordy

import (
	"fmt"
	"regexp"
	"strings"
)

type validator struct {
	prefix   string
	suffix   string
	template string
	regexp   *regexp.Regexp
}

func newValidator(o operators) validator {
	v := validator{
		prefix:   "What is",
		suffix:   "\\?",
		template: `^%s (?P<first>-?\d+)(?P<operators>(?: %s -?\d+)*)%s$`,
		regexp:   nil,
	}
	v.genRegExp(o)
	return v
}

func (v *validator) genRegExp(o operators) {
	opString := fmt.Sprintf("(?:%s)", strings.Join(o.names(), "|"))
	s := fmt.Sprintf(v.template, v.prefix, opString, v.suffix)
	v.regexp = regexp.MustCompile(s)
}

func (v validator) validate(in string) bool {
	return v.regexp.MatchString(in)
}
