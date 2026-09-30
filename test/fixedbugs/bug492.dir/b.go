package main

import "./a"

func main() {
	defer a.F1(a.F2())
}
