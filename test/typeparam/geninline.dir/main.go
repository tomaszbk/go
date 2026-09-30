package main

import "./a"

// Testing inlining of functions that refer to instantiated exported and non-exported
// generic types.

func main() {
	a.Test1()
	a.Test2()
	a.Test3()
}
