// -lang=go1.17

// The predeclared type comparable is not visible before Go 1.18.

package p

type _ comparable // ERROR "predeclared comparable"
