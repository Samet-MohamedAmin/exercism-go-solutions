package queenattack

import (
	"errors"
	"regexp"
)

type coords struct {
	x, y int
}

var (
	r                   = regexp.MustCompile(`[a-h][1-8]`)
	errBadPositionFomat = errors.New("bad position format")
	errSamePosition     = errors.New("the two queens have the same position")
)

func getCoords(s string) (coords, error) {
	c := coords{}
	if !r.MatchString(s) {
		return c, errBadPositionFomat
	}

	c.x = int(s[0]) - 'a' + 1
	c.y = int(s[1]) - '1' + 1
	return c, nil
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// CanQueenAttack given two queen position, returns boolean value and error
func CanQueenAttack(w, b string) (bool, error) {
	if w == b {
		return false, errSamePosition
	}
	cs := [2]coords{}
	for i, s := range []string{w, b} {
		c, err := getCoords(s)
		if err != nil {
			return false, err
		}
		cs[i] = c
	}

	ok := false
	if cs[0].x-cs[1].x == 0 || cs[0].y-cs[1].y == 0 || abs(cs[0].x, cs[1].x) == abs(cs[0].y, cs[1].y) {
		ok = true
	}
	return ok, nil
}
