package a

import "fmt"

func F[T any](a T) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("panic: %v\n", x)
		}
	}()
	panic(a)
}
