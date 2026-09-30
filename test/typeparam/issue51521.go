// run

package main

import (
	"fmt"
	"strings"
)

type I interface{ M() }

func F[P I](p P) { defer catch(); p.M() }
func G[T any]()  { defer catch(); interface{ M() T }.M(nil) }

func main() {
	F[I](nil)
	G[int]()
}

func catch() {
	err := recover()
	if err, ok := err.(error); ok && strings.Contains(err.Error(), "nil pointer dereference") {
		return
	}
	fmt.Println("FAIL", err)
}
