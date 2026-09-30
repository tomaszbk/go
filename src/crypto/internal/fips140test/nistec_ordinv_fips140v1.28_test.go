//go:build !(fips140v1.0 || fips140v1.26)

package fipstest

import (
	"crypto/internal/fips140/nistec"
	"testing"
)

func p256OrdInverse(t *testing.T, k *[4]uint64) {
	nistec.P256OrdInverse(k)
}
