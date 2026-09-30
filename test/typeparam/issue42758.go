// run

package main

func F[T, U int]() interface{} {
	switch interface{}(nil) {
	case int(0), T(0), U(0):
	}

	return map[interface{}]int{int(0): 0, T(0): 0, U(0): 0}
}

func main() {
	F[int, int]()
}
