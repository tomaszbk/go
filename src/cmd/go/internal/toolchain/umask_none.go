//go:build !(darwin || freebsd || linux || netbsd || openbsd)

package toolchain

import "io/fs"

func sysWriteBits() fs.FileMode {
	return 0700
}
