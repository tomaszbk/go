//go:build ignore

package main

// // No C code required.
import "C"

func FuncInt() int { return 2 }

func FuncRecursive() X { return X{} }

type Y struct {
	X *X
}
type X struct {
	Y Y
}

func main() {}
