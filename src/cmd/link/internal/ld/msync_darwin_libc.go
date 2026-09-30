//go:build darwin

package ld

import _ "unsafe" // for go:linkname

//go:linkname msync syscall.msync
func msync(b []byte, flags int) (err error)
