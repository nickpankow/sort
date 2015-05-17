package sort

// Sortable Type
type SortableObject interface {
	GetValue() int  // Method by which sorting algorithms compare objects
}

// Implementation of SortableObject used in testing
type sortTestObj struct {
	i int
	data interface{}
}

// Implement SortableObject Interface for sortTestObj
func (s sortTestObj) GetValue() int {
	return s.i
}

type SortableObjectArray []SortableObject

func (s0 SortableObjectArray) isEqual(s1 []SortableObject) bool {
	for i := 0; i < len(s0); i++ {
		if s0[i] != s1[i] || s0[i].GetValue() != s1[i].GetValue() {
			return false
		}
	}
	return true
}