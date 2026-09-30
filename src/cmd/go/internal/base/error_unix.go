//go:build unix

package base

import (
	"errors"
	"syscall"
)

func IsETXTBSY(err error) bool {
	return errors.Is(err, syscall.ETXTBSY)
}
