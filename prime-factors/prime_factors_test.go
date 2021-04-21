package prime

import (
	"reflect"
	"sort"
	"testing"
)

func TestAddPrime(t *testing.T) {
	testCases := []struct {
		input        int64
		initialState []int64
		expected     []int64
		description  string
	}{
		{5, []int64{}, []int64{5}, "empty"},
		{7, []int64{}, []int64{7}, "empty"},
		{3, []int64{5}, []int64{3, 5}, "append at the beginning"},
		{2, []int64{3, 7}, []int64{2, 3, 7}, "append at the beginning"},
		{5, []int64{3}, []int64{3, 5}, "append at the end"},
		{7, []int64{2, 3}, []int64{2, 3, 7}, "append at the end"},
		{5, []int64{2, 7}, []int64{2, 5, 7}, "insert between"},
		{11, []int64{2, 3, 5, 7, 13, 19, 23}, []int64{2, 3, 5, 7, 11, 13, 19, 23}, "insert between"},
	}
	for _, tc := range testCases {
		primes = append([]int64{}, tc.initialState...)
		addPrime(tc.input)
		if !reflect.DeepEqual(primes, tc.expected) {
			t.Fatalf("FAIL %s\ninitialState = %v, addPrimes(%d) => primes = %v\nexpected primes = %v",
				tc.description, tc.initialState, tc.input, primes, tc.expected)
		}
		t.Logf("PASS %s", tc.description)
	}
}

func TestFirstPrimeFactor(t *testing.T) {
	testCases := []struct {
		input    int64
		expected bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{9, false},
		{11, true},
		{16, false},
		{19, true},
	}
	for _, tc := range testCases {
		_, actual := firstPrimeFactor(tc.input)
		if actual != tc.expected {
			t.Fatalf("FAIL firstPrimeFactor(%d) = %t; expected %t", tc.input, actual, tc.expected)
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	for _, test := range tests {
		actual := Factors(test.input)
		sort.Slice(actual, ascending(actual))
		sort.Slice(test.expected, ascending(test.expected))
		if !reflect.DeepEqual(actual, test.expected) {
			t.Fatalf("FAIL %s\nFactors(%d) = %#v;\nexpected %#v",
				test.description, test.input,
				actual, test.expected)
		}
		t.Logf("PASS %s", test.description)
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Factors(test.input)
		}
	}
}

func ascending(list []int64) func(int, int) bool {
	return func(i, j int) bool {
		return list[i] < list[j]
	}
}
