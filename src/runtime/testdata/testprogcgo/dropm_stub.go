package main

import _ "unsafe" // for go:linkname

// Defined in the runtime package.
//
//go:linkname runtime_getm_for_test runtime.getm
func runtime_getm_for_test() uintptr
