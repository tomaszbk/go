package a

type Comparator[T any] func(v1, v2 T) int

func CompareInt[T ~int](a, b T) int {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}
