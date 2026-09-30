package runtime

import _ "unsafe"

func initSecureMode() {
	// We have already initialized the secureMode bool in sysauxv.
}

func isSecureMode() bool {
	return secureMode
}
