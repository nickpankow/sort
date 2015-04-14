package sort

// import "fmt"

func MergeSort(a []SortableObject, ascending bool) []SortableObject {
	switch len(a) {
		case 0,1: 
			return a
		case   2: 
			if a[0].GetValue() < a[1].GetValue() == ascending {
				x := a[0]
				a[0] = a[1]
				a[1] = x
			}
			return a
		default:
			left := MergeSort(a[:len(a) / 2], ascending)
			right := MergeSort(a[len(a) / 2: len(a)], ascending)
			return sortedMerge(left, right, ascending)
	}
}

func merge(a []SortableObject, b []SortableObject) []SortableObject {
	return nil
}

func sortSegment(a []SortableObject, start int, end int, ascending bool) {
	if end - start == 2 {
		if a[start].GetValue() > a[end].GetValue() == ascending {
			x := a[start]
			a[start] = a[end]
			a[end] = x
		}
	}
}

func sortedMerge(a []SortableObject, start int, end int, ascending bool) []SortableObject {
	
}