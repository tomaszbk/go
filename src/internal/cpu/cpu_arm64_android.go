//go:build arm64

package cpu

func osInit() {
	hwcapInit("android")
}
