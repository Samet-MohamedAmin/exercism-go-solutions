package tree

import (
	"errors"
	"sort"
)

// Record is a struct containing int fields ID and Parent.
type Record struct {
	ID     int
	Parent int
}

// Node is a struct containing int field ID and []*Node field Children.
type Node struct {
	ID       int
	Children []*Node
}

// Find search for a record by id inside an array of records
func Find(node *Node, id int) *Node {

	if node.ID == id {
		return node
	}

	for _, c := range node.Children {
		found := Find(c, id)
		if found != nil {
			return found
		}
	}

	return nil
}

// FindConcurrent concurrently search for a record by id inside an array of records
func FindConcurrent(node *Node, id int, c chan *Node) {
	select {
	case <-c:
		return
	default:
	}

	if node.ID == id {
		c <- node
		close(c)
		return
	}

	for _, child := range node.Children {
		go FindConcurrent(child, id, c)
	}

}

// Build builds records
func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	rootRecord := records[0]
	if rootRecord.Parent != 0 {
		return nil, errors.New("root node should have no parent")
	}
	if rootRecord.ID != 0 {
		return nil, errors.New("no root node")
	}

	root := &Node{ID: 0}
	for i := 1; i < len(records); i++ {
		r := records[i]
		if Find(root, r.ID) != nil {
			return nil, errors.New("duplicated node")
		}
		if r.ID >= len(records) {
			return nil, errors.New("non-continuous")
		}
		if r.ID == r.Parent {
			return nil, errors.New("cycle directly")
		}
		if r.ID < r.Parent {
			return nil, errors.New("higher id parent of lower id")
		}
		n := &Node{ID: r.ID}
		p := Find(root, r.Parent)
		p.Children = append(p.Children, n)

	}
	return root, nil
}
