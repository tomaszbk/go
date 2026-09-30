//go:build wasip1

package osinfo

import (
	"errors"
	"fmt"
)

// Version returns the OS version name/number.
func Version() (string, error) {
	return "", fmt.Errorf("unable to determine OS version: %w", errors.ErrUnsupported)
}
