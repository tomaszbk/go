//go:build aix || darwin

package runtime

func nonblockingPipe() (r, w int32, errno int32) {
	r, w, errno = pipe()
	if errno != 0 {
		return -1, -1, errno
	}
	closeonexec(r)
	setNonblock(r)
	closeonexec(w)
	setNonblock(w)
	return r, w, errno
}
