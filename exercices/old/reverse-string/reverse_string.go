package reverse

// Reverse returns reverse string from given string
func Reverse(input string) string {
	res := ""
	for _, l := range input {
		res = string(l) + res
	}
	return res
}
