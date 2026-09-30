//go:build darwin || dragonfly || freebsd || (js && wasm) || netbsd || openbsd || wasip1

package os

import "syscall"

func hostname() (name string, err error) {
	name, err = syscall.Sysctl("kern.hostname")
	if err != nil {
		return "", NewSyscallError("sysctl kern.hostname", err)
	}
	return name, nil
}
