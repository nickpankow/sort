package sort

import (
	"testing"
	"math/rand"
	"time"
)

func TestHeapSortAscending(t *testing.T) {
	const num_tests = 10
	const array_size = 8
	const max_value = 100
	a := make([]SortableObject, array_size)
	rand.Seed(time.Now().Unix())

	for i := 0; i < num_tests; i++ {
		for j := 0; j < array_size; j++ {
			a[j] = sortTestObj{rand.Int() % max_value, nil}
		}

		HeapSort(a, true)
		if !isArraySorted(a, true) {
			t.Log("Started a:", getSortedObjectArrayString(a))
			t.Error("Merge array random order failed - expected ascending.\na:", getSortedObjectArrayString(a))
		}
		t.Log("Sorted Ascending a:", getSortedObjectArrayString(a))
	}
}

func TestHeapSortDescending(t *testing.T) {
	const num_tests = 10
	const array_size = 8
	const max_value = 100
	a := make([]SortableObject, array_size)
	rand.Seed(time.Now().Unix())

	for i := 0; i < num_tests; i++ {
		for j := 0; j < array_size; j++ {
			a[j] = sortTestObj{rand.Int() % max_value, nil}
		}

		HeapSort(a, false)
		if !isArraySorted(a, false) {
			t.Error("Merge array random order failed - expected descending.\na:", getSortedObjectArrayString(a))
			}
	}
}

func TestHeapifyAsc(t *testing.T) {
	var a, expected SortableObjectArray
	a = []SortableObject{
		sortTestObj{38, nil},
		sortTestObj{27, nil},
		sortTestObj{43, nil},
		sortTestObj{3, nil},
		sortTestObj{9, nil},
		sortTestObj{82, nil},
		sortTestObj{10, nil},
	}
	expected = []SortableObject{
		sortTestObj{82, nil},
		sortTestObj{27, nil},
		sortTestObj{43, nil},
		sortTestObj{3, nil},
		sortTestObj{9, nil},
		sortTestObj{38, nil},
		sortTestObj{10, nil},
	}

	t.Log("Start: ", getSortedObjectArrayString(a))
	heapify_asc(a, len(a))
	if !a.isEqual(expected) {
		t.Error("Heap not sifted correctly. a:", getSortedObjectArrayString(a), " expected:", getSortedObjectArrayString(expected))
	}
}
