//go:build !compiler_bootstrap

package bits

import _ "unsafe"

//go:linkname overflowError runtime.overflowError
var overflowError error

//go:linkname divideError runtime.divideError
var divideError error
