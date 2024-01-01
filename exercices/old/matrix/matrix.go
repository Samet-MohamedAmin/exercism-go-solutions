package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix type
type Matrix [][]int

// New returns new matrix
func New(in string) (Matrix, error) {
	m := Matrix{}
	rows := strings.Split(in, "\n")
	for i, row := range rows {
		r := []int{}
		cols := strings.Split(row, " ")
		for _, c := range cols {
			if c == "" {
				continue
			}
			v, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			r = append(r, v)
		}
		if (i > 0) && (len(r) != len(m[0])) {
			return nil, errors.New("Invalid columns number")
		}
		m = append(m, r)
	}

	return m, nil
}

// Rows return rows
func (m Matrix) Rows() [][]int {
	var rows [][]int = [][]int{}
	for _, r := range m {
		row := []int{}
		for _, c := range r {
			row = append(row, c)
		}
		rows = append(rows, row)
	}
	return rows
}

// Cols returns columns
func (m Matrix) Cols() [][]int {
	var cols [][]int = [][]int{}
	for i := 0; i < len(m[0]); i++ {
		col := []int{}
		for j := 0; j < len(m); j++ {
			col = append(col, m[j][i])
		}
		cols = append(cols, col)
	}
	return cols
}

// Set returns set
func (m *Matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 || r >= len(*m) || (len(*m) > 0 && c >= len((*m)[0])) {
		return false
	}
	(*m)[r][c] = val
	return true
}
