package sort

import (
	"testing"
	// "fmt"
)

func TestInsertSortInt(t *testing.T) {
	a := []int{0,1,2,3,4,5}
	InsertSortInt(a, true)
	if !isArraySortedInt(a, true) {
		t.Error("Array did not sort (asc) correctly.", a)
	}

	a = []int{5,4,3,2,1}
	InsertSortInt(a, true)
	if !isArraySortedInt(a, true) {
		t.Error("Array did not sort (asc) correctly.", a)
	}

	a = []int{6,2,7,5,10,52,-5,9,3,17,1}
	InsertSortInt(a, true)
	if !isArraySortedInt(a, true) {
		t.Error("Array did not sort (asc) correctly.", a)
	}

	a = []int{0,1,2,3,4,5}
	InsertSortInt(a, false)
	if !isArraySortedInt(a, false) {
		t.Error("Array did not sort (desc) correctly.", a)
	}

	a = []int{5,4,3,2,1}
	InsertSortInt(a, false)
	if !isArraySortedInt(a, false) {
		t.Error("Array did not sort (desc) correctly.", a)
	}

	a = []int{6,2,7,5,10,52,-5,9,3,17,1}
	InsertSortInt(a, false)
	if !isArraySortedInt(a, false) {
		t.Error("Array did not sort (desc) correctly.", a)
	}
}

func TestisArraySortedIntAscendingInt(t *testing.T){
	a := []int{0,1,2,3,4,5}
	if !isArraySortedInt(a, true) {
		t.Error("Array is ascending, but method returned that it is not.")
	}
	a = []int{1,0,2,3,4,5}
	if isArraySortedInt(a, true) {
		t.Error("Array is not ascending, but method returned that it is.")
	}
	a = []int{}
	if !isArraySortedInt(a, true) {
		t.Error("Array is ascending, but method returned that it is not.")
	}
}

func TestisArraySortedIntDescendingInt(t *testing.T){
	a := []int{5,4,3,2,1}
	if !isArraySortedInt(a, false) {
		t.Error("Array is descending, but method returned that it is not.", a)
	}
	a = []int{5,3,4,2,1}
	if isArraySortedInt(a, false) {
		t.Error("Array is not descending, but method returned that it is.", a)
	}
	a = []int{}
	if !isArraySortedInt(a, false) {
		t.Error("Array is descending, but method returned that it is not.", a)
	}
}

func isArraySortedInt(a []int, ascending bool) bool {
	for i := 0; i < len(a) - 1; i++ {
		if ascending == true {
			if a[i] > a[i + 1] {
				return false
			}
		} else {
			if a[i] < a[i + 1] {
				return false
			}
		}
	}
	return true
}
