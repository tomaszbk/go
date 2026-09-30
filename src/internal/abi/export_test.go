package abi

func FuncPCTestFn()

var FuncPCTestFnAddr uintptr // address of FuncPCTestFn, directly retrieved from assembly

//go:noinline
func FuncPCTest() uintptr {
	return FuncPCABI0(FuncPCTestFn)
}
