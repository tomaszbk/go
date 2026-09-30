//go:build loong64 && linux

package cpu

func osInit() {
	hwcapInit()
}
