package armstrong

import (
	"fmt"
	"testing"
)

func TestPow(t *testing.T) {
	testCasesPow := []struct {
		n        int
		m        int
		expected int
	}{
		{1, 3, 1},
		{2, 2, 4},
		{2, 5, 32},
		{10, 2, 100},
		{10, 5, 100000},
	}
	for _, tc := range testCasesPow {
		if actual := pow(tc.n, tc.m); actual != tc.expected {
			t.Fatalf("FAIL: \npow(%d, %d) = %d\n expected = %d",
				tc.n, tc.m, actual, tc.expected)
		}
		t.Logf("PASS: %v", tc)
	}
}

func TestArmstrong(t *testing.T) {
	for _, tc := range testCases[3:] {
		if actual := IsNumber(tc.input); actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %v\nActual: %v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkIsNumber(b *testing.B) {
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("%d", tc.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsNumber(tc.input)
			}
		})
	}
}
