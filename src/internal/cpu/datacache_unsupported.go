//go:build !386 && !amd64

package cpu

func DataCacheSizes() []uintptr {
	return nil
}
