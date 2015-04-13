package sort

import (
	"testing"
)

type sortTestObj struct {
	i int
	data interface{}
}

func (s sortTestObj) GetValue() int {
	return s.i
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

func isArraySorted(a []SortableObject, ascending bool) bool {
	for i := 0; i < len(a) - 1; i++ {
		if ascending == true {
			if a[i].GetValue() > a[i + 1].GetValue() {
				return false
			}
		} else {
			if a[i].GetValue() < a[i + 1].GetValue() {
				return false
			}
		}
	}
	return true
}