package sort

func Heapsort(data Interface) {
	heapSort(data, 0, data.Len())
}

func ReverseRange(data Interface, a, b int) {
	reverseRange(data, a, b)
}
