package runtime

import _ "unsafe"

//go:linkname fips140_setBypass crypto/fips140.setBypass
func fips140_setBypass() {
	getg().fipsOnlyBypass = true
}

//go:linkname fips140_unsetBypass crypto/fips140.unsetBypass
func fips140_unsetBypass() {
	getg().fipsOnlyBypass = false
}

//go:linkname fips140_isBypassed crypto/fips140.isBypassed
func fips140_isBypassed() bool {
	return getg().fipsOnlyBypass
}
