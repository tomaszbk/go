// run


package main

import (
	"fmt"
	"unsafe"
)

type S struct {
	p *byte
	a string
	b string
	c int64
	d int64
}

func main() {
	s := &S{p: nil, a: "foo", b: "foo", c: 0, d: 0}
	s.a = ""
	s.b = "bar"
	s.c = 33

	z := (*[2]uintptr)(unsafe.Pointer(&s.a))
	fmt.Printf("%x %x\n", z[0], z[1])
}
