// run


// issue 5820: register clobber when clearfat and 64 bit arithmetic is interleaved.

package main

func main() {
	array := make([][]int, 2)
	index := uint64(1)
	array[index] = nil
	if array[1] != nil {
		panic("array[1] != nil")
	}
}
