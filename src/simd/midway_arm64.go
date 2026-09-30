//go:build goexperiment.simd && arm64

package simd

import (
	"internal/cpu"
)

const archHasHwClmul = true

func archMaxVectorSize() (size, allFeatureSize int, arch string) {
	arch = "arm64"
	// This describes Neon, SVE is still TBD.
	size = 128
	if cpu.ARM64.HasPMULL {
		allFeatureSize = 128
	}
	return
}
