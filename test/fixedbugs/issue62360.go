// run


package main

import (
	"fmt"
	"math/big"
)

//go:noinline
func f(x uint32) *big.Int {
	return big.NewInt(int64(x))
}
func main() {
	b := f(0xffffffff)
	c := big.NewInt(0xffffffff)
	if b.Cmp(c) != 0 {
		panic(fmt.Sprintf("b:%x c:%x", b, c))
	}
}
