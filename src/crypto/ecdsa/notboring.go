//go:build !boringcrypto

package ecdsa

import "crypto/internal/boring"

func boringPublicKey(*PublicKey) (*boring.PublicKeyECDSA, error) {
	panic("boringcrypto: not available")
}
func boringPrivateKey(*PrivateKey) (*boring.PrivateKeyECDSA, error) {
	panic("boringcrypto: not available")
}
