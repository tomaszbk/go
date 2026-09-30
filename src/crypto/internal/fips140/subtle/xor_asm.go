//go:build (amd64 || arm64 || ppc64 || ppc64le) && !purego

package subtle

//go:noescape
func xorBytes(dst, a, b *byte, n int)
