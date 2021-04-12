package sort

// MergeSort implements merge sort algorithm.
func MergeSort(items []string) []string {
	if len(items) <= 1 {
		return items
	}

	middlePoint := len(items) / 2
	left, right := items[:middlePoint], items[middlePoint:]
	sortedLeft, sortedRight := MergeSort(left), MergeSort(right)

	merged := merge(sortedLeft, sortedRight)

	return merged
}

func merge(left, right []string) []string {
	out := make([]string, len(left)+len(right))

	// This loop picks the smallest items one-by-one from left and right
	// until one of them runs out of item.
	// Note that left and right are already sorted.
	i, j, k := 0, 0, 0
	for ; i < len(left) && j < len(right); k++ {
		if left[i] <= right[j] {
			out[k] = left[i] // pick item from left
			i++
		} else {
			out[k] = right[j] // pick item from right
			j++
		}
	}

	// This simply copies the remaining items over to the out array.
	var remaining []string
	if i < len(left) {
		remaining = left[i:]
	} else if j < len(right) {
		remaining = right[j:]
	}
	for _, item := range remaining {
		out[k] = item
		k++
	}

	return out
}
