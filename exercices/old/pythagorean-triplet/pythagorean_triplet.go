package pythagorean

import "math"

// Triplet holds pythagore triplets
type Triplet [3]int

// Range returns range
func Range(min, max int) (out []Triplet) {
	for i := min; 2*i*i < max*max; i++ {
		for j := i + 1; j <= max; j++ {
			p1 := i*i + j*j
			if p1 > max*max {
				break
			}
			start := int(math.Sqrt(float64(p1)))
			for k := start; k <= max; k++ {
				if p1 < k*k {
					break
				}
				if p1 == k*k {
					out = append(out, Triplet{i, j, k})
				}
			}
		}
	}
	return
}

// Sum returns sum
func Sum(sum int) (out []Triplet) {
	for i := 1; 2*i*i < (sum-2*i)*(sum-2*i); i++ {
		for j := i + 1; i+2+j+1 <= sum; j++ {
			k := sum - i - j
			p1 := i*i + j*j
			if p1 > k*k {
				break
			}
			if p1 == k*k {
				out = append(out, Triplet{i, j, k})
			}
		}
	}
	return
}
