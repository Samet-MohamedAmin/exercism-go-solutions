package diamond

import (
	"bytes"
	"errors"
)

const sep = ' '

var (
	errWrongKey = errors.New("keys must be between 'A' and 'Z'")
)

// generateMatrix generates matrix of bytes for upper half and middle line diamond
func generateMatrix(n int) (m [][]byte) {
	// fill with separator
	rows := 2*n + 1
	for i := 0; i < n+1; i++ {
		m = append(m, bytes.Repeat([]byte{sep}, rows))
	}
	// put diamond letters
	for i := 0; i < n+1; i++ {
		c := byte('A' + i)
		m[i][n-i] = c
		m[i][n+i] = c
	}
	return
}

// Gen generates diamong string
func Gen(b byte) (s string, err error) {
	if b < 'A' || b > 'Z' {
		err = errWrongKey
		return
	}

	n := int(b - 'A')
	m := generateMatrix(n)

	// first half
	for i := 0; i < n; i++ {
		s += string(m[i]) + "\n"
	}
	// middle line
	s += string(m[n]) + "\n"
	// last inverted half
	for i := 0; i < n; i++ {
		s += string(m[n-1-i]) + "\n"
	}

	return
}
