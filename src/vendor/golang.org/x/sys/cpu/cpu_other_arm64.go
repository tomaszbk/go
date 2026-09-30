//go:build !darwin && !linux && !netbsd && !openbsd && !windows && arm64

package cpu

func doinit() {
	setMinimalFeatures()
}
