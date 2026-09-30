//go:build !amd64 || !gc || purego

package chacha20poly1305

func (c *chacha20poly1305) seal(dst, nonce, plaintext, additionalData []byte) []byte {
	return c.sealGeneric(dst, nonce, plaintext, additionalData)
}

func (c *chacha20poly1305) open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
	return c.openGeneric(dst, nonce, ciphertext, additionalData)
}
