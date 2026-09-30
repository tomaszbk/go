//go:build dragonfly || freebsd || linux || netbsd || openbsd || solaris

package runtime

func Pipe() (r, w int32, errno int32) {
	return pipe2(0)
}
