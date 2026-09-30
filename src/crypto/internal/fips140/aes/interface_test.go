package aes_test

import (
	"crypto/cipher"
	"crypto/internal/fips140/aes"
)

var _ cipher.Block = (*aes.Block)(nil)
var _ cipher.Stream = (*aes.CTR)(nil)
var _ cipher.BlockMode = (*aes.CBCDecrypter)(nil)
var _ cipher.BlockMode = (*aes.CBCEncrypter)(nil)
