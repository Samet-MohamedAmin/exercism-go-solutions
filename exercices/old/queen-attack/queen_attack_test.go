package queenattack

import "testing"

// Arguments to CanQueenAttack are in algebraic notation.
// See http://en.wikipedia.org/wiki/Algebraic_notation_(chess)

var tests = []struct {
	w, b   string
	attack bool
	ok     bool
}{
	{"b4", "b4", false, false}, // same square
	{"a8", "b9", false, false}, // off board
	{"a0", "b1", false, false},
	{"g3", "i5", false, false},
	{"here", "there", false, false}, // invalid
	{"", "", false, false},

	{"b3", "d7", false, true}, // no attack
	{"a1", "f8", false, true},
	{"b4", "b7", true, true}, // same file
	{"e4", "b4", true, true}, // same rank
	{"a1", "f6", true, true}, // common diagonals
	{"a6", "b7", true, true},
	{"d1", "f3", true, true},
	{"f1", "a6", true, true},
	{"a1", "h8", true, true},
	{"a8", "h1", true, true},
}

func TestCoords(t *testing.T) {
	testCases := []struct {
		s  string
		c  coords
		ok bool
	}{
		{"a1", coords{1, 1}, true},
		{"a2", coords{1, 2}, true},
		{"b1", coords{2, 1}, true},
		{"b2", coords{2, 2}, true},
		{"d8", coords{4, 8}, true},
		{"h8", coords{8, 8}, true},
		{"a0", coords{0, 0}, false},
		{"a9", coords{0, 0}, false},
		{"A1", coords{0, 0}, false},
		{"B1", coords{0, 0}, false},
		{"i1", coords{0, 0}, false},
		{"k1", coords{0, 0}, false},
	}
	for _, tc := range testCases {
		c, err := getCoords(tc.s)
		if !tc.ok {
			if err == nil {
				t.Fatalf("coords(%s)\nexepected to return an error", tc.s)
			}
			continue
		}
		if err != nil {
			t.Fatalf("coords(%s)\nreturned an error \"%v\" but expected error to be nil", tc.s, err)
		}
		if c.x != tc.c.x || c.y != tc.c.y {
			t.Fatalf("coords(%s)\nactual x : %d, y : %d, expected x : %d, y : %d", tc.s, c.x, c.y, tc.c.x, tc.c.y)
		}
	}
}

func TestCanQueenAttack(t *testing.T) {
	for _, test := range tests {
		switch attack, err := CanQueenAttack(test.w, test.b); {
		case err != nil:
			var _ error = err
			if test.ok {
				t.Fatalf("CanQueenAttack(%s, %s) returned error %q.  "+
					"Error not expected.",
					test.w, test.b, err)
			}
		case !test.ok:
			t.Fatalf("CanQueenAttack(%s, %s) = %t, %v.  Expected error.",
				test.w, test.b, attack, err)
		case attack != test.attack:
			t.Fatalf("CanQueenAttack(%s, %s) = %t, want %t.",
				test.w, test.b, attack, test.attack)
		}
	}
}

// Benchmark combined time for all test cases
func BenchmarkCanQueenAttack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			CanQueenAttack(test.w, test.b)
		}
	}
}
