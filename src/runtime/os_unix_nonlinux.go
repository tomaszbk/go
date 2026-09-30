//go:build unix && !linux

package runtime

// sigFromUser reports whether the signal was sent because of a call
// to kill.
//
//go:nosplit
func (c *sigctxt) sigFromUser() bool {
	return c.sigcode() == _SI_USER
}

// sigFromSeccomp reports whether the signal was sent from seccomp.
//
//go:nosplit
func (c *sigctxt) sigFromSeccomp() bool {
	return false
}
