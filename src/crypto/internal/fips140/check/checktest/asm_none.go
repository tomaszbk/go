//go:build (!386 && !amd64 && !arm && !arm64 && !loong64) || purego

package checktest

import "unsafe"

func PtrStaticData() *uint32        { return nil }
func PtrStaticText() unsafe.Pointer { return nil }
