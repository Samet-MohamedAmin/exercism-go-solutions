package allyourbase

import (
	"reflect"
	"testing"
)

func TestBaseToDecimal(t *testing.T) {
	testCasesBaseToDecimal := []struct {
		inputBase   int
		inputDigits []int
		expected    int
		fail        bool
	}{
		{2, []int{1}, 1, false},
		{4, []int{1}, 1, false},
		{15, []int{1}, 1, false},
		{2, []int{1, 0, 1}, 5, false},
		{2, []int{1, 1, 1}, 7, false},
		{2, []int{1, 0, 1, 0, 1, 0}, 42, false},
		{3, []int{1, 1, 2, 0}, 42, false},
		{16, []int{2, 10}, 42, false},
		{2, []int{2}, 0, true},
		{4, []int{4}, 0, true},
		{15, []int{15}, 0, true},
		{2, []int{1, 3, 1}, 0, true},
		{2, []int{1, 2, 1}, 0, true},
		{2, []int{2, 0, 1, 0, 1, 0}, 0, true},
		{3, []int{4, 1, 2, 0}, 0, true},
		{16, []int{2, 32}, 0, true},
	}
	for _, c := range testCasesBaseToDecimal {
		actual, err := baseToDecimal(c.inputBase, c.inputDigits)
		if c.fail {
			if err == nil {
				t.Fatalf("FAIL: baseToDecimal(%d, %v) should fail", c.inputBase, c.inputDigits)
			} else {
				continue
			}
		}
		if actual != c.expected {
			t.Fatalf("FAIL:\n\tbaseToDecimal(%d, %v) = %d\n\texpected %d",
				c.inputBase,
				c.inputDigits,
				actual,
				c.expected,
			)
		}
	}
}

func TestConvertToBase(t *testing.T) {
	for _, c := range testCases {
		output, err := ConvertToBase(c.inputBase, c.inputDigits, c.outputBase)
		if c.err != "" {
			if err == nil || c.err != err.Error() {
				t.Fatalf(`FAIL: %s
	Expected error: %s
	Got: %v`, c.description, c.err, err)
			}
		} else if !reflect.DeepEqual(c.expected, output) {
			t.Fatalf(`FAIL: %s
    Input base: %d
    Input digits: %#v
    Output base: %d
    Expected output digits: %#v
    Got: %#v`, c.description, c.inputBase, c.inputDigits, c.outputBase, c.expected, output)
		}
		t.Logf("PASS: %s", c.description)
	}
}
