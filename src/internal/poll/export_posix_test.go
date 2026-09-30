//go:build unix || windows

// Export guts for testing on posix.
// Since testing imports os and os imports internal/poll,
// the internal/poll tests can not be in package poll.

package poll

func (fd *FD) EOFError(n int, err error) error {
	return fd.eofError(n, err)
}
