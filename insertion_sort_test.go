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


func TestInsertSort(t *testing.T) {
	testCases := []struct{
		data 		[]SortableObject
		ascending 	bool
	} {
		{ []SortableObject{sortTestObj{5,nil},sortTestObj{4,nil},sortTestObj{3,nil},sortTestObj{2,nil},sortTestObj{1,nil}}, true },
		{ []SortableObject{sortTestObj{3,nil},sortTestObj{1,nil},sortTestObj{5,nil},sortTestObj{2,nil},sortTestObj{4,nil}}, true },
		{ []SortableObject{sortTestObj{1,nil},sortTestObj{2,nil},sortTestObj{3,nil},sortTestObj{4,nil},sortTestObj{5,nil}}, true },
		{ []SortableObject{sortTestObj{5,nil},sortTestObj{4,nil},sortTestObj{3,nil},sortTestObj{2,nil},sortTestObj{1,nil}}, false },
		{ []SortableObject{sortTestObj{3,nil},sortTestObj{1,nil},sortTestObj{5,nil},sortTestObj{2,nil},sortTestObj{4,nil}}, false },
		{ []SortableObject{sortTestObj{1,nil},sortTestObj{2,nil},sortTestObj{3,nil},sortTestObj{4,nil},sortTestObj{5,nil}}, false },
	}

	for _, x := range testCases { 
		t.Log("Presort: ", x.data)
		InsertSort(x.data, x.ascending)
		if isArraySorted(x.data, x.ascending) == false {
			t.Error("Array did not sort properly. ", x.data)
		} else {
			t.Log("Sorted: ", x.data)
		}
		t.Log("\n")
	}
}

func TestIsArraySorted(t *testing.T){
	testCases := []struct{
		data 		[]SortableObject
		ascending 	bool
		expected	bool
	} {
		{ []SortableObject{sortTestObj{5,nil},sortTestObj{4,nil},sortTestObj{3,nil},sortTestObj{2,nil},sortTestObj{1,nil}}, true, false },
		{ []SortableObject{sortTestObj{3,nil},sortTestObj{1,nil},sortTestObj{5,nil},sortTestObj{2,nil},sortTestObj{4,nil}}, true, false },
		{ []SortableObject{sortTestObj{1,nil},sortTestObj{2,nil},sortTestObj{3,nil},sortTestObj{4,nil},sortTestObj{5,nil}}, true, true },
		{ []SortableObject{sortTestObj{5,nil},sortTestObj{4,nil},sortTestObj{3,nil},sortTestObj{2,nil},sortTestObj{1,nil}}, false, true },
		{ []SortableObject{sortTestObj{3,nil},sortTestObj{1,nil},sortTestObj{5,nil},sortTestObj{2,nil},sortTestObj{4,nil}}, false, false },
		{ []SortableObject{sortTestObj{1,nil},sortTestObj{2,nil},sortTestObj{3,nil},sortTestObj{4,nil},sortTestObj{5,nil}}, false, false },
	}


	for _,x := range testCases {
		if isArraySorted(x.data, x.ascending) != x.expected{
			t.Error("Array does not match expected: ", x)
		}
	}
}
