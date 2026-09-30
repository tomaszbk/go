// errorcheck


package main

type T struct {
	x, x int  // ERROR "duplicate|redeclared"
}
