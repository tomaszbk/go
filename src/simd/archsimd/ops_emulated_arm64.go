//go:build goexperiment.simd && arm64

package archsimd

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: NEON
func (x Float32x4) ReduceSum() float32 {
	x = x.ConcatAddPairs(x) // [x0+x1, x2+x3, x0+x1, x2+x3]
	x = x.ConcatAddPairs(x) // [(x0+x1)+(x2+x3), ...]
	return x.GetElem(0)
}

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: NEON
func (x Float64x2) ReduceSum() float64 {
	return x.ConcatAddPairs(x).GetElem(0) // [x0+x1, x0+x1]
}
