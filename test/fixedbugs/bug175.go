// errorcheck


package main

func f() (int, bool) { return 0, true }

func main() {
	x, y := f(), 2 // ERROR "multi|2-valued"
	_, _ = x, y
}
