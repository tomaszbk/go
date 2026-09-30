// build


// Issue 15926: linker was adding .def to the end of symbols, causing
// a name collision with a method actually named def.

package main

type S struct{}

func (s S) def() {}

var I = S.def

func main() {
    I(S{})
}
