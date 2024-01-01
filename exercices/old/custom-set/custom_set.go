package stringset

import (
	"fmt"
)

// Set holds content
type Set struct {
	values *[]string
}

// New returns new Set
func New() Set {
	return Set{values: &[]string{}}

}

// NewFromSlice returns new set from slice
func NewFromSlice(in []string) Set {
	s := New()
	for _, e := range in {
		s.Add(e)
	}
	return s
}

// String returns Set string value
func (s Set) String() string {
	out := "{"
	for _, e := range *s.values {
		out += fmt.Sprintf("\"%s\", ", e)
	}
	if len(*s.values) > 0 {
		out = out[:len(out)-2]
	}
	out += "}"
	return out

}

// IsEmpty returns true if set is empty. Otherwise it returns false
func (s Set) IsEmpty() bool {
	return len(*s.values) == 0
}

// Has for given string returns true if Set contains that string, otherwise it returns false
func (s Set) Has(in string) bool {
	for _, e := range *s.values {
		if e == in {
			return true
		}
	}
	return false

}

// Subset returns subset for given two Set
func Subset(s1, s2 Set) bool {
	for _, e := range *s1.values {
		if !s2.Has(e) {
			return false
		}
	}
	return true
}

// Disjoint returns disjoint for given two Set
func Disjoint(s1, s2 Set) bool {
	for _, e := range *s1.values {
		if s2.Has(e) {
			return false
		}
	}
	return true
}

// Equal returns true if the two Sets are equal, otherwise it returns false
func Equal(s1, s2 Set) bool {
	if len(*s1.values) != len(*s2.values) {
		return false
	}

	for _, e := range *s1.values {
		if !s2.Has(e) {
			return false
		}
	}
	return true
}

// Add adds given string to the Set
func (s Set) Add(in string) {
	if !s.Has(in) {
		*s.values = append(*s.values, in)
	}
}

// Intersection returns intersection for given two Set
func Intersection(s1, s2 Set) Set {
	out := New()
	for _, e := range *s1.values {
		if s2.Has(e) {
			out.Add(e)
		}
	}
	return out
}

// Difference returns difference for given two Set
func Difference(s1, s2 Set) Set {
	out := New()
	for _, e := range *s1.values {
		if !s2.Has(e) {
			out.Add(e)
		}
	}
	return out
}

// Union returns union for given two Set
func Union(s1, s2 Set) Set {
	out := New()
	for _, s := range []Set{s1, s2} {
		for _, e := range *s.values {
			out.Add(e)
		}
	}
	return out
}
