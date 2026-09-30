//go:build !linux

package net

import "io"

func spliceFrom(_ *netFD, _ io.Reader) (int64, error, bool) {
	return 0, nil, false
}

func spliceTo(_ io.Writer, _ *netFD) (int64, error, bool) {
	return 0, nil, false
}
