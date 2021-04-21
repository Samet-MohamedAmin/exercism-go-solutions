package perfect

import "errors"

// Classification holds number classification
type Classification byte

const (
	// ClassificationAbundant means that the sum facors is strictly greater than the number
	ClassificationAbundant Classification = iota
	// ClassificationDeficient means that the sum facors is strictly lower than the number
	ClassificationDeficient
	// ClassificationPerfect means that the sum facors is equal to the number
	ClassificationPerfect
)

var (
	// ErrOnlyPositive is error returned in case of given negative number
	ErrOnlyPositive = errors.New("only strictly positive numbers are allowed")
)

func factors(n int64) (ds []int64) {
	if n < 1 {
		return
	}
	for i := int64(1); i < n/2+1; i++ {
		if n%i == 0 {
			ds = append(ds, i)
		}
	}
	return
}

func factors2(n int64) (ds []int64) {
	if n <= 1 {
		return
	}
	ds = []int64{1}
	max := n/2 + 1
	var step int64 = 1
	if n%2 != 0 {
		step = 2
	}
	for i := step + 1; i < max; i += step {
		if n%i == 0 {
			ds = append(ds, i)
			max = n / i
			if i >= max {
				break
			}
			ds = append(ds, max)
		}
	}
	return
}

// Classify given integer input, returns class of integers and error if exists
func Classify(n int64) (class Classification, err error) {
	if n < 1 {
		err = ErrOnlyPositive
		return
	}

	var sum int64 = 0
	for _, f := range factors2(n) {
		sum += f
	}

	switch {
	case sum > n:
		class = ClassificationAbundant
	case sum < n:
		class = ClassificationDeficient
	default:
		class = ClassificationPerfect
	}
	return
}
