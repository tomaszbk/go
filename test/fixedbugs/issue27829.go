// run


// Bad AND/BTR combination rule.

package main

import "fmt"

//go:noinline
func f(x uint64) uint64 {
	return (x >> 48) &^ (uint64(0x4000))
}

func main() {
	bad := false
	if got, want := f(^uint64(0)), uint64(0xbfff); got != want {
		fmt.Printf("got %x, want %x\n", got, want)
		bad = true
	}
	if bad {
		panic("bad")
	}
}
