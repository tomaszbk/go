//go:build arm64 && linux && !android

package cpu

func osInit() {
	hwcapInit("linux")
}
