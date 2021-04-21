package summultiples

// SumMultiples returns sum of multiples
func SumMultiples(limit int, divisors ...int) int {
	if len(divisors) == 0 {
		return 0
	}
	div := []int{}
	for _, d := range divisors {
		if d > 0 {
			div = append(div, d)
		}
	}
	res := 0
	for i := 1; i < limit; i++ {
		for _, d := range div {
			if i%d == 0 {
				res += i
				break
			}
		}
	}
	return res
}
