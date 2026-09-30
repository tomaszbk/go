package os

import (
	"syscall"
)

func (ph *processHandle) closeHandle() {
	syscall.Close(int(ph.handle))
}
