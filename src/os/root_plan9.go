//go:build plan9

package os

import (
	"internal/filepathlite"
)

func checkPathEscapes(r *Root, name string) error {
	if r.root.closed.Load() {
		return ErrClosed
	}
	if !filepathlite.IsLocal(name) {
		return errPathEscapes
	}
	return nil
}

func checkPathEscapesLstat(r *Root, name string) error {
	return checkPathEscapes(r, name)
}
