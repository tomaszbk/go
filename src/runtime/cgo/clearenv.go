//go:build linux

package cgo

import _ "unsafe" // for go:linkname

//go:cgo_import_static x_cgo_clearenv
//go:linkname x_cgo_clearenv x_cgo_clearenv
//go:linkname _cgo_clearenv runtime._cgo_clearenv
var x_cgo_clearenv byte
var _cgo_clearenv = &x_cgo_clearenv
