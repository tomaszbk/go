//go:build !linux || !go1.24

package unix

func vgetrandom(p []byte, flags uint32) (ret int, supported bool) {
	return -1, false
}
