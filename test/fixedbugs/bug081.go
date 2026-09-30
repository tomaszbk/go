// errorcheck

package main

const x x = 2 // ERROR "loop|type|cycle"

/*
bug081.go:3: first constant must evaluate an expression
Bus error
*/
