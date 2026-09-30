// Check Go language version-specific errors.

//go:build go1.19

package p

type Slice []byte
type Array [8]byte

var s Slice
var p = (Array)(s /* ok because Go 1.X prior to Go 1.21 ignored the //go:build go1.19 */)
