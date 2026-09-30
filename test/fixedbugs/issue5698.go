// errorcheck

// Issue 5698: can define a key type with slices.

package main

type Key struct {
	a int16 // the compiler was confused by the padding.
	b []int
}

type Val struct{}

type Map map[Key]Val // ERROR "invalid map key type"
