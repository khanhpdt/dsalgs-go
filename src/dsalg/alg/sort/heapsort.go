package sort

import "dsalg/ds/heap"

// HeapSort sorts the given items in ascending order.
// Complexity: O(Nlog(N)), where N is the number of items to sort.
func HeapSort(items []string) []string {
	// build the initial max heap (O(N))
	h := heap.NewMaxHeap(items)

	n := h.Len()
	for i := h.Len() - 1; i >= 1; i-- {
		h.Exchange(0, i) // move the largest item to the end of the heap
		n--              // reduce heap size as the last items are already sorted
		h.Sink(0, n)     // ensure max-heap property
	}

	sorted := h.Keys()
	return sorted
}
