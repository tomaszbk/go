package fake

import (
	"errors"
	"syscall"
)

func init() {
	// constants copied from GOROOT/src/internal/syscall/windows/syscall_windows.go
	const (
		ERROR_LOCK_VIOLATION syscall.Errno = 33
	)

	isWindowsErrLockViolation = func(err error) bool {
		return errors.Is(err, ERROR_LOCK_VIOLATION)
	}
}
