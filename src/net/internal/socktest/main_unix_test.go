//go:build !js && !plan9 && !wasip1 && !windows

package socktest_test

import "syscall"

var (
	socketFunc func(int, int, int) (int, error)
	closeFunc  func(int) error
)

func installTestHooks() {
	socketFunc = sw.Socket
	closeFunc = sw.Close
}

func uninstallTestHooks() {
	socketFunc = syscall.Socket
	closeFunc = syscall.Close
}
