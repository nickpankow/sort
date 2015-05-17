package sort

// Sort an array using a Selection Sort algorithm.
func SelectionSort(a []SortableObject, ascending bool) {
	if ascending {
		selectionSortAscending(a)
	} else {
		selectionSortDescending(a)
	}
}

// Sort an array in ascending order using a Selection Sort algorithm.
func selectionSortAscending(a []SortableObject) {
	for i := 0; i < len(a); i++ {
		lowest := i
		for j := i; j < len(a); j++ {
			if a[j].GetValue() < a[lowest].GetValue() {
				lowest = j
			}
		}
		
		x := a[i]
		a[i] = a[lowest]
		a[lowest] = x
	}
}

// Sort an array in descending order using a Selection Sort algorithm.
func selectionSortDescending(a []SortableObject) {
	for i := 0; i < len(a); i++ {
		lowest := i
		for j := i; j < len(a); j++ {
			if a[j].GetValue() > a[lowest].GetValue() {
				lowest = j
			}
		}
		
		x := a[i]
		a[i] = a[lowest]
		a[lowest] = x
	}
}