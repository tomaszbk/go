// run

// Missing zero extension when converting a float32
// to a uint64.

package main

import (
	"fmt"
	"math"
)

func Foo(v float32) {
	fmt.Printf("%x\n", uint64(math.Float32bits(v)))
}

func main() {
	Foo(2.0)
}
