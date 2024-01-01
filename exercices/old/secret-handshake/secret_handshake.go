package secret

import "strconv"

var reactions = []string{"wink", "double blink", "close your eyes", "jump"}

func reverse(in []string) {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
}

// Handshake return correct Handshake action
func Handshake(in uint) (out []string) {
	s := strconv.FormatInt(int64(in), 2)
	for i := 0; i < len(s) && i < len(reactions); i++ {
		if s[len(s)-1-i] == '1' {
			out = append(out, reactions[i])
		}
	}

	if len(s) > len(reactions) {
		reverse(out)
	}
	return
}
