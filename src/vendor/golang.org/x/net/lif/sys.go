//go:build solaris

package lif

import "unsafe"

var nativeEndian binaryByteOrder

func init() {
	i := uint32(1)
	b := (*[4]byte)(unsafe.Pointer(&i))
	if b[0] == 1 {
		nativeEndian = littleEndian
	} else {
		nativeEndian = bigEndian
	}
}
