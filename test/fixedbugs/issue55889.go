// errorcheck -0 -lang=go1.17

// Prior to Go 1.18, ineffectual //go:linkname directives were treated
// as noops. Ensure that modules that contain these directives (e.g.,
// x/sys prior to go.dev/cl/274573) continue to compile.

package p

import _ "unsafe"

//go:linkname nonexistent nonexistent

//go:linkname constant constant
const constant = 42

//go:linkname typename typename
type typename int
