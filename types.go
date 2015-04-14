package sort

type SortableObject interface {
	GetValue() int
}

type sortTestObj struct {
	i int
	data interface{}
}

func (s sortTestObj) GetValue() int {
	return s.i
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