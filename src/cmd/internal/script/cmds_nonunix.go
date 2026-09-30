//go:build !unix

package script

func isETXTBSY(err error) bool {
	// syscall.ETXTBSY is only meaningful on Unix platforms.
	return false
}
