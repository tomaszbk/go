//go:build goexperiment.simd && !amd64

package simd_test

import (
	"simd"
)

func sum(x simd.Float32s) float32 {
	return boringSum(x)
}
