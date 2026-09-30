//go:build aix || dragonfly || freebsd || linux || netbsd || openbsd

package unix

// ReadDirent reads directory entries from fd and writes them into buf.
func ReadDirent(fd int, buf []byte) (n int, err error) {
	return Getdents(fd, buf)
}
