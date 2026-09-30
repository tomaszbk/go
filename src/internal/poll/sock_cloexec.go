// This file implements accept for platforms that provide a fast path for
// setting SetNonblock and CloseOnExec.

//go:build dragonfly || freebsd || linux || netbsd || openbsd

package poll

import "syscall"

// Wrapper around the accept system call that marks the returned file
// descriptor as nonblocking and close-on-exec.
func accept(s int) (int, syscall.Sockaddr, string, error) {
	ns, sa, err := Accept4Func(s, syscall.SOCK_NONBLOCK|syscall.SOCK_CLOEXEC)
	if err != nil {
		return -1, nil, "accept4", err
	}
	return ns, sa, "", nil
}
