//go:build !aix && !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !solaris && !windows && !zos

package nettest

func supportsRawSocket() bool {
	return false
}
