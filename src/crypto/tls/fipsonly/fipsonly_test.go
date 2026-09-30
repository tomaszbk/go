//go:build boringcrypto

package fipsonly

import (
	"crypto/tls/internal/fips140tls"
	"testing"
)

func Test(t *testing.T) {
	if !fips140tls.Required() {
		t.Fatal("fips140tls.Required() = false, must be true")
	}
}
