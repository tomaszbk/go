// run


// Test that pointer bitmaps of types with large scalar tails are
// correctly repeated when unrolled into the heap bitmap.

package main

import "runtime"

const D = 57

type T struct {
	a [D]float64
	b map[string]int
	c [D]float64
}

var ts []T

func main() {
	ts = make([]T, 4)
	for i := range ts {
		ts[i].b = make(map[string]int)
	}
	ts[3].b["abc"] = 42
	runtime.GC()
	if ts[3].b["abc"] != 42 {
		panic("bad field value")
	}
}
