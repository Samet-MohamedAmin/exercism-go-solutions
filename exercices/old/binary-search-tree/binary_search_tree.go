package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	right *SearchTreeData
	data  int
}

func Bst(data int) *SearchTreeData {
	return &SearchTreeData{data: data}
}

func (s *SearchTreeData) Insert(data int) {
	if data <= s.data {
		if s.left == nil {
			s.left = Bst(data)
		} else {
			s.left.Insert(data)
		}
	} else {
		if s.right == nil {
			s.right = Bst(data)
		} else {
			s.right.Insert(data)
		}
	}
}

func (s *SearchTreeData) MapString(mapFunc func(int) string) (out []string) {
	if s == nil {
		return
	}
	out = append(out, s.left.MapString(mapFunc)...)
	out = append(out, mapFunc(s.data))
	out = append(out, s.right.MapString(mapFunc)...)
	return
}

func (s *SearchTreeData) MapInt(mapFunc func(int) int) (out []int) {
	if s == nil {
		return
	}
	out = append(out, s.left.MapInt(mapFunc)...)
	out = append(out, mapFunc(s.data))
	out = append(out, s.right.MapInt(mapFunc)...)
	return
}
