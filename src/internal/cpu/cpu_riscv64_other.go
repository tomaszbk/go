//go:build riscv64 && !linux

package cpu

func osInit() {
	// Other operating systems do not support the riscv_hwprobe syscall.
}
