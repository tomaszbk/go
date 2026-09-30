// errorcheck

package main

func f9(a int) (i int, f float64) {
	i := 9          // ERROR "redecl|no new"
	f := float64(9) // ERROR "redecl|no new"
	return i, f
}
