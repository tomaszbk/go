package x

import "internal/abi"

func Fn0() // defined in assembly

func Fn1() {}

var FnExpr func()

func test() {
	_ = abi.FuncPCABI0(Fn0)           // line 12, no error
	_ = abi.FuncPCABIInternal(Fn0)    // line 13, error
	_ = abi.FuncPCABI0(Fn1)           // line 14, error
	_ = abi.FuncPCABIInternal(Fn1)    // line 15, no error
	_ = abi.FuncPCABI0(FnExpr)        // line 16, error
	_ = abi.FuncPCABIInternal(FnExpr) // line 17, no error
}
