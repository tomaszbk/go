//go:build !(linux || (darwin && !ios) || dragonfly || freebsd || solaris || windows)

package net

import "io"

var testHookSupportsSendfile func() bool

func supportsSendfile() bool {
	return false
}

func sendFile(c *netFD, r io.Reader) (n int64, err error, handled bool) {
	return 0, nil, false
}
