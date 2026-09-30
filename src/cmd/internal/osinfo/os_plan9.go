//go:build plan9

package osinfo

import (
	"os"
)

// Version returns the OS version name/number.
func Version() (string, error) {
	b, err := os.ReadFile("/dev/osversion")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
