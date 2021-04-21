// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Kind represnets triangle kind
// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// NaT is not a triangle
	NaT Kind = iota
	// Equ is equilateral
	Equ
	// Iso is isosceles
	Iso
	// Sca is scalene
	Sca
)

func isTriangle(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	if a+b < c || a+c < b || b+c < a {
		return false
	}
	return true
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	} else if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}
