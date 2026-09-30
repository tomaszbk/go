//go:build !(fips140v1.0 || fips140v1.26)

package rand

import (
	"crypto/internal/fips140/drbg"
	"io"
)

// IsDefaultReader reports whether r is the default [crypto/rand.Reader].
//
// If true, the Read method of r can be assumed to call [drbg.Read].
func IsDefaultReader(r io.Reader) bool {
	return drbg.IsDefaultReader(r)
}
