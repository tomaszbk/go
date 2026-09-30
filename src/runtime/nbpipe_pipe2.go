//go:build dragonfly || freebsd || linux || netbsd || openbsd || solaris

package runtime

func nonblockingPipe() (r, w int32, errno int32) {
	return pipe2(_O_NONBLOCK | _O_CLOEXEC)
}
