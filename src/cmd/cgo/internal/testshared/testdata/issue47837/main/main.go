package main

import (
	"testshared/issue47837/a"
)

func main() {
	var vara a.ImplA
	a.TheFuncWithArgA(&vara)
}
