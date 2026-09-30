//go:build goexperiment.runtimesecret && !arm64 && !amd64 && !loong64

package secret

import "unsafe"

func loadRegisters(p unsafe.Pointer)          {}
func spillRegisters(p unsafe.Pointer) uintptr { return 0 }
func useSecret(secret []byte)                 {}
