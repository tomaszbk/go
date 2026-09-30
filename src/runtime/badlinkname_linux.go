//go:build amd64 || arm64

package runtime

import _ "unsafe"

// As of Go 1.22, the symbols below are found to be pulled via
// linkname in the wild. We provide a push linkname here, to
// keep them accessible with pull linknames.
// This may change in the future. Please do not depend on them
// in new code.

//go:linkname vdsoClockgettimeSym
