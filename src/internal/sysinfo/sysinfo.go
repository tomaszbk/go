// Package sysinfo implements high level hardware information gathering
// that can be used for debugging or information purposes.
package sysinfo

import (
	"internal/cpu"
	"sync"
)

var CPUName = sync.OnceValue(func() string {
	if name := cpu.Name(); name != "" {
		return name
	}

	if name := osCPUInfoName(); name != "" {
		return name
	}

	return ""
})
