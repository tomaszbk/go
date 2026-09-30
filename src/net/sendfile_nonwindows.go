//go:build linux || (darwin && !ios) || dragonfly || freebsd || solaris

package net

// Always true except for workstation and client versions of Windows
func supportsSendfile() bool {
	return true
}
