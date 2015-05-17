package sort

import (
	"bytes"
	"strconv"
)

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

func getSortedObjectArrayString(s []SortableObject) string {
	var buf bytes.Buffer
	if (len(s) == 0) { return "" }

	for _, x := range s {
		if x == nil {
			buf.WriteString("nil")
		} else {
			buf.WriteString(strconv.Itoa(x.GetValue()))
		}
		buf.WriteString(",")
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}

func getSortedObjectArraySegmentString(s []SortableObject, start int, end int) string {
	var buf bytes.Buffer
	var x SortableObject
	if (len(s) == 0) { return "" }
	if (start == end) { 
		if s[start] == nil { return "nil" }
		return strconv.Itoa(s[start].GetValue())
	}

	for i := start; i <= end; i++ {
		x = s[i]
		if x == nil {
			buf.WriteString("nil")
		} else {
			buf.WriteString(strconv.Itoa(x.GetValue()))
		}
		buf.WriteString(",")
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}