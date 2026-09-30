//go:build (!arm64 && !loong64 && !s390x && !ppc64 && !ppc64le) || !gc || purego

package chacha20

const bufSize = blockSize

func (s *Cipher) xorKeyStreamBlocks(dst, src []byte) {
	s.xorKeyStreamBlocksGeneric(dst, src)
}
