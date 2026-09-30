// errorcheck


package main

var (
	a, b = f() // ERROR "initialization cycle|depends upon itself|depend upon each other"
	c    = b   // GCCGO_ERROR "depends upon itself|depend upon each other"
)

func f() (int, int) {
	return c, c
}

func main() {}
