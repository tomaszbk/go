//go:build goexperiment.arenas

package main

import "arena"

func main() {
	a := arena.NewArena()
	x := arena.New[[200]byte](a)
	x[0] = 9
	a.Free()
	// Use after free.
	//
	// ASAN should detect this deterministically as Free
	// should poison the arena memory.
	//
	// MSAN should detect that this access is to freed
	// memory. This may crash with an "accessed freed arena
	// memory" error before MSAN gets a chance, but if MSAN
	// was not enabled there would be a chance that this
	// could fail to crash on its own.
	println(x[0])
}
