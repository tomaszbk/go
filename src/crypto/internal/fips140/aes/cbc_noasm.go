//go:build (!s390x && !ppc64 && !ppc64le) || purego

package aes

func cryptBlocksEnc(b *Block, civ *[BlockSize]byte, dst, src []byte) {
	cryptBlocksEncGeneric(b, civ, dst, src)
}

func cryptBlocksDec(b *Block, civ *[BlockSize]byte, dst, src []byte) {
	cryptBlocksDecGeneric(b, civ, dst, src)
}
