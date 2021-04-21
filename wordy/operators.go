package wordy

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type opMapping map[string]func(int, int) int

type operators struct {
	mapping  opMapping
	template string
	regexp   *regexp.Regexp
}

func newOperators() operators {
	ops := operators{
		mapping: opMapping{
			"plus":          func(a int, b int) int { return a + b },
			"minus":         func(a int, b int) int { return a - b },
			"multiplied by": func(a int, b int) int { return a * b },
			"divided by":    func(a int, b int) int { return a / b },
		},
		template: `^ (?P<operator>%s) (?P<number>-?\d+)(?P<other>.*)`,
		regexp:   nil,
	}
	ops.genRegExp()
	return ops
}

func (o *operators) genRegExp() {
	rs := fmt.Sprintf(o.template, strings.Join(o.names(), "|"))
	o.regexp = regexp.MustCompile(rs)
}

func (o operators) names() (out []string) {
	for k := range o.mapping {
		out = append(out, k)
	}
	return
}

func (o operators) evaluate(first int, in string) int {
	sm := o.regexp.FindStringSubmatch(in)
	if sm == nil {
		return first
	}

	opName := captureGroup(o.regexp, sm, "operator")
	op1 := first
	op2, _ := strconv.Atoi(captureGroup(o.regexp, sm, "number"))
	opFunc := o.mapping[opName]
	other := captureGroup(o.regexp, sm, "other")
	return o.evaluate(opFunc(op1, op2), other)
}
