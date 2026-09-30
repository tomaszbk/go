//go:build !aix && !darwin && !js && !((linux && (amd64 || arm64 || loong64 || riscv64)) || (freebsd && amd64)) && !openbsd && !plan9 && !solaris && !windows

package runtime

import "unsafe"

// mmap calls the mmap system call. It is implemented in assembly.
// We only pass the lower 32 bits of file offset to the
// assembly routine; the higher bits (if required), should be provided
// by the assembly routine as 0.
// The err result is an OS error code such as ENOMEM.
func mmap(addr unsafe.Pointer, n uintptr, prot, flags, fd int32, off uint32) (p unsafe.Pointer, err int)

// munmap calls the munmap system call. It is implemented in assembly.
func munmap(addr unsafe.Pointer, n uintptr)
