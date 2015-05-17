package sort

import (
	"testing"
	"math/rand"
	"time"
)

func TestCopyArray(t *testing.T) {
	var a,b []SortableObject

	a = []SortableObject{sortTestObj{5,nil},sortTestObj{4,nil},sortTestObj{3,nil},sortTestObj{2,nil},sortTestObj{1,nil}}
	b = []SortableObject{sortTestObj{0,nil},sortTestObj{0,nil},sortTestObj{0,nil},sortTestObj{0,nil},sortTestObj{0,nil}}

	copy_array(a, 0, len(a) - 1, b)

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Error("Copy was not successful.  Items at index ", i, " are not equal. a:", getSortedObjectArrayString(a), " b:", getSortedObjectArrayString(b) )
		}
	}
}

func TestMergeAscSizeOne(t *testing.T) {
	var a,b SortableObjectArray

	a = []SortableObject{sortTestObj{5,nil}}
	b = make([]SortableObject, 1)
	merge_asc(a, 0, (len(a) - 0) / 2 , len(a) - 1, b)
	if !a.isEqual(b) {
		t.Error("Merge array of size one failed. a:", getSortedObjectArrayString(a), " b:", getSortedObjectArrayString(b))
	}

}

func TestMergeAscSizeTwo(t *testing.T) {
	var a,b SortableObjectArray

	a = []SortableObject{sortTestObj{5,nil}, sortTestObj{7,nil}}
	b = make([]SortableObject, 2)
	merge_asc(a, 0, len(a) / 2, len(a) - 1, b)
	if !isArraySorted(b, true) {
		t.Error("Merge array of size two failed. a:", getSortedObjectArrayString(a), " b:", getSortedObjectArrayString(b))
	}

	a = []SortableObject{sortTestObj{7,nil}, sortTestObj{5,nil}}
	b = make([]SortableObject, 2)
	merge_asc(a, 0, len(a) / 2, len(a) - 1, b)
	if !isArraySorted(b, true) {
		t.Error("Merge array of size two failed. a:", getSortedObjectArrayString(a), " b:", getSortedObjectArrayString(b))
	}
}

func TestMergeRandomOrder(t *testing.T) {
	var a,b 	SortableObjectArray
	a = []SortableObject{sortTestObj{1,nil},sortTestObj{7,nil},sortTestObj{3,nil},sortTestObj{5,nil},sortTestObj{9,nil}}
	b = []SortableObject{sortTestObj{0,nil},sortTestObj{0,nil},sortTestObj{0,nil},sortTestObj{0,nil},sortTestObj{0,nil}}

	merge_asc(a, 0, len(a) / 2, len(a) - 1, b)
	if !isArraySorted(b, true) {
		t.Error("Merge array random order failed - expected ascending. a:", getSortedObjectArrayString(a), " b:", getSortedObjectArrayString(b))
	}
}

func TestMergeSortAscending(t *testing.T) {
	const num_tests = 10
	const array_size = 8
	const max_value = 100
	a := make([]SortableObject, array_size)
	b := make([]SortableObject, array_size)
	rand.Seed(time.Now().Unix())

	for i := 0; i < num_tests; i++ {
		for j := 0; j < array_size; j++ {
			a[j] = sortTestObj{rand.Int() % max_value, nil}
			b[j] = sortTestObj{0, nil}
		}

		merge_sort_ascending(a, 0, len(a) - 1, b)
		if !isArraySorted(a, true) {
			t.Error("Merge array random order failed - expected ascending.\na:", getSortedObjectArrayString(a), " \nb:", getSortedObjectArrayString(b))
		}
	}
}

func TestMergeSortDecsending(t *testing.T) {
	const num_tests = 10
	const array_size = 8
	const max_value = 100
	a := make([]SortableObject, array_size)
	b := make([]SortableObject, array_size)
	rand.Seed(time.Now().Unix())

	for i := 0; i < num_tests; i++ {
		for j := 0; j < array_size; j++ {
			a[j] = sortTestObj{rand.Int() % max_value, nil}
			b[j] = sortTestObj{0, nil}
		}

		merge_sort_descending(a, 0, len(a) - 1, b)
		if !isArraySorted(a, false) {
			t.Error("Merge array random order failed - expected ascending.\na:", getSortedObjectArrayString(a), " \nb:", getSortedObjectArrayString(b))
			}
	}
}
