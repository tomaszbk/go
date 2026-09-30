//go:build !boringcrypto

package rsa

import "crypto/internal/boring"

func boringPublicKey(*PublicKey) (*boring.PublicKeyRSA, error) {
	panic("boringcrypto: not available")
}
func boringPrivateKey(*PrivateKey) (*boring.PrivateKeyRSA, error) {
	panic("boringcrypto: not available")
}
