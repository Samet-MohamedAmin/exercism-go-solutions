package listops

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// IntList contains list of integers
type IntList []int

// Foldl (reduce) each item into the accumulator from the left using `fn binFunc`
func (l IntList) Foldl(fn binFunc, initial int) (out int) {
	out = initial
	for _, v := range l {
		out = fn(out, v)
	}
	return
}

// Foldr (reduce) each item into the accumulator from the right using `fn binFunc`
func (l IntList) Foldr(fn binFunc, initial int) (out int) {
	out = initial
	for i := l.Length() - 1; i >= 0; i-- {
		out = fn(l[i], out)
	}
	return
}

// Filter return the list of all items for which `predicate(item)` is True
func (l IntList) Filter(fn predFunc) (out IntList) {
	if l.Length() == 0 {
		return l
	}
	for _, v := range l {
		if fn(v) {
			out = append(out, v)
		}
	}
	return
}

// Length return the total number of items within it
func (l IntList) Length() (out int) {
	for range l {
		out++
	}
	return
}

// Map return the list of the results of applying `function(item)` on all items
func (l IntList) Map(fn unaryFunc) IntList {
	for i := range l {
		l[i] = fn(l[i])
	}
	return l
}

// Reverse return a list with all the original items, but in reversed order
func (l IntList) Reverse() IntList {
	for i, j := 0, l.Length()-1; i < l.Length()/2; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
	return l
}

// Append add all items in the second list to the end of the first list
func (l IntList) Append(la IntList) (out IntList) {
	out = make(IntList, l.Length()+la.Length())
	for i, v := range l {
		out[i] = v
	}
	for i, v := range la {
		out[i+l.Length()] = v
	}
	return
}

// Concat combine all items in all lists into one flattened list
func (l IntList) Concat(ls []IntList) IntList {
	for _, la := range ls {
		l = l.Append(la)
	}
	return l
}
