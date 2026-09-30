// Linkname fastrand is allowed _for now_, as it has a
// linknamed definition, for legacy reason.
// NOTE: this may not be allowed in the future. Don't do this!

package main

import _ "unsafe"

//go:linkname fastrand runtime.fastrand
func fastrand() uint32

func main() {
	println(fastrand())
}
