package sort

import (
	"testing"
	// "fmt"
	"math/rand"
	"time"
)

func TestPartitionAsc(t *testing.T) {
	var a, expected SortableObjectArray
	a = []SortableObject{
		sortTestObj{1,nil},sortTestObj{12,nil},sortTestObj{5,nil},sortTestObj{26,nil},sortTestObj{7,nil},
		sortTestObj{14,nil},sortTestObj{3,nil},sortTestObj{7,nil},sortTestObj{2,nil},
	}

	_ = partition_asc(a, 0, len(a) - 1)

	// fmt.Println(getSortedObjectArrayString(a))
	expected = []SortableObject{
		sortTestObj{1,nil},sortTestObj{2,nil},sortTestObj{5,nil},sortTestObj{7,nil},sortTestObj{3,nil},
		sortTestObj{14,nil},sortTestObj{7,nil},sortTestObj{26,nil},sortTestObj{12,nil},
	}
	if !a.isEqual(expected) {
		t.Error("Partition incorrect.  a:", getSortedObjectArrayString(a), " expected:", expected)
	}
}

//81,68,96,76,64,40,17,44

func TestQuickSortAscHelper(t *testing.T) {
	var a SortableObjectArray
	a = []SortableObject{
		sortTestObj{25,nil},sortTestObj{73,nil},sortTestObj{76,nil},sortTestObj{91,nil},sortTestObj{38,nil},
		sortTestObj{21,nil},sortTestObj{3,nil},sortTestObj{72,nil},sortTestObj{51,nil},sortTestObj{65,nil},
		sortTestObj{9,nil},sortTestObj{89,nil},
	}

	// fmt.Println(getSortedObjectArrayString(a))
	// quick_sort_asc(a, 0, len(a) - 1)
	QuickSort(a, true)
	// fmt.Println(getSortedObjectArrayString(a))
	if !isArraySorted(a, true) {
		t.Error("QuickSort array ascending failed - a:", getSortedObjectArrayString(a))
	}
}

func TestQuickSortAsc(t *testing.T) {
	const num_tests = 3
	const array_size = 15
	const max_value = 100
	var a []SortableObject
	rand.Seed(time.Now().Unix())

	for i := 0; i < num_tests; i++ {
		a = make([]SortableObject, array_size)
		for j := 0; j < array_size; j++ {
			a[j] = sortTestObj{rand.Int() % max_value, nil}
		}

		t.Log("Pre:", getSortedObjectArrayString(a))
		QuickSort(a, true)
		t.Log("Post:", getSortedObjectArrayString(a))
		if !isArraySorted(a, true) {
			t.Error("QuickSort array random order failed - expected ascending.")
		}
		t.Log("----")
	}
}

func TestQuickSortDesc(t *testing.T) {
	const num_tests = 3
	const array_size = 15
	const max_value = 100
	var a []SortableObject
	rand.Seed(time.Now().Unix())

	for i := 0; i < num_tests; i++ {
		a = make([]SortableObject, array_size)
		for j := 0; j < array_size; j++ {
			a[j] = sortTestObj{rand.Int() % max_value, nil}
		}

		t.Log("Pre:", getSortedObjectArrayString(a))
		QuickSort(a, false)
		t.Log("Post:", getSortedObjectArrayString(a))
		if !isArraySorted(a, false) {
			t.Error("QuickSort array random order failed - expected descending.")
		}
		t.Log("----")
	}
}