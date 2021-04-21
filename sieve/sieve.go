package sieve

func generateNumbers(n int) (out []int) {
	for i := 2; i <= n; i++ {
		out = append(out, i)
	}
	return
}

func findFromIndex(index, n int, ns []int) (int, bool) {
	var i int
	for i = index; i < len(ns) && ns[i] != n; i++ {
	}
	var ok bool = false
	if i < len(ns) && ns[i] == n {
		ok = true
	}
	return i, ok
}

func findIndex(n int, ns []int) (int, bool) {
	return findFromIndex(0, n, ns)
}

func removeIndex(i int, ns []int) (bool, []int) {
	if i < 0 || i >= len(ns) {
		return false, ns
	}
	ns = append(ns[:i], ns[i+1:]...)
	return true, ns
}

func removeFromIndex(index, n int, ns []int) (bool, []int) {
	i, ok := findFromIndex(index, n, ns)
	if ok {
		return removeIndex(i, ns)
	}
	return false, ns
}

func remove(n int, ns []int) (bool, []int) {
	return removeFromIndex(0, n, ns)
}

// Sieve returns all the primes from 2 up to a given number
func Sieve(n int) []int {
	ns := generateNumbers(n)
	for i := 0; i < len(ns); i++ {
		max := ns[len(ns)-1]
		if ns[i]*2 > max {
			break
		}
		for j := 2; j <= max/ns[i]; j++ {
			_, ns = removeFromIndex(i, j*ns[i], ns)
		}
	}
	return ns
}
