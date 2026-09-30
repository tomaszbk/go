//go:build !386 && !amd64 && !amd64p32 && !arm64

package cpu

func archInit() {
	if err := readHWCAP(); err != nil {
		return
	}
	doinit()
	Initialized = true
}
