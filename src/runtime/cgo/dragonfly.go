//go:build dragonfly

package cgo

import _ "unsafe" // for go:linkname

// Supply environ and __progname, because we don't
// link against the standard DragonFly crt0.o and the
// libc dynamic library needs them.

//go:linkname _environ environ
//go:linkname _progname __progname

var _environ uintptr
var _progname uintptr
