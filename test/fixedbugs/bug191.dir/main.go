package main

import . "./a"
import . "./b"

var _ T
var _ V

func main() {
	if A != 1 || B != 2 {
		panic("wrong vars")
	}
}
