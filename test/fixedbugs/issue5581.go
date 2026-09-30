// errorcheck

// Used to emit a spurious "invalid recursive type" error.
// See golang.org/issue/5581.

package main

import "fmt"

func NewBar() *Bar { return nil }

func (x *Foo) Method() (int, error) {
	for y := range x.m {
		_ = y.A
	}
	return 0, nil
}

type Foo struct {
	m map[*Bar]int
}

type Bar struct {
	A *Foo
	B chan Blah // ERROR "undefined.*Blah"
}

func main() {
	fmt.Println("Hello, playground")
}
