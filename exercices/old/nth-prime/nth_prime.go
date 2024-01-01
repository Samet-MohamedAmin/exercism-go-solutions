package prime

var primes []int = []int{2}

// isPrime returns true if n is prime and returns false if n is not prime
// n is prime if not divided by prime numbers from prime array
func isPrime(n int) bool {
	if n < primes[0] {
		return false
	}
	for _, p := range primes {
		if n%p == 0 {
			return false
		}
	}
	return true
}

// nextPrime returns the first prime number that is greater than given number n
func nextPrime(n int) int {
	if isPrime(n) {
		return n
	}
	return nextPrime(n + 1)
}

// Nth returns the n'th prime number
func Nth(n int) (int, bool) {
	if n <= 0 {
		return n, false
	}
	for i := len(primes) - 1; i < n; i++ {
		primes = append(primes, nextPrime(primes[i]+1))
	}
	return primes[n-1], true
}
