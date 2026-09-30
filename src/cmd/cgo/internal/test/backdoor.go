package cgotest

import _ "unsafe"

//go:linkname lockedOSThread runtime.lockedOSThread
//extern runtime_lockedOSThread
func lockedOSThread() bool
