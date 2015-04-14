package sort

// import "fmt"

func SelectionSort(a []SortableObject, ascending bool) {
	if ascending {
		selectionSortAscending(a)
	} else {
		selectionSortDescending(a)
	}
}

func selectionSortAscending(a []SortableObject) {
	// fmt.Println("===")
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
		// fmt.Println(a)
	}
	// fmt.Println("===")
}

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