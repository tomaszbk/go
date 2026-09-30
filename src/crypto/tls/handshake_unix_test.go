//go:build unix

package tls

import (
	"errors"
	"syscall"
)

func init() {
	isConnRefused = func(err error) bool {
		return errors.Is(err, syscall.ECONNREFUSED)
	}
}
