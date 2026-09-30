package unix_test

import (
	"internal/syscall/unix"
	"syscall"
	"testing"
)

func TestSupportCopyFileRange(t *testing.T) {
	major, minor := unix.KernelVersion()
	t.Logf("Running on FreeBSD %d.%d\n", major, minor)

	_, err := unix.CopyFileRange(0, nil, 0, nil, 0, 0)
	want := err != syscall.ENOSYS
	got := unix.SupportCopyFileRange()
	if want != got {
		t.Fatalf("SupportCopyFileRange, got %t; want %t", got, want)
	}
}
