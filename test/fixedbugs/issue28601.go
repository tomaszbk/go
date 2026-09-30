// compile


// Failed to compile with gccgo.

package p

import "unsafe"

const w int = int(unsafe.Sizeof(0))

var a [w]byte
