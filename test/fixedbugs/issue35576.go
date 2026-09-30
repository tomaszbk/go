// run


// Check print/println(f()) is allowed where f() is multi-value.

package main

func f() (int16, float64, string) { return -42, 42.0, "x" }

func main() {
	print(f())
	println(f())
}
