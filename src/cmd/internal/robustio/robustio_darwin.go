package robustio

import (
	"errors"
	"syscall"
)

const errFileNotFound = syscall.ENOENT

// isEphemeralError returns true if err may be resolved by waiting.
func isEphemeralError(err error) bool {
	errno, ok := errors.AsType[syscall.Errno](err)
	return ok && errno == errFileNotFound
}
