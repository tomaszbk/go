// This file is compiled and then imported by ddd3.go.

package ddd

func Sum(args ...int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}

