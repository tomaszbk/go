//go:build amd64 && linux && gc

package unix

import "syscall"

//go:noescape
func gettimeofday(tv *Timeval) (err syscall.Errno)
