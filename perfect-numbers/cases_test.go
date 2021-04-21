package perfect

// Source: exercism/problem-specifications
// Commit: d7c0227 perfect-numbers: Apply new "input" policy
// Problem Specifications Version: 1.1.0

var factorsTestCases = []struct {
	input    int64
	expected []int64
}{
	{1, []int64{}},
	{2, []int64{1}},
	{3, []int64{1}},
	{4, []int64{1, 2}},
	{6, []int64{1, 2, 3}},
	{12, []int64{1, 2, 3, 4, 6}},
	{25, []int64{1, 5}},
	{100, []int64{1, 2, 4, 5, 10, 20, 25, 50}},
}

var classificationTestCases = []struct {
	description string
	input       int64
	ok          bool
	expected    Classification
}{
	{
		description: "Smallest perfect number is classified correctly",
		input:       6,
		ok:          true,
		expected:    ClassificationPerfect,
	},
	{
		description: "Medium perfect number is classified correctly",
		input:       28,
		ok:          true,
		expected:    ClassificationPerfect,
	},
	{
		description: "Large perfect number is classified correctly",
		input:       33550336,
		ok:          true,
		expected:    ClassificationPerfect,
	},
	{
		description: "Smallest abundant number is classified correctly",
		input:       12,
		ok:          true,
		expected:    ClassificationAbundant,
	},
	{
		description: "Medium abundant number is classified correctly",
		input:       30,
		ok:          true,
		expected:    ClassificationAbundant,
	},
	{
		description: "Large abundant number is classified correctly",
		input:       33550335,
		ok:          true,
		expected:    ClassificationAbundant,
	},
	{
		description: "Smallest prime deficient number is classified correctly",
		input:       2,
		ok:          true,
		expected:    ClassificationDeficient,
	},
	{
		description: "Smallest non-prime deficient number is classified correctly",
		input:       4,
		ok:          true,
		expected:    ClassificationDeficient,
	},
	{
		description: "Medium deficient number is classified correctly",
		input:       32,
		ok:          true,
		expected:    ClassificationDeficient,
	},
	{
		description: "Large deficient number is classified correctly",
		input:       33550337,
		ok:          true,
		expected:    ClassificationDeficient,
	},
	{
		description: "Edge case (no factors other than itself) is classified correctly",
		input:       1,
		ok:          true,
		expected:    ClassificationDeficient,
	},
	{
		description: "Zero is rejected (not a natural number)",
		input:       0,
		ok:          false,
	},
	{
		description: "Negative integer is rejected (not a natural number)",
		input:       -1,
		ok:          false,
	},
}
