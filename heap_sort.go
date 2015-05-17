package sort

// Sort an array based off the Heap Sort algorithm
func HeapSort(a []SortableObject, ascending bool) {
	if ascending {
		heap_sort_ascending(a)
	} else {
		heap_sort_descending(a)
	}
}

// Sort an array in ascending order using a heap sort
func heap_sort_ascending(a []SortableObject) {
	// Convert to heap
	heapify_asc(a, len(a))

	end := len(a) - 1
	for end >= 0 {
		x := a[0]
		a[0] = a[end]
		a[end] = x

		end--
		sift_down_asc(a, 0, end)
	}
}

// Sort an array in descending order using a heap sort
func heap_sort_descending(a []SortableObject) {
	// Convert to heap
	heapify_desc(a, len(a))

	end := len(a) - 1
	for end >= 0 {
		x := a[0]
		a[0] = a[end]
		a[end] = x

		end--
		sift_down_desc(a, 0, end)
	}
}

// Heapify an array in ascending order
func heapify_asc(a []SortableObject, count int) {
	start := (count - 2) / 2
	for start >= 0 {
		sift_down_asc(a, start, count - 1)
		start--
	}
}

// Heapify an array in descending order
func heapify_desc(a []SortableObject, count int) {
	start := (count - 2) / 2
	for start >= 0 {
		sift_down_desc(a, start, count - 1)
		start--
	}
}

func sift_down_asc(a []SortableObject, start int, end int) {
	root := start
	for root * 2 + 1 <= end {
		left_child := root * 2 + 1
		right_child := left_child + 1
		swap := root

		// if the parent is smaller than the left child, the left child should be swapped
		if a[swap].GetValue() < a[left_child].GetValue() {
			swap = left_child
		}
		// If the right child is larger than the left child and parent, it should be swapped
		if right_child <= end && a[swap].GetValue() < a[right_child].GetValue() {
			swap = right_child
		}
		// If swap is still at the root, the heap is already in order
		if swap == root {
			return
		} else {
			// Swap
			x := a[swap]
			a[swap] = a[root]
			a[root] = x
			root = swap
		}
	}
}

func sift_down_desc(a []SortableObject, start int, end int) {
	root := start
	for root * 2 + 1 <= end {
		left_child := root * 2 + 1
		right_child := left_child + 1
		swap := root

		// if the parent is smaller than the left child, the left child should be swapped
		if a[swap].GetValue() > a[left_child].GetValue() {
			swap = left_child
		}
		// If the right child is larger than the left child and parent, it should be swapped
		if right_child <= end && a[swap].GetValue() > a[right_child].GetValue() {
			swap = right_child
		}
		// If swap is still at the root, the heap is already in order
		if swap == root {
			return
		} else {
			// Swap
			x := a[swap]
			a[swap] = a[root]
			a[root] = x
			root = swap
		}
	}
}