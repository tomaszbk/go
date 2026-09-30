//go:build !unix

package base

func IsETXTBSY(err error) bool {
	// syscall.ETXTBSY is only meaningful on Unix platforms.
	return false
}
