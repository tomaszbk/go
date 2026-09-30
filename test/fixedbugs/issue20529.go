// errorcheck

//go:build amd64


// Issue 20529: Large stack frames caused compiler panics.
// Only tested on amd64 because the test only makes sense
// on a 64 bit system, and it is platform-agnostic,
// so testing one suffices.

package p

import "runtime"

func f() { // GC_ERROR "stack frame too large"
	x := [][]int{1e9: []int{}}
	runtime.KeepAlive(x)
}
