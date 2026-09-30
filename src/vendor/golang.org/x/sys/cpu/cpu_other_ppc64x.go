//go:build !aix && !linux && (ppc64 || ppc64le)

package cpu

func archInit() {
	PPC64.IsPOWER8 = true
	Initialized = true
}
