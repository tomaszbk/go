//go:build openbsd

package cgo

import _ "unsafe" // for go:linkname

// Supply __guard_local because we don't link against the standard
// OpenBSD crt0.o and the libc dynamic library needs it.

//go:linkname _guard_local __guard_local

var _guard_local uintptr

// This is normally marked as hidden and placed in the
// .openbsd.randomdata section.
//
//go:cgo_export_dynamic __guard_local __guard_local
