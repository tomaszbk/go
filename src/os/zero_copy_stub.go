//go:build !freebsd && !linux && !solaris

package os

import "io"

func (f *File) writeTo(w io.Writer) (written int64, handled bool, err error) {
	return 0, false, nil
}

func (f *File) readFrom(r io.Reader) (n int64, handled bool, err error) {
	return 0, false, nil
}
