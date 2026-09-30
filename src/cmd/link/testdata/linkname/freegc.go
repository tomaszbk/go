// Linkname runtime.freegc is not allowed.

package main

import (
	_ "unsafe"
)

//go:linkname freegc runtime.freegc
func freegc()

func main() {
	freegc()
}
