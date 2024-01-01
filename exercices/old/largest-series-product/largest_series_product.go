package lsproduct

import (
	"errors"
	"strconv"
)

func product(digits string) (int64, error) {
	var out int64 = 1
	for _, c := range digits {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			return -1, errors.New("digits input must only contain digits")
		}
		out *= int64(n)
	}
	return out, nil
}

func maxProduct(pa []int64) int64 {
	var max int64 = -1
	for _, p := range pa {
		if p > max {
			max = p
		}
	}
	return max
}

// LargestSeriesProduct returns largest series product from a given string of numbers
func LargestSeriesProduct(digits string, n int) (int64, error) {
	switch {
	case n == 0:
		return 1, nil
	case n < 0:
		return -1, errors.New("span must be greater than zero")
	case n > len(digits):
		return -1, errors.New("span must be smaller than string length")
	}

	pa := []int64{}
	for i := 0; i <= len(digits)-n; i++ {
		p, err := product(digits[i : i+n])
		if err != nil {
			return -1, err
		}
		pa = append(pa, p)
	}

	return maxProduct(pa), nil
}
