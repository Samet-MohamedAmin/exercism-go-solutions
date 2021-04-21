package armstrong

import (
	"math"
)

func pow(n, m int) int {
	result := 1
	for i := 0; i < m; i++ {
		result *= n
	}
	return result
}

// IsNumber returns true if n is armstrong number and false if not
func IsNumber(n int) bool {
	count := int(math.Log10(float64(n))) + 1
	sum := 0
	for n2 := n; n2 > 0; n2 /= 10 {
		sum += pow(n2%10, count)
	}
	return n == sum
}
