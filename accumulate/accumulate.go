package accumulate

// Accumulate returns
func Accumulate(given []string, converter func(string) string) (res []string) {
	for _, e := range given {
		res = append(res, converter(e))
	}
	return
}
