//go:build unix

package runtime

import _ "unsafe"

// used in internal/syscall/unix
//go:linkname fcntl
