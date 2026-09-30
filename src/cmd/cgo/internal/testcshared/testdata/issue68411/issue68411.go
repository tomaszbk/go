package main

import "C"

//export exportFuncWithNoParams
func exportFuncWithNoParams() {}

//export exportFuncWithParams
func exportFuncWithParams(a, b int) {}

func main() {}
