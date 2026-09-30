package main

import "./a"

func main() {
	a.V.Print()
	a.FnPrint(&a.V)
}
