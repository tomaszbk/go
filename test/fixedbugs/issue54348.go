// run

package main

func main() {
	F[T[int]]()
}

func F[X interface{ M() }]() {
	var x X
	x.M()
}

type T[X any] struct{ E }

type E struct{}

func (h E) M() {}
