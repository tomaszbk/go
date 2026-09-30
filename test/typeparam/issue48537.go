// compile

package main

func main() {
}

type C interface {
	map[int]string
}

func f[A C]() A {
	return A{
		1: "a",
		2: "b",
	}
}
