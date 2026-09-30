//go:build (unix && !linux) || (js && wasm) || wasip1 || windows

package os

import "syscall"

func ensurePidfd(sysAttr *syscall.SysProcAttr) (*syscall.SysProcAttr, bool) {
	return sysAttr, false
}

func getPidfd(_ *syscall.SysProcAttr, _ bool) (uintptr, bool) {
	return 0, false
}

func pidfdFind(_ int) (uintptr, error) {
	return 0, syscall.ENOSYS
}

func (_ *Process) pidfdWait() (*ProcessState, error) {
	panic("unreachable")
}

func (_ *Process) pidfdSendSignal(_ syscall.Signal) error {
	panic("unreachable")
}
