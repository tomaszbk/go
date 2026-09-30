// run

package main

type SomeInterface interface {
	Whatever()
}

func X[T any]() T {
	var m T

	// for this example, this block should never run
	if _, ok := any(m).(SomeInterface); ok {
		var dst SomeInterface
		_, _ = dst.(T)
		return dst.(T)
	}

	return m
}

type holder struct{}

func main() {
	X[holder]()
}
