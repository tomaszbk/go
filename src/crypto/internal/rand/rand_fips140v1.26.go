//go:build !fips140v1.0

package rand

import (
	"crypto/internal/fips140/drbg"
	"io"
)

func fips140SetTestingReader(r io.Reader) {
	drbg.SetTestingReader(r)
}
