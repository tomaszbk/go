//go:build !linux && (mips64 || mips64le)

package cpu

func archInit() {
	Initialized = true
}
