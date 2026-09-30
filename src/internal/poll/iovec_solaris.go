package poll

import (
	"syscall"
	"unsafe"
)

func newIovecWithBase(base *byte) syscall.Iovec {
	return syscall.Iovec{Base: (*int8)(unsafe.Pointer(base))}
}
