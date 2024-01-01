package perfect

import (
	"reflect"
	"sort"
	"testing"
)

var _ error = ErrOnlyPositive

func TestFactors(t *testing.T) {
	for _, tc := range factorsTestCases {
		actual := factors2(tc.input)
		if len(actual) == 0 && len(tc.expected) == 0 {
			continue
		}
		sort.SliceStable(actual, func(i, j int) bool { return actual[i] < actual[j] })
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Fatalf("FAIL factors(%d)\nactual %v\nexpected %v", tc.input, actual, tc.expected)
		}
	}

}

func TestGivesPositiveRequiredError(t *testing.T) {
	if _, err := Classify(0); err != ErrOnlyPositive {
		t.Fatalf("FAIL GivesPositiveRequiredError Expected error %q but got %q", ErrOnlyPositive, err)
	}
	t.Logf("PASS GivesPositiveRequiredError")
}

func TestClassifiesCorrectly(t *testing.T) {
	for _, c := range classificationTestCases {
		cat, err := Classify(c.input)
		switch {
		case err != nil:
			if c.ok {
				t.Fatalf("FAIL %s\nClassify(%d)\nExpected no error but got error %q", c.description, c.input, err)
			}
		case !c.ok:
			t.Fatalf("FAIL %s\nClassify(%d)\nExpected error but got %q", c.description, c.input, cat)
		case cat != c.expected:
			t.Fatalf("FAIL %s\nClassify(%d)\nExpected %d, got %d", c.description, c.input, c.expected, cat)
		}
		t.Logf("PASS %s", c.description)
	}
}

// Test that the classifications are not equal to each other.
// If they are equal, then the tests will return false positives.
func TestClassificationsNotEqual(t *testing.T) {
	classifications := []struct {
		class Classification
		name  string
	}{
		{ClassificationAbundant, "ClassificationAbundant"},
		{ClassificationDeficient, "ClassificationDeficient"},
		{ClassificationPerfect, "ClassificationPerfect"},
	}

	for i, pair1 := range classifications {
		for j := i + 1; j < len(classifications); j++ {
			pair2 := classifications[j]
			if pair1.class == pair2.class {
				t.Fatalf("%s should not be equal to %s", pair1.name, pair2.name)
			}
		}
	}
}

func BenchmarkFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range classificationTestCases {
			factors2(c.input)
		}
	}
}

// BenchmarkFactors-4   	       9	 159882348 ns/op	    2096 B/op	      36 allocs/op

func BenchmarkClassify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range classificationTestCases {
			Classify(c.input)
		}
	}
}
