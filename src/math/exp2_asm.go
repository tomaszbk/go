//go:build arm64 || loong64 || riscv64

package math

const haveArchExp2 = true

func archExp2(x float64) float64
