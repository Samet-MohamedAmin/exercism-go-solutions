package prime

import (
	"math"
)

var primes = []int64{}

func addPrime(n int64) {
	var i int = 0
	for ; i < len(primes); i++ {
		if n < primes[i] {
			break
		}
	}
	switch i {
	case 0:
		primes = append([]int64{n}, primes...)
		return
	case len(primes):
		primes = append(primes, n)
		return
	default:
		primes = append(primes[:i], append([]int64{n}, primes[i:]...)...)
	}
}

func firstPrimeFactor(n int64) (int64, bool) {
	if n == 1 {
		return 0, false
	}
	for _, p := range primes {
		if p == n {
			return n, true
		}
	}
	for i := int64(2); i <= int64(math.Round(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return i, false
		}
	}
	addPrime(n)
	return n, true
}

// Factors returns prime factors for given number n
func Factors(n int64) []int64 {
	out := []int64{}
	if n == 1 {
		return out
	}
	f, isPrime := firstPrimeFactor(n)
	if isPrime {
		return append(out, n)
	}
	out = append(out, f)
	return append(out, Factors(n/f)...)
}
