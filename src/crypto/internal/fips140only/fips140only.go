package fips140only

import (
	"crypto/fips140"
	"crypto/internal/fips140/sha256"
	"crypto/internal/fips140/sha3"
	"crypto/internal/fips140/sha512"
	"hash"
)

// Enforced reports whether FIPS 140-only mode is enabled and enforced, in which non-approved
// cryptography returns an error or panics.
func Enforced() bool {
	return fips140.Enforced()
}

func ApprovedHash(h hash.Hash) bool {
	switch h.(type) {
	case *sha256.Digest, *sha512.Digest, *sha3.Digest:
		return true
	default:
		return false
	}
}
