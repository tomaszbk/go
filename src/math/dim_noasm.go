//go:build !amd64 && !arm64 && !loong64 && !riscv64 && !s390x

package math

const haveArchMax = false

func archMax(x, y float64) float64 {
	panic("not implemented")
}

const haveArchMin = false

func archMin(x, y float64) float64 {
	panic("not implemented")
}
