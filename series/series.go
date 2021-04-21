package series

// First returns the first substring of s with length n
// if the input string has more or equal to n letters
func First(n int, s string) (out string, ok bool) {
	if n <= len(s) {
		out, ok = UnsafeFirst(n, s), true
	}
	return
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) (out string) {
	return s[:n]
}

// All returns a list of all substrings of s with length n.
func All(n int, s string) (out []string) {
	if f, ok := First(n, s); ok {
		out = append(out, f)
		out = append(out, All(n, s[1:])...)
	}
	return
}
