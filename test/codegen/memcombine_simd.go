// asmcheck
//go:build goexperiment.simd

package codegen

import "simd/archsimd"

// TODO: Move these tests to memcombine.go when GOEXPERIMENT=simd becomes the default
func dwloadInt64x2(p *struct{ a, b archsimd.Int64x2 }) (archsimd.Int64x2, archsimd.Int64x2) {
	// arm64:"FLDPQ "
	return p.a, p.b
}

func dwstoreInt64x2(p *struct{ a, b archsimd.Int64x2 }, a, b archsimd.Int64x2) {
	// arm64:`FSTPQ\s\(F[0-9]+, F[0-9]+\), \(R[0-9]+\)`
	p.a = a
	p.b = b
}
