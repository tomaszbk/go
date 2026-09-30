// compile


package main

type T struct {
	a [20]int
}

func f(x [4]int) {
	g(T{}, x)
}

func g(t T, x [4]int)
