//go:build !arm64 && !loong64 && !riscv64

package math

const haveArchExp2 = false

func archExp2(x float64) float64 {
	panic("not implemented")
}
