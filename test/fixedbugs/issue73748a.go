// run


package main

import (
	"context"
	"io"
	"runtime/trace"
)

type T struct {
	a [16]int
}

//go:noinline
func f(x *T) {
	*x = T{}
}

func main() {
	trace.Start(io.Discard)
	defer func() {
		recover()
		trace.Log(context.Background(), "a", "b")

	}()
	f(nil)
}
