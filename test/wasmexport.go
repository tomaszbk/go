// errorcheck


// Verify that misplaced directives are diagnosed.

//go:build wasm

package p

//go:wasmexport F
func F() {} // OK

type S int32

//go:wasmexport M
func (S) M() {} // ERROR "cannot use //go:wasmexport on a method"
