//go:build amd64 || arm64 || loong64 || riscv64 || s390x

package math

const haveArchMax = true

func archMax(x, y float64) float64

const haveArchMin = true

func archMin(x, y float64) float64
