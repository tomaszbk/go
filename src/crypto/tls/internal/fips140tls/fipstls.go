// Package fips140tls controls whether crypto/tls requires FIPS-approved settings.
package fips140tls

import (
	"crypto/fips140"
	"sync/atomic"
)

var required atomic.Bool

func init() {
	if fips140.Enabled() {
		Force()
	}
}

// Force forces crypto/tls to restrict TLS configurations to FIPS-approved settings.
// By design, this call is impossible to undo (except in tests).
func Force() {
	required.Store(true)
}

// Required reports whether FIPS-approved settings are required.
//
// Required is true if FIPS 140-3 mode is enabled with GODEBUG=fips140=on, or if
// the crypto/tls/fipsonly package is imported by a Go+BoringCrypto build.
func Required() bool {
	return required.Load()
}

func TestingOnlyAbandon() {
	required.Store(false)
}
