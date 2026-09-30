//go:build fips140v1.0

package hpke

import (
	"crypto/aes"
	"crypto/cipher"
)

func newAESGCM(key []byte) (cipher.AEAD, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(b)
}
