//go:build goexperiment.simd

package pkg

// For testing purposes, F and V are exported simd types,
// and should have the proper (variable) unsafe.Sizeof

import (
	"simd"
)

var V simd.Float32s

func F() simd.Float32s {
	var x simd.Float32s
	return x
}
