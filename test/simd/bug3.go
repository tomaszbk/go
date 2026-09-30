// run -gcflags=all=-d=checkptr

//go:build goexperiment.simd && amd64

// Test case for issue #78413.

package main

import (
	"simd/archsimd"
)

//go:noinline
func F() []int32 {
	return []int32{0}
}

func main() {
	archsimd.LoadInt32x8Part(F())
}
