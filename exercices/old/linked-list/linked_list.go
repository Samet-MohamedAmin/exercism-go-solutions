package linkedlist

import "errors"

// Node holds node value
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

// List holds linked list
type List struct {
	head *Node
	tail *Node
}

// ErrEmptyList is error thrown when given empty list
var ErrEmptyList = errors.New("empty list")

// NewList returns new list
func NewList(args ...interface{}) *List {
	l := &List{}
	for _, v := range args {
		l.PushBack(v)
	}
	return l
}

func (l *List) isEmpty() bool {
	return l.head == nil
}

// Next returns next node
func (e *Node) Next() *Node {
	return e.next
}

// Prev returns previous node
func (e *Node) Prev() *Node {
	return e.prev
}

// PushFront pushs front new node
func (l *List) PushFront(v interface{}) {
	oldHead := l.head
	l.head = &Node{Val: v, next: oldHead}
	if oldHead == nil {
		l.tail = l.head
	}
	if l.head.next != nil {
		l.head.next.prev = l.head
	}
}

// PushBack pushs back new node
func (l *List) PushBack(v interface{}) {
	oldTail := l.tail
	l.tail = &Node{Val: v, prev: oldTail}
	if oldTail == nil {
		l.head = l.tail
	}
	if l.tail.prev != nil {
		l.tail.prev.next = l.tail
	}
}

func (l *List) pop(e *Node) (interface{}, error) {
	if l.isEmpty() {
		return nil, ErrEmptyList
	}

	if e.next != nil {
		e.next.prev = e.prev
	}
	if e.prev != nil {
		e.prev.next = e.next
	}

	if l.head == e {
		l.head = e.next
	}
	if l.tail == e {
		l.tail = e.prev
	}
	return e.Val, nil
}

// PopFront pop front a node
func (l *List) PopFront() (interface{}, error) {
	return l.pop(l.head)
}

// PopBack pop back a node
func (l *List) PopBack() (interface{}, error) {
	return l.pop(l.tail)
}

// Reverse reverses linked list order
func (l *List) Reverse() *List {
	if l.isEmpty() || l.head.next == nil {
		return l
	}
	// reverse head and tail
	l.head, l.tail = l.tail, l.head
	// reverse nodes links
	for e := l.head; e.prev != nil; e = e.next {
		e.prev, e.next = e.next, e.prev
	}
	// reverse tail.prev link
	l.tail.prev = l.tail.next
	// remove head.prev and tail.prev links
	l.head.prev, l.tail.next = nil, nil
	return l
}

// First returns first element
func (l *List) First() *Node {
	return l.head
}

// Last returns last element
func (l *List) Last() *Node {
	return l.tail
}
