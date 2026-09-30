//go:build !windows

package os_test

import (
	"errors"
)

// isOSSymlinkUnsupportedError returns true when err is an error
// returned by os.Symlink when symlinks are unsupported by OS.
func isOSSymlinkUnsupportedError(err error) bool {
	return errors.Is(err, errors.ErrUnsupported)
}
