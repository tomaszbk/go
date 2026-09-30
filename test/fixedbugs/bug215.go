// errorcheck


// Used to crash the compiler.
// https://golang.org/issue/158

package main

type A struct {	a A }	// ERROR "recursive|cycle"
func foo()		{ new(A).bar() }
func (a A) bar()	{}
