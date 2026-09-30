// run


package main

func g[P any](...P) P { var zero P; return zero }

var (
	_ int        = g(1, 2)
	_ rune       = g(1, 'a')
	_ float64    = g(1, 'a', 2.3)
	_ float64    = g('a', 2.3)
	_ complex128 = g(2.3, 'a', 1i)
)

func main() {}
