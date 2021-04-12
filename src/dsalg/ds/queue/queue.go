// Package queue provides implementation of queue data structure using linked list.
package queue

import "dsalg/ds/ll"

type Queue struct {
	head *ll.Node
	tail *ll.Node
}

// Enqueue adds a new node with the given key to the end of the queue.
func (q *Queue) Enqueue(k string) {
	newNode := &ll.Node{Key: k}

	if q.tail != nil {
		q.tail.Next = newNode
	}
	q.tail = newNode

	if q.head == nil {
		q.head = q.tail
	}
}

// Dequeue removes and returns the item at the beginning of the queue.
// Return NOK if queue is empty.
func (q *Queue) Dequeue() (string, bool) {
	if q.head == nil {
		return "", false
	}
	key := q.head.Key
	q.head = q.head.Next
	return key, true
}

func (q *Queue) String() string {
	l := &ll.LinkedList{Head: q.head}
	return l.String()
}

func (q *Queue) Equal(other *Queue) bool {
	switch {
	case q == nil && other == nil:
		return true
	case q == nil && other != nil:
		return false
	case q != nil && other == nil:
		return false
	default:
		return (&ll.LinkedList{Head: q.head}).Equal(&ll.LinkedList{Head: other.head})
	}
}

func NewQueue(keys ...string) *Queue {
	q := new(Queue)
	for _, k := range keys {
		q.Enqueue(k)
	}
	return q
}
