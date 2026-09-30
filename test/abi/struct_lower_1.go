// run

//go:build !wasm


package main

import "fmt"

//go:registerparams
//go:noinline
func passStruct6(a Struct6) Struct6 {
	return a
}

type Struct6 struct {
	Struct1
}

type Struct1 struct {
	A, B, C uint
}

func main() {
	fmt.Println(passStruct6(Struct6{Struct1{1, 2, 3}}))
}
