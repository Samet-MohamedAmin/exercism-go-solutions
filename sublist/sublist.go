package sublist

// Relation represents relation between 2 integer arrays
type Relation string

// Sublist returns the relation between 2 given integer arrays
func Sublist(a1 []int, a2 []int) Relation {
	switch {
	case len(a1) == 0 && len(a2) == 0:
		return "equal"
	case len(a1) == 0:
		return "sublist"
	case len(a2) == 0:
		return "superlist"
	}

	var minA, maxA []int
	if len(a1) < len(a2) {
		minA, maxA = a1, a2
	} else {
		minA, maxA = a2, a1

	}

	unequal := true
	for i := 0; i < len(maxA); i++ {
		if maxA[i] == minA[0] {
			unequal = false
			if len(maxA)-i < len(minA) {
				return "unequal"
			}
			var j int
			for j = 0; j < len(minA); j++ {
				if minA[j] != maxA[j+i] {
					unequal = true
					break
				}
			}
			if j >= len(minA) {
				break
			}
		}
	}
	if unequal {
		return "unequal"
	}
	if len(a1) == len(a2) {
		return "equal"
	}
	if len(a1) < len(a2) {
		return "sublist"
	}
	return "superlist"
}
