//go:build amd64 || arm64 || loong64 || riscv64 || s390x

package math

const haveArchExp = true

func archExp(x float64) float64
