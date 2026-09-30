//go:build !linux

package runtime

import "unsafe"

// setVMAName isn’t implemented
func setVMAName(start unsafe.Pointer, len uintptr, name string) {}

func setVMANameSupported() bool { return false }
