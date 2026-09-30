//go:build linux && go1.24

package unix

import _ "unsafe"

//go:linkname vgetrandom runtime.vgetrandom
//go:noescape
func vgetrandom(p []byte, flags uint32) (ret int, supported bool)
