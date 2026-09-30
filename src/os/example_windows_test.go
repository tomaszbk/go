package os_test

import (
	"errors"
	"syscall"
)

// isOSSymlinkUnsupportedError returns true when err is an error
// returned by os.Symlink when symlinks are unsupported by OS.
func isOSSymlinkUnsupportedError(err error) bool {
	return errors.Is(err, errors.ErrUnsupported) ||
		errors.Is(err, syscall.EWINDOWS) ||
		errors.Is(err, syscall.ERROR_PRIVILEGE_NOT_HELD)
}
