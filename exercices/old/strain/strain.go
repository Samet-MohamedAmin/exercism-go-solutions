package strain

// Ints represents int array
type Ints []int

// Lists represents array of int arrays
type Lists [][]int

// Strings represents array of strings
type Strings []string

// Keep keeps elements that fit the input predicate and discords others
func (l Ints) Keep(pred func(int) bool) (out Ints) {
	for _, e := range l {
		if pred(e) {
			out = append(out, e)
		}
	}
	return
}

// Discard discards elements that fit the input predicate and keeps others
func (l Ints) Discard(pred func(int) bool) (out Ints) {
	return l.Keep(func(e int) bool { return !pred(e) })
}

// Keep keeps elements that fit the input predicate and discords others
func (l Lists) Keep(pred func([]int) bool) (out Lists) {
	for _, e := range l {
		if pred(e) {
			out = append(out, e)
		}
	}
	return
}

// Keep keeps elements that fit the input predicate and discords others
func (l Strings) Keep(pred func(string) bool) (out Strings) {
	for _, e := range l {
		if pred(e) {
			out = append(out, e)
		}
	}
	return
}
