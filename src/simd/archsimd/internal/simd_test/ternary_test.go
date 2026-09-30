//go:build goexperiment.simd && amd64

package simd_test

import (
	"simd/archsimd"
	"testing"
)

func TestFMA(t *testing.T) {
	if archsimd.X86.AVX512() {
		testFloat32x4TernaryFlaky(t, archsimd.Float32x4.MulAdd, fmaSlice[float32], 0.001)
		testFloat32x8TernaryFlaky(t, archsimd.Float32x8.MulAdd, fmaSlice[float32], 0.001)
		testFloat32x16TernaryFlaky(t, archsimd.Float32x16.MulAdd, fmaSlice[float32], 0.001)
		testFloat64x2Ternary(t, archsimd.Float64x2.MulAdd, fmaSlice[float64])
		testFloat64x4Ternary(t, archsimd.Float64x4.MulAdd, fmaSlice[float64])
		testFloat64x8Ternary(t, archsimd.Float64x8.MulAdd, fmaSlice[float64])
	}
}
