//go:build aix || darwin || openbsd || solaris

package syscall

import _ "unsafe"

// used by internal/poll
//go:linkname writev
