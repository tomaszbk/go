// Linkname builtin symbols (that is not already linknamed,
// e.g. mapaccess1) is not allowed.

package main

import "unsafe"

func main() {
	mapaccess1(nil, nil, nil)
}

//go:linkname mapaccess1 runtime.mapaccess1
func mapaccess1(t, m, k unsafe.Pointer) unsafe.Pointer
