package beer

import (
	"errors"
	"fmt"
	"strings"
)

const (
	template = "%s bottle%s of beer on the wall, %s bottle%s of beer.\n%s and %s, %s bottle%s of beer on the wall.\n"
	empty    = "No more"
	max      = 99
)

// Verse return verse
func Verse(n int) (out string, err error) {
	if n < 0 || n > max {
		err = errors.New("invalid number")
	}
	count := fmt.Sprint(n)
	missing := fmt.Sprint(n - 1)
	plural1 := "s"
	plural2 := "s"
	action1 := "Take one down"
	action2 := "pass it around"
	switch n {
	case 0:
		count = empty
		missing = fmt.Sprint(max)
		action1 = "Go to the store"
		action2 = "buy some more"
	case 1:
		missing = empty
		plural1 = ""
		action1 = "Take it down"
	case 2:
		plural2 = ""
	}
	out = fmt.Sprintf(template, count, plural1, strings.ToLower(count), plural1, action1, action2, strings.ToLower(missing), plural2)
	return
}

// Verses return verses
func Verses(a, b int) (out string, err error) {
	if a < b || b < 0 || a > max {
		err = errors.New("invalid numbers")
		return
	}
	for i := a; i >= b; i-- {
		v, _ := Verse(i)
		out += v + "\n"
	}
	return
}

// Song return song
func Song() (out string) {
	out, _ = Verses(max, 0)
	return
}
