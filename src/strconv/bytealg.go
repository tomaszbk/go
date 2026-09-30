//go:build !compiler_bootstrap

package strconv

import "internal/bytealg"

// index returns the index of the first instance of c in s, or -1 if missing.
func index(s string, c byte) int {
	return bytealg.IndexByteString(s, c)
}
