//go:build compiler_bootstrap

package strconv

// index returns the index of the first instance of c in s, or -1 if missing.
func index(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}
