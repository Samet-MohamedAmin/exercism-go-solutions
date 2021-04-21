package sieve

// Source: exercism/problem-specifications
// Commit: 8bbb634 sieve: Apply new "input" policy
// Problem Specifications Version: 1.1.0

var testCasesFill = []struct {
	limit    int
	expected []int
}{
	{1, []int{}},
	{2, []int{2}},
	{3, []int{2, 3}},
	{9, []int{2, 3, 4, 5, 6, 7, 8, 9}},
	{15, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
}

var testCasesIndexFound = []struct {
	value    int
	numbers  []int
	expected int
}{
	{2, []int{2, 3, 4}, 0},
	{2, []int{2, 3, 4, 5, 6}, 0},
	{4, []int{2, 3, 4, 5}, 2},
	{8, []int{2, 3, 4, 5, 6, 7, 8, 9}, 6},
	{12, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 10},
	{4, []int{2, 3, 4}, 2},
	{6, []int{2, 3, 4, 5, 6}, 4},
	{8, []int{2, 3, 4, 5, 6, 7, 8}, 6},
}

var testCasesIndexNotFound = []struct {
	value   int
	numbers []int
}{
	{5, []int{2, 3, 4}},
	{7, []int{2, 3, 4, 5, 6}},
	{9, []int{2, 3, 4, 5}},
	{10, []int{2, 3, 4, 5, 6, 7, 8, 9}},
	{17, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
	{-1, []int{2, 3, 4}},
	{-2, []int{2, 3, 4, 5, 6}},
	{-4, []int{2, 3, 4, 5, 6}},
}

var testCasesRemove = []struct {
	value       int
	numbers     []int
	expected    []int
	description string
}{
	{4, []int{2, 3, 4}, []int{2, 3}, "remove last element"},
	{6, []int{2, 3, 4, 5, 6}, []int{2, 3, 4, 5}, "remove last element"},
	{8, []int{2, 3, 4, 5, 6, 7, 8}, []int{2, 3, 4, 5, 6, 7}, "remove last element"},
	{4, []int{2, 3, 4, 5}, []int{2, 3, 5}, "remove between element"},
	{8, []int{2, 3, 4, 5, 6, 7, 8, 9}, []int{2, 3, 4, 5, 6, 7, 9}, "remove between element"},
	{12, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15}, "remove between element"},
	{2, []int{2, 3, 4}, []int{3, 4}, "BONUS: remove first element"},
	{2, []int{2, 3, 4, 5, 6}, []int{3, 4, 5, 6}, "BONUS: remove first element"},
	{2, []int{2, 3, 4, 5, 6, 7, 8}, []int{3, 4, 5, 6, 7, 8}, "BONUS: remove first element"},
}

var testCases = []struct {
	description string
	limit       int
	expected    []int
}{
	{
		"no primes under two",
		1,
		[]int{},
	},
	{
		"find first prime",
		2,
		[]int{2},
	},
	{
		"find primes up to 10",
		10,
		[]int{2, 3, 5, 7},
	},
	{
		"limit is prime",
		13,
		[]int{2, 3, 5, 7, 11, 13},
	},
	{
		"find primes up to 1000",
		1000,
		[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997},
	},
}
