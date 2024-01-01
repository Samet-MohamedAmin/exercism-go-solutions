package darts

import "math"

var scores = map[float64]int{1: 10, 5: 5, 10: 1}

const outside = 0

// Score returns score
func Score(x, y float64) int {
	d := math.Hypot(x, y)
	for k, v := range scores {
		if d <= k {
			return v
		}
	}
	return outside
}
