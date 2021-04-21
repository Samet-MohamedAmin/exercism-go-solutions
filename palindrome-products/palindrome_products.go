package palindrome

import (
	"errors"
	"fmt"
	"math"
)

// Product contains product type
type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func getLimits(fmin, fmax int) (int, int) {
	l1, l2 := fmin*fmin, fmax*fmax
	if l1 < l2 {
		return l2, l1
	}
	return l1, l2
}

func updateProduct(p *Product, condition bool, i, j int) {
	if condition {
		p.Product = i * j
		p.Factorizations = [][2]int{{i, j}}
	} else if p.Product == i*j {
		p.Factorizations = append(p.Factorizations, [2]int{i, j})
	}
}

// Products returns products
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = errors.New("fmin > fmax")
		return
	}
	pmin.Product, pmax.Product = getLimits(fmin, fmax)
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			p := i * j
			if isPalindrome(fmt.Sprint(math.Abs(float64(p)))) {
				updateProduct(&pmin, pmin.Product > p, i, j)
				updateProduct(&pmax, pmax.Product < p, i, j)
			}
		}
	}
	if pmin.Factorizations == nil && pmax.Factorizations == nil {
		err = errors.New("no palindromes")
	}
	return
}
