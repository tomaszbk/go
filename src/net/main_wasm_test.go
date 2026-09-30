//go:build wasip1 || js

package net

import "os/exec"

func installTestHooks() {}

func uninstallTestHooks() {}

func forceCloseSockets() {}

func addCmdInheritedHandle(cmd *exec.Cmd, fd uintptr) {}
