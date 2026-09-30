package sysrand

import "internal/syscall/windows"

func read(b []byte) error {
	return windows.ProcessPrng(b)
}
