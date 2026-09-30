//go:build goexperiment.simd

package bridge

// ZeroSized is used as the definition for type _simd in package simd, to create
// a hard dependence between the two packages before code transformation.
type ZeroSized struct {
	_ [0]func(*ZeroSized) *ZeroSized
}
