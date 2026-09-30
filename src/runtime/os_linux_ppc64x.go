//go:build linux && (ppc64 || ppc64le)

package runtime

import "internal/cpu"

func archauxv(tag, val uintptr) {
	switch tag {
	case _AT_HWCAP:
		// ppc64x doesn't have a 'cpuid' instruction
		// equivalent and relies on HWCAP/HWCAP2 bits for
		// hardware capabilities.
		cpu.HWCap = uint(val)
	case _AT_HWCAP2:
		cpu.HWCap2 = uint(val)
	}
}
