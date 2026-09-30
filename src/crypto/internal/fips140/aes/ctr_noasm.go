//go:build (!amd64 && !arm64 && !loong64 && !s390x && !ppc64 && !ppc64le) || purego

package aes

func ctrBlocks1(b *Block, dst, src *[BlockSize]byte, ivlo, ivhi uint64) {
	ctrBlocks(b, dst[:], src[:], ivlo, ivhi)
}

func ctrBlocks2(b *Block, dst, src *[2 * BlockSize]byte, ivlo, ivhi uint64) {
	ctrBlocks(b, dst[:], src[:], ivlo, ivhi)
}

func ctrBlocks4(b *Block, dst, src *[4 * BlockSize]byte, ivlo, ivhi uint64) {
	ctrBlocks(b, dst[:], src[:], ivlo, ivhi)
}

func ctrBlocks8(b *Block, dst, src *[8 * BlockSize]byte, ivlo, ivhi uint64) {
	ctrBlocks(b, dst[:], src[:], ivlo, ivhi)
}
