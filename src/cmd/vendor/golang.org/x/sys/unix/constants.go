//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos

package unix

const (
	R_OK = 0x4
	W_OK = 0x2
	X_OK = 0x1
)
