//go:build aix || darwin || dragonfly || freebsd || openbsd || solaris || zos

package unix

var mapper = &mmapper{
	active: make(map[*byte][]byte),
	mmap:   mmap,
	munmap: munmap,
}
