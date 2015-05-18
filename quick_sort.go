package sort

func QuickSort(a []SortableObject, ascending bool) {
	if ascending {
		quick_sort_asc(a, 0, len(a) - 1)
	} else {
		quick_sort_desc(a, 0, len(a) - 1)
	}
}

func quick_sort_asc(a []SortableObject, begin int, end int) {
	pivot := partition_asc(a, begin, end)
	if begin < pivot - 1 {
		quick_sort_asc(a, begin, pivot - 1)
	}
	if pivot < end {
		quick_sort_asc(a, pivot, end)
	}
}

func quick_sort_desc(a []SortableObject, begin int, end int) {
	pivot := partition_desc(a, begin, end)
	if begin < pivot - 1 {
		quick_sort_desc(a, begin, pivot - 1)
	}
	if pivot < end {
		quick_sort_desc(a, pivot, end)
	}
}

func partition_asc(a []SortableObject, begin int, end int) int {
	pivot := a[(end + begin) / 2]
	i := begin
	j := end

	for i <= j {
		for a[i].GetValue() < pivot.GetValue() { i++ }
		for a[j].GetValue() > pivot.GetValue() { j-- }

		if i <= j {
			swap := a[i]
			a[i] = a[j]
			a[j] = swap
			i++
			j--
		}
	}
	return i
}

func partition_desc(a []SortableObject, begin int, end int) int {
	pivot := a[(end + begin) / 2]
	i := begin
	j := end

	for i <= j {
		for a[i].GetValue() > pivot.GetValue() { i++ }
		for a[j].GetValue() < pivot.GetValue() { j-- }

		if i <= j {
			swap := a[i]
			a[i] = a[j]
			a[j] = swap
			i++
			j--
		}
	}
	return i
}