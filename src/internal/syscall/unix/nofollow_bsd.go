//go:build dragonfly || freebsd

package unix

import "syscall"

// References:
// - https://man.freebsd.org/cgi/man.cgi?open(2)
// - https://man.dragonflybsd.org/?command=open&section=2
const noFollowErrno = syscall.EMLINK
