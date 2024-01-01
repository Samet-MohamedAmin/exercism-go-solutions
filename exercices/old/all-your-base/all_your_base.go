package allyourbase

import "errors"

var (
	errInvalidDigit = errors.New("all digits must satisfy 0 <= d < input base")
	errInputBase    = errors.New("input base must be >= 2")
	errOutputBase   = errors.New("output base must be >= 2")
)

func baseToDecimal(inputBase int, inputDigits []int) (int, error) {
	out := 0
	for i, pow := len(inputDigits)-1, 1; i >= 0; i, pow = i-1, inputBase*pow {
		if inputDigits[i] < 0 || inputDigits[i] >= inputBase {
			return 0, errInvalidDigit
		}
		out += inputDigits[i] * pow
	}
	return out, nil
}

// ConvertToBase converts a number, represented as a sequence of digits in one base, to any other base
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	switch {
	case inputBase < 2:
		return nil, errInputBase
	case outputBase < 2:
		return nil, errOutputBase
	case len(inputDigits) == 0:
		return []int{0}, nil
	}

	sum, err := baseToDecimal(inputBase, inputDigits)
	switch {
	case err != nil:
		return nil, err
	case sum == 0:
		return []int{0}, nil
	}

	out := []int{}
	for ; sum > 0; sum /= outputBase {
		out = append([]int{sum % outputBase}, out...)
	}
	return out, nil
}
