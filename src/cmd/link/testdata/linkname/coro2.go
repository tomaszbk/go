// Linkname corostart is not allowed, as it doesn't have
// a linknamed definition.

package main

import _ "unsafe"

//go:linkname corostart runtime.corostart
func corostart()

func main() {
	corostart()
}
