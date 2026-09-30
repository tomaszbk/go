// errorcheck


// Issue 4452. Used to print many errors, now just one.

package main

func main() {
	_ = [...]int(4) // ERROR "\[\.\.\.\].*outside of array literal|invalid use of \[\.\.\.\] array"
}
