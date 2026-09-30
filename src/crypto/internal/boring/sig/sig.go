// Package sig holds “code signatures” that can be called
// and will result in certain code sequences being linked into
// the final binary. The functions themselves are no-ops.
package sig

// BoringCrypto indicates that the BoringCrypto module is present.
func BoringCrypto()

// FIPSOnly indicates that package crypto/tls/fipsonly is present.
func FIPSOnly()

// StandardCrypto indicates that standard Go crypto is present.
func StandardCrypto()
