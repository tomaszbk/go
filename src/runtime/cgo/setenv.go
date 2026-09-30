//go:build unix

package cgo

import _ "unsafe" // for go:linkname

//go:cgo_import_static x_cgo_setenv
//go:linkname x_cgo_setenv x_cgo_setenv
//go:linkname _cgo_setenv runtime._cgo_setenv
var x_cgo_setenv byte
var _cgo_setenv = &x_cgo_setenv

//go:cgo_import_static x_cgo_unsetenv
//go:linkname x_cgo_unsetenv x_cgo_unsetenv
//go:linkname _cgo_unsetenv runtime._cgo_unsetenv
var x_cgo_unsetenv byte
var _cgo_unsetenv = &x_cgo_unsetenv
