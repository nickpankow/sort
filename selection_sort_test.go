package sort

import (
	"testing"
)

func (s sortTestObj) GetValue() int {
	return s.i
}

func TestSelectionSort(t *testing.T) {
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
		SelectionSort(x.data, x.ascending)
		if isArraySorted(x.data, x.ascending) == false {
			t.Error("Array did not sort properly. ", x.data)
		} else {
			t.Log("Sorted: ", x.data)
		}
		t.Log("\n")
	}
}
