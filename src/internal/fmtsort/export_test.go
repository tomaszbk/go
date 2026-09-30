package fmtsort

import "reflect"

func Compare(a, b reflect.Value) int {
	return compare(a, b)
}
