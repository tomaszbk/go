package gcm_test

import (
	"crypto/cipher"
	"crypto/internal/fips140/aes/gcm"
)

var _ cipher.AEAD = (*gcm.GCM)(nil)
