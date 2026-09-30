//go:build !(fips140v1.0 || fips140v1.26)

package fips140only

import (
	"crypto/internal/fips140/drbg"
	"io"
)

func ApprovedRandomReader(r io.Reader) bool {
	return drbg.IsDefaultReader(r)
}
