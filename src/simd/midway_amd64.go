//go:build goexperiment.simd && amd64

package simd

import (
	"internal/cpu"
	"simd/archsimd"
)

const archHasHwClmul = true

func archMaxVectorSize() (size, allFeatureSize int, arch string) {
	arch = "amd64"
	if archsimd.X86.AVX() {
		size = 128
		allFeatureSize = 128
	}
	if archsimd.X86.AVX2() {
		size = 256
		if cpu.X86.HasVPCLMULQDQ {
			allFeatureSize = 256
		}
	}
	if archsimd.X86.AVX512() {
		size = 512
		if cpu.X86.HasAVX512VPCLMULQDQ {
			allFeatureSize = 512
		}
	}
	return
}
