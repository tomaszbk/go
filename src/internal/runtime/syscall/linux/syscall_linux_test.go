package linux_test

import (
	"internal/runtime/syscall/linux"
	"testing"
)

func TestEpollctlErrorSign(t *testing.T) {
	v := linux.EpollCtl(-1, 1, -1, &linux.EpollEvent{})

	const EBADF = 0x09
	if v != EBADF {
		t.Errorf("epollctl = %v, want %v", v, EBADF)
	}
}
