//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package telemetry

import (
	"os/exec"
	"syscall"
)

func init() {
	daemonize = daemonizePosix
}

func daemonizePosix(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
}
