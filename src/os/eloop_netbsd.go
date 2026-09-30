//go:build netbsd

package os

import "syscall"

// isNoFollowErr reports whether err may result from O_NOFOLLOW blocking an open operation.
func isNoFollowErr(err error) bool {
	// NetBSD returns EFTYPE, but check the other possibilities as well.
	switch err {
	case syscall.ELOOP, syscall.EMLINK, syscall.EFTYPE:
		return true
	}
	return false
}
