// errorcheck -lang=go1.17


// Issue 10975: Returning an invalid interface would cause
// `internal compiler error: getinarg: not a func`.

package main

type I interface {
	int // ERROR "interface contains embedded non-interface|embedding non-interface type"
}

func New() I {
	return struct{}{}
}
