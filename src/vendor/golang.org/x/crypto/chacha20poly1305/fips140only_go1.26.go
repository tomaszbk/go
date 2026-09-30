//go:build go1.26

package chacha20poly1305

import "crypto/fips140"

func fips140Enforced() bool { return fips140.Enforced() }
