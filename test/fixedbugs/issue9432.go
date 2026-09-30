// errorcheck

// gc used to recurse infinitely when dowidth is applied
// to a broken recursive type again.
// See golang.org/issue/9432.
package p

type foo struct { // ERROR "invalid recursive type|cycle"
	bar  foo
	blah foo
}
