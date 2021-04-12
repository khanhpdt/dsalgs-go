// Package heap provides implementations for max-heap and min heap using arrays.
package heap

import "fmt"

// Node is a node in a heap.
type Node struct {
	key string
}

// MaxHeap represents the max-heap data structure.
type MaxHeap struct {
	nodes []*Node
}

// swim re-heapify the maxheap from i up to the root.
func (h *MaxHeap) swim(i int) {
	if i <= 0 || i >= len(h.nodes) {
		return
	}

	parent := (i - 1) / 2
	left, right := parent*2+1, parent*2+2
	largest := left
	if right < len(h.nodes) && h.nodes[left].key < h.nodes[right].key {
		largest = right
	}

	if h.nodes[parent].key < h.nodes[largest].key {
		h.Exchange(parent, largest)
		// continue up to the root
		h.swim(parent)
	}
}

// Insert inserts a new node with the given key to the heap.
func (h *MaxHeap) Insert(k string) {
	h.nodes = append(h.nodes, &Node{key: k})
	h.swim(len(h.nodes) - 1)
}

// Root returns the root item in the heap. It's the largest item.
func (h *MaxHeap) Root() (string, bool) {
	if h == nil || h.nodes == nil {
		return "", false
	}
	return h.nodes[0].key, true
}

// Sink re-heapify the maxheap from i down to n.
func (h *MaxHeap) Sink(i, n int) {
	left, right := i*2+1, i*2+2

	// no child
	if left >= n {
		return
	}

	largest := left
	if right < n && h.nodes[left].key < h.nodes[right].key {
		largest = right
	}

	if h.nodes[i].key < h.nodes[largest].key {
		h.Exchange(i, largest)
		// continue down to the bottom
		h.Sink(largest, n)
	}
}

// NewMaxHeap constructs a new max heap from the given items.
func NewMaxHeap(items []string) *MaxHeap {
	h := new(MaxHeap)

	if len(items) == 0 {
		return h
	}

	// map items to nodes
	for i := 0; i < len(items); i++ {
		h.nodes = append(h.nodes, &Node{key: items[i]})
	}

	nonLeaveNodeStart := (h.Len() - 1) / 2

	// though this for loop seems to take O(NlogN), it actually only takes O(N).
	for i := nonLeaveNodeStart; i >= 0; i-- {
		h.Sink(i, h.Len())
	}

	return h
}

// Len returns the number of items in the heap.
func (h *MaxHeap) Len() int {
	return len(h.nodes)
}

// Exchange swaps the ith and jth items in the heap
func (h *MaxHeap) Exchange(i, j int) {
	if i < 0 || i >= h.Len() || j < 0 || j >= h.Len() {
		panic(fmt.Sprintf("Either i or j is in invalid range: i=%d, j=%d.", i, j))
	}
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

// Keys returns the keys in the nodes.
func (h *MaxHeap) Keys() []string {
	keys := make([]string, h.Len())
	for i, n := range h.nodes {
		keys[i] = n.key
	}
	return keys
}
