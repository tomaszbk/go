//go:build gccgo && linux && amd64

package unix

import "syscall"

//extern gettimeofday
func realGettimeofday(*Timeval, *byte) int32

func gettimeofday(tv *Timeval) (err syscall.Errno) {
	r := realGettimeofday(tv, nil)
	if r < 0 {
		return syscall.GetErrno()
	}
	return 0
}
