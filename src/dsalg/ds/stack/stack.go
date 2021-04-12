// Package stack provides implementation of stacks using linked list.
package stack

import (
	"dsalg/ds/ll"
)

type Stack struct {
	head *ll.Node
}

// Pop removes and returns the top item (the item added last) from the stack.
// Return NOK if the stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.head != nil {
		key := s.head.Key
		s.head = s.head.Next
		return key, true
	}
	return "", false
}

// Push adds the item to the top of the stack.
func (s *Stack) Push(key string) {
	newNode := &ll.Node{Key: key, Next: s.head}
	s.head = newNode
}

func (s *Stack) String() string {
	l := &ll.LinkedList{Head: s.head}
	return l.String()
}
