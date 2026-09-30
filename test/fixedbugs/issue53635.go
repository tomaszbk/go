// run

package main

func main() {
	f[int]()
}

func f[T any]() {
	switch []T(nil) {
	case nil:
	default:
		panic("FAIL")
	}

	switch (func() T)(nil) {
	case nil:
	default:
		panic("FAIL")
	}

	switch (map[int]T)(nil) {
	case nil:
	default:
		panic("FAIL")
	}
}
