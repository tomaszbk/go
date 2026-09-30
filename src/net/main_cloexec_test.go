//go:build dragonfly || freebsd || linux || netbsd || openbsd || solaris

package net

import "internal/poll"

func init() {
	extraTestHookInstallers = append(extraTestHookInstallers, installAccept4TestHook)
	extraTestHookUninstallers = append(extraTestHookUninstallers, uninstallAccept4TestHook)
}

var (
	// Placeholders for saving original socket system calls.
	origAccept4 = poll.Accept4Func
)

func installAccept4TestHook() {
	poll.Accept4Func = sw.Accept4
}

func uninstallAccept4TestHook() {
	poll.Accept4Func = origAccept4
}
