//go:build plan9

package os

import (
	"internal/strconv"
	"syscall"
)

func executable() (string, error) {
	fn := "/proc/" + strconv.Itoa(Getpid()) + "/text"
	f, err := Open(fn)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return syscall.Fd2path(int(f.Fd()))
}
