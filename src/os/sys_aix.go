package os

import "syscall"

// gethostname syscall cannot be used because it also returns the domain.
// Therefore, hostname is retrieve with uname syscall and the Nodename field.

func hostname() (name string, err error) {
	var u syscall.Utsname
	if errno := syscall.Uname(&u); errno != nil {
		return "", NewSyscallError("uname", errno)
	}
	b := make([]byte, len(u.Nodename))
	i := 0
	for ; i < len(u.Nodename); i++ {
		if u.Nodename[i] == 0 {
			break
		}
		b[i] = byte(u.Nodename[i])
	}
	return string(b[:i]), nil
}
