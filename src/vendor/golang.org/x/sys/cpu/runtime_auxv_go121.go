//go:build go1.21

package cpu

import (
	_ "unsafe" // for linkname
)

//go:linkname runtime_getAuxv runtime.getAuxv
func runtime_getAuxv() []uintptr

func init() {
	getAuxvFn = runtime_getAuxv
}
