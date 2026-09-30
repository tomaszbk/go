package net

import "internal/syscall/windows"

func supportsUnixSocket() bool {
	return windows.SupportUnixSocket()
}
