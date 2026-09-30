// run

package main

type T[_ any] struct{}

var m = map[interface{}]int{
	T[struct{ int }]{}: 0,
	T[struct {
		int "x"
	}]{}: 0,
}

func main() {
	if len(m) != 2 {
		panic(len(m))
	}
}
