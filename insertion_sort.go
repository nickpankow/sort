package sort

func InsertSortInt(a []int, ascending bool) {
	if ascending {
		insertSortAscendingInt(a)
	} else {
		insertSortDescendingInt(a)
	}
}

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

func InsertSort(a []SortableObject, ascending bool) {
	if ascending {
		insertSortAscending(a)
	} else {
		insertSortDescending(a)
	}
}

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