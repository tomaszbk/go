// -lang=go1.19

// Check Go language version-specific errors.

package p

type Slice []byte
type Array [8]byte

var s Slice
var p = (Array)(s /* ERROR "requires go1.20 or later" */)
