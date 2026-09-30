// -lang=go1.19

// Check Go language version-specific errors.

//go:build go1.20

package p

type Slice []byte
type Array [8]byte

var s Slice
var p = (Array)(s /* ok */)
