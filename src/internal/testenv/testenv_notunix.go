//go:build windows || plan9 || (js && wasm) || wasip1

package testenv

import (
	"errors"
	"io/fs"
	"os"
)

// Sigquit is the signal to send to kill a hanging subprocess.
// On Unix we send SIGQUIT, but on non-Unix we only have os.Kill.
var Sigquit = os.Kill

func syscallIsNotSupported(err error) bool {
	return errors.Is(err, fs.ErrPermission) || errors.Is(err, errors.ErrUnsupported)
}
