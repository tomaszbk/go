//go:build goexperiment.simd && wasm

package simd

const archHasHwClmul = false

func archMaxVectorSize() (size, allFeatureSize int, arch string) {
	return 128, 128, "wasm"
}
