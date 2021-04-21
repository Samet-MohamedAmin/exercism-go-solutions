package grains

import (
	"errors"
	"fmt"
)

// Square returns the value for a square in chessboard
func Square(input int) (uint64, error) {

	if input < 1 || input > 64 {
		errMsg := fmt.Sprintf("INVALID INPUT: Square(%d) expected \"input < 0 || input > 64\", Actual %d", input, input)
		return 0, errors.New(errMsg)
	}

	var result uint64 = 1

	for i := 1; i < input; i++ {
		result *= 2
	}

	return result, nil
}

// Total returns total sum for every chessboard squares
func Total() uint64 {

	var result uint64 = 0

	for i := 1; i <= 64; i++ {
		s, _ := Square(i)
		result += s
	}

	return result
}
