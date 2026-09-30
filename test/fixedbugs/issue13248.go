// errorcheck


// This program caused an infinite loop with the recursive-descent parser.

package main

func main() {
    foo( // GCCGO_ERROR "undefined name"
} // ERROR "unexpected }|expected operand|missing"
