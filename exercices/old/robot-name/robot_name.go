package robotname

import (
	"errors"
)

// Robot has n property
type Robot struct {
	n string
}

// Format represent name character code
type Format []struct {
	start rune
	end   rune
}

var (
	usedIndexes = 0
	format      = Format{
		{start: 'A', end: 'Z'},
		{start: 'A', end: 'Z'},
		{start: '0', end: '9'},
		{start: '0', end: '9'},
		{start: '0', end: '9'},
	}
	// first empty name
	nextIndex  = ""
	maxIndexes = 0
)

// InitiateNextIndex return firstEmptyIndex from format
func (f Format) InitiateNextIndex() string {
	index := ""
	for _, c := range f {
		index += string(c.start)
	}
	nextIndex = index
	return index
}

// IntitiateMaxNamesNumber return firstEmptyIndex from format
func (f Format) IntitiateMaxNamesNumber() int {
	max := 1
	for _, c := range f {
		max *= int(c.end - c.start + 1)
	}
	usedIndexes++
	maxIndexes = max
	return max
}

// IncrementNextIndex returns the next empty index
func IncrementNextIndex() string {
	add := 1
	newIndex := ""
	for i := len(format) - 1; i >= 0; i-- {
		c := rune(nextIndex[i])
		if add == 1 {
			f := format[i]
			if c == f.end {
				c = f.start
				add = 1
			} else {
				add = 0
				c++
			}
		}
		newIndex = string(c) + newIndex
	}

	nextIndex = newIndex
	return nextIndex
}

// Reset resets valid name for Robot
func (r *Robot) Reset() error {
	if nextIndex == "" {
		format.IntitiateMaxNamesNumber()
		r.n = format.InitiateNextIndex()
		return nil
	}

	if usedIndexes >= maxIndexes {
		return errors.New("exhausted namespace")
	}

	r.n = IncrementNextIndex()
	usedIndexes++

	return nil
}

// Name create Robot name if not exist and return Robot name
func (r *Robot) Name() (string, error) {
	if r.n != "" {
		return r.n, nil
	}

	if err := r.Reset(); err != nil {
		return "", err
	}
	return r.n, nil

}
