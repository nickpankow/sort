package sort

// Sort an array of integers using an Insertion sort algorithm.
func InsertSortInt(a []int, ascending bool) {
	if ascending {
		insertSortAscendingInt(a)
	} else {
		insertSortDescendingInt(a)
	}
}

// Sort an array of integers in ascending order using an Insertion sort algorithm.
func insertSortAscendingInt(a []int) {
	for i := 0; i < len(a); i++ {
		// Move search left while it is greater than the item to its left
		for j := i; j > 0; j-- {
			if a[j] > a[j-1] {
				break
			}
			x := a[j]
			a[j] = a[j - 1]
			a[j - 1] = x
		}
	}
}

// Sort an array of integers in descending order using an Insertion sort algorithm.
func insertSortDescendingInt(a []int) {
	for i := 0; i < len(a); i++ {
		// Move search left while it is greater than the item to its left
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				break
			}
			x := a[j]
			a[j] = a[j - 1]
			a[j - 1] = x
		}
	}
}

// Sort an array using an Insertion sort algorithm.
func InsertSort(a []SortableObject, ascending bool) {
	if ascending {
		insertSortAscending(a)
	} else {
		insertSortDescending(a)
	}
}

// Sort an array in ascending order using an Insertion sort algorithm.
func insertSortAscending(a []SortableObject) {
	for i := 0; i < len(a); i++ {
		// Move search left while it is greater than the item to its left
		for j := i; j > 0; j-- {
			if a[j].GetValue() > a[j-1].GetValue() {
				break
			}
			x := a[j]
			a[j] = a[j - 1]
			a[j - 1] = x
		}
	}
}

// Sort an array in descending order using an Insertion sort algorithm.
func insertSortDescending(a []SortableObject) {
	for i := 0; i < len(a); i++ {
		// Move search left while it is greater than the item to its left
		for j := i; j > 0; j-- {
			if a[j].GetValue() < a[j-1].GetValue(){
				break
			}
			x := a[j]
			a[j] = a[j - 1]
			a[j - 1] = x
		}
	}
}