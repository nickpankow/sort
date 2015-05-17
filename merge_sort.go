package sort

// import "fmt"

var x int

// Sort an array using a merge sort algorithm.
func MergeSort(a []SortableObject, ascending bool) {
	b := make([]SortableObject, len(a))
	if (ascending) {
		merge_sort_ascending(a, 0, len(a) - 1, b)
	} else {
		merge_sort_descending(a, 0, len(a) - 1, b)
	}
}

// Sort a given array (a) in ascending order using a standard merge sort algorithm.  Requires a temporary buffer array (b)
// of equal size.
func merge_sort_ascending(array_to_sort []SortableObject, begin int, end int, buffer_array []SortableObject) {
	if end - begin < 1 { return }

	middle := (end - begin) / 2 
	middle += begin + 1

	merge_sort_ascending(array_to_sort, begin, middle - 1, buffer_array)
	merge_sort_ascending(array_to_sort, middle, end, buffer_array)
	merge_asc(array_to_sort, begin, middle, end, buffer_array)
	copy_array(buffer_array, begin, end, array_to_sort)
}

// Sort a given array (a) in descending order using a standard merge sort algorithm.  Requires a temporary buffer array (b)
// of equal size.
func merge_sort_descending(array_to_sort []SortableObject, begin int, end int, buffer_array []SortableObject) {
	if end - begin < 1 { return }

	middle := (end - begin) / 2 
	middle += begin + 1

	merge_sort_descending(array_to_sort, begin, middle - 1, buffer_array)
	merge_sort_descending(array_to_sort, middle, end, buffer_array)
	merge_desc(array_to_sort, begin, middle, end, buffer_array)
	copy_array(buffer_array, begin, end, array_to_sort)
}

// Merge an array of two parts separated by the middle index into a buffer array in ascending order.
func merge_asc(a []SortableObject, begin int, middle int, end int, b []SortableObject) {
	i0 := begin
	i1 := middle
	for i := begin; i <= end; i++ {
		if (i0 < middle) && (i1 > end || a[i0].GetValue() < a[i1].GetValue()) {
			b[i] = a[i0]
			i0++	
		} else {
			b[i] = a[i1]
			i1++
		}
	}
}

// Merge an array of two parts separated by the middle index into a buffer array in descending order.
func merge_desc(a []SortableObject, begin int, middle int, end int, b []SortableObject) {
	i0 := begin
	i1 := middle
	for i := begin; i <= end; i++ {
		if (i0 < middle) && (i1 > end || a[i0].GetValue() > a[i1].GetValue()) {
			b[i] = a[i0]
			i0++	
		} else {
			b[i] = a[i1]
			i1++
		}
	}
}

// Copy a user-defined segment of an array into the a different array.
func copy_array(b []SortableObject, begin int, end int, a []SortableObject) {
	for i := begin; i <= end; i++ {
		a[i] = b[i]
	}
}
