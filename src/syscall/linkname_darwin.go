package syscall

import _ "unsafe"

// used by os
//go:linkname closedir
//go:linkname readdir_r

// used by internal/poll
//go:linkname fdopendir

// used by internal/syscall/unix
//go:linkname unlinkat
//go:linkname openat
//go:linkname fstatat

// used by cmd/link
//go:linkname msync
//go:linkname fcntl
