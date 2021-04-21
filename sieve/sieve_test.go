package sieve

import (
	"reflect"
	"testing"
)

func TestFill(t *testing.T) {
	for _, tc := range testCasesFill {
		ns := generateNumbers(tc.limit)
		if len(ns) != 0 || len(tc.expected) != 0 {
			if !reflect.DeepEqual(ns, tc.expected) {
				t.Fatalf("FAIL: \nfill(%d) => numbers =  %v\nExpected  %v",
					tc.limit, ns, tc.expected)
			}
		}
		t.Logf("PASS: %d", tc.limit)
	}
}

func TestFindIndex(t *testing.T) {
	t.Logf("TEST: TestCasesIndexFound")
	for i, tc := range testCasesIndexFound {
		actual, ok := findIndex(tc.value, tc.numbers)
		if !ok {
			t.Fatalf("FAIL: \nindex(%d) not found. It should return a valid index",
				tc.value)
		}
		if actual != tc.expected {
			t.Fatalf("FAIL: \nindex(%d) = %d\nExpected  %d",
				tc.value, tc.numbers, tc.expected)
		}
		t.Logf("PASS: %d", i)
	}

	t.Logf("TEST: TestCasesIndexNotFound")
	for i, tc := range testCasesIndexNotFound {
		found, ok := findIndex(tc.value, tc.numbers)
		if ok {
			t.Fatalf("FAIL: \nindex(%d) = %d should not return a result",
				tc.value, found)
		}
		t.Logf("PASS: %d", i)
	}

}

func TestRemove(t *testing.T) {
	for _, tc := range testCasesRemove {
		_, ns := remove(tc.value, tc.numbers)
		if !reflect.DeepEqual(ns, tc.expected) {
			t.Fatalf("FAIL: %s\nremove(%d) => numbers =  %v\nExpected  %v",
				tc.description, tc.value, ns, tc.expected)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestSieve(t *testing.T) {
	for _, tc := range testCases {
		p := Sieve(tc.limit)
		if len(p) != 0 || len(tc.expected) != 0 {
			if !reflect.DeepEqual(p, tc.expected) {
				t.Fatalf("FAIL: %s\nSieve(%d)\nExpected %v\nActual  %v",
					tc.description, tc.limit, tc.expected, p)
			}
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Sieve(tc.limit)
		}
	}
}
