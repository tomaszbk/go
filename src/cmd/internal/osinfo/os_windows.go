//go:build windows

package osinfo

import (
	"fmt"
	"internal/syscall/windows"
)

// Version returns the OS version name/number.
func Version() (string, error) {
	major, minor, build := windows.Version()
	return fmt.Sprintf("%d.%d.%d", major, minor, build), nil
}
