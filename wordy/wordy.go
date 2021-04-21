package wordy

import (
	"strconv"
)

// Answer evaluate given question response
func Answer(question string) (answer int, ok bool) {
	o := newOperators()
	v := newValidator(o)

	sm := v.regexp.FindStringSubmatch(question)
	if sm == nil {
		ok = false
		return
	}

	ok = true
	first, _ := strconv.Atoi(captureGroup(v.regexp, sm, "first"))
	answer = o.evaluate(first, captureGroup(v.regexp, sm, "operators"))

	return
}
