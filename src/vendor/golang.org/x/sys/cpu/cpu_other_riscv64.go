//go:build !linux && riscv64

package cpu

func archInit() {
	Initialized = true
}
