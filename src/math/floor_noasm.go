//go:build !386 && !amd64 && !arm64 && !loong64 && !ppc64 && !ppc64le && !riscv64 && !s390x && !wasm

package math

const haveArchFloor = false

func archFloor(x float64) float64 {
	panic("not implemented")
}

const haveArchCeil = false

func archCeil(x float64) float64 {
	panic("not implemented")
}

const haveArchTrunc = false

func archTrunc(x float64) float64 {
	panic("not implemented")
}
