package net

import "internal/syscall/windows"

// Workstation and client versions of Windows limit the number
// of concurrent TransmitFile operations allowed on the system
// to a maximum of two. Please see:
// https://learn.microsoft.com/en-us/windows/win32/api/mswsock/nf-mswsock-transmitfile
// https://golang.org/issue/73746
func supportsSendfile() bool {
	if testHookSupportsSendfile != nil {
		return testHookSupportsSendfile()
	}
	return windows.SupportUnlimitedTransmitFile()
}
