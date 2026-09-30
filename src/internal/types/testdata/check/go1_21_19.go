// -lang=go1.21

// Check Go language version-specific errors.

//go:build go1.19

package p

type Slice []byte
type Array [8]byte

var s Slice
var p = (Array)(s /* ok because file versions below go1.21 set the language version to go1.21 */)
