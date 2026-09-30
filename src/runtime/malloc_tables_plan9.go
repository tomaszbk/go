//go:build plan9

package runtime

import "unsafe"

var (
	mallocScanTable   []func(size uintptr, typ *_type, needzero bool) unsafe.Pointer
	mallocNoScanTable []func(size uintptr, typ *_type, needzero bool) unsafe.Pointer
)
