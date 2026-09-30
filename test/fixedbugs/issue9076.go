// errorcheck


// Issue 9076: cmd/gc shows computed values in error messages instead of original expression.

package main

import "unsafe"

const Hundred = 100
var _ int32 = 100/unsafe.Sizeof(int(0)) + 1 // ERROR "100 \/ unsafe.Sizeof\(int\(0\)\) \+ 1|incompatible type"
var _ int32 = Hundred/unsafe.Sizeof(int(0)) + 1 // ERROR "Hundred \/ unsafe.Sizeof\(int\(0\)\) \+ 1|incompatible type"
