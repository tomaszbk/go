package syscall

import _ "unsafe"

// used by internal/syscall/unix
//go:linkname unlinkat
//go:linkname openat
//go:linkname fstatat
