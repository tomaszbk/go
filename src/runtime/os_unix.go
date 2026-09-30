//go:build unix

package runtime

const (
	// These values are the same on all known Unix systems.
	// If we find a discrepancy some day, we can split them out.
	_F_SETFD    = 2
	_FD_CLOEXEC = 1
)

//go:nosplit
func closeonexec(fd int32) {
	fcntl(fd, _F_SETFD, _FD_CLOEXEC)
}
