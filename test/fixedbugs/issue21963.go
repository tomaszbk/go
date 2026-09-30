// run


package main

import (
	"fmt"
	"runtime"
)

//go:noinline
func f(x []int32, y *int8) int32 {
	c := int32(int16(*y))
	runtime.GC()
	return x[0] * c
}

func main() {
	var x = [1]int32{5}
	var y int8 = -1
	if got, want := f(x[:], &y), int32(-5); got != want {
		panic(fmt.Sprintf("wanted %d, got %d", want, got))
	}
}
