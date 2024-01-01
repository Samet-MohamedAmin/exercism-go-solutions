package wordy

import "regexp"

func captureGroup(r *regexp.Regexp, sm []string, group string) string {
	if len(sm) != len(r.SubexpNames()) {
		return ""
	}
	for i, g := range r.SubexpNames() {
		if g == group {
			return sm[i]
		}
	}
	return ""
}
