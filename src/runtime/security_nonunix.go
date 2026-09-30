//go:build !unix

package runtime

func isSecureMode() bool {
	return false
}

func secure() {}
