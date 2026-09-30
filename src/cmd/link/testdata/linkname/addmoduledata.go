// Linkname runtime.addmoduledata is not allowed.

package main

import (
	_ "unsafe"
)

//go:linkname addmoduledata runtime.addmoduledata
func addmoduledata()

func main() {
	addmoduledata()
}
