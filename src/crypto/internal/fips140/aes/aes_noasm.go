//go:build (!amd64 && !s390x && !ppc64 && !ppc64le && !arm64 && !loong64) || purego

package aes

type block struct {
	blockExpanded
}

func newBlock(c *Block, key []byte) *Block {
	newBlockExpanded(&c.blockExpanded, key)
	return c
}

func encryptBlock(c *Block, dst, src []byte) {
	encryptBlockGeneric(&c.blockExpanded, dst, src)
}

func decryptBlock(c *Block, dst, src []byte) {
	decryptBlockGeneric(&c.blockExpanded, dst, src)
}

func checkGenericIsExpected() {}
