//go:build unix

package runtime

import (
	"internal/stringslite"
)

func secure() {
	initSecureMode()

	if !isSecureMode() {
		return
	}

	// When secure mode is enabled, we do one thing: enforce specific
	// environment variable values (currently we only force GOTRACEBACK=none)
	//
	// Other packages may also disable specific functionality when secure mode
	// is enabled (determined by using linkname to call isSecureMode).

	secureEnv()
}

func secureEnv() {
	var hasTraceback bool
	for i := 0; i < len(envs); i++ {
		if stringslite.HasPrefix(envs[i], "GOTRACEBACK=") {
			hasTraceback = true
			envs[i] = "GOTRACEBACK=none"
		}
	}
	if !hasTraceback {
		envs = append(envs, "GOTRACEBACK=none")
	}
}
