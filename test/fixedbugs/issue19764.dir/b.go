package main

import "./a"

func main() {
	var x a.I = &a.T{}
	x.M() // call to the wrapper (*T).M
	a.F() // make sure a.F is not dead, which also calls (*T).M inside package a
}
