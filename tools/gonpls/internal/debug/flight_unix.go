//go:build unix

package debug

import (
	"os"
	"syscall"
)

func init() {
	// UNIX: kill the whole process group, since
	// "go tool trace" starts a cmd/trace child.
	kill = killGroup
	sysProcAttr.Setpgid = true
}

func killGroup(p *os.Process) error {
	return syscall.Kill(-p.Pid, syscall.SIGKILL)
}
