package main

// // No C code required.
import "C"

import "testplugin/common"

func F() int { return 17 }

var FuncVar = func() {}

func ReadCommonX() int {
	FuncVar()
	return common.X
}

func main() {
	panic("plugin1.main called")
}
