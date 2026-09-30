// errorcheck

// This code was incorrectly accepted by gccgo.

package main

type N string
type M string

const B N = "B"
const C M = "C"

func main() {
	q := B + C // ERROR "mismatched types|incompatible types"
	println(q)
}
