// run

package main

var (
	_ = d
	_ = f("_", c, b)
	a = f("a")
	b = f("b")
	c = f("c")
	d = f("d")
)

func f(s string, rest ...int) int {
	print(s)
	return 0
}

func main() {
	println()
}
