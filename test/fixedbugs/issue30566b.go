// run


package main

import (
	"bytes"
	"fmt"
)

func main() {
	_, _ = false || g(1), g(2)
	if !bytes.Equal(x, []byte{1, 2}) {
		panic(fmt.Sprintf("wanted [1,2], got %v", x))
	}
}

var x []byte

//go:noinline
func g(b byte) bool {
	x = append(x, b)
	return false
}
