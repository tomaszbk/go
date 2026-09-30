// errorcheck


// check that compiler doesn't stop reading struct def
// after first unknown type.

// Fixes issue 2110.

package main

type S struct {
	err foo.Bar // ERROR "undefined|expected package"
	Num int
}

func main() {
	s := S{}
	_ = s.Num // no error here please
}
