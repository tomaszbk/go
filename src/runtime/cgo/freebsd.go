//go:build freebsd

package cgo

import _ "unsafe" // for go:linkname

// Supply environ and __progname, because we don't
// link against the standard FreeBSD crt0.o and the
// libc dynamic library needs them.

//go:linkname _environ environ
//go:linkname _progname __progname

//go:cgo_export_dynamic environ
//go:cgo_export_dynamic __progname

var _environ uintptr
var _progname uintptr
