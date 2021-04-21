package brackets

var bracketPairs = map[rune]rune{
	'[': ']',
	'(': ')',
	'{': '}',
}

// Bracket verify that any and all pairs are matched and nested correctly
func Bracket(in string) bool {
	queue := []rune{}
	for _, c := range in {
		for open, end := range bracketPairs {
			if c == open {
				queue = append(queue, c)
				break
			} else if c == end {
				if len(queue) == 0 || bracketPairs[queue[len(queue)-1]] != c {
					return false
				}
				queue = queue[:len(queue)-1]
				break
			}
		}
	}
	if len(queue) > 0 {
		return false
	}
	return true
}
