//go:build !plan9 && !windows

package modfetch

import (
	"fmt"
	"syscall"
	"testing"
)

func TestIsErrReadOnlyFS(t *testing.T) {
	if isErrReadOnlyFS(nil) {
		t.Error("isErrReadOnlyFS(nil) = true, want false")
	}
	if isErrReadOnlyFS(fmt.Errorf("some error")) {
		t.Error("isErrReadOnlyFS(non-EROFS) = true, want false")
	}
	if !isErrReadOnlyFS(syscall.EROFS) {
		t.Error("isErrReadOnlyFS(syscall.EROFS) = false, want true")
	}
	if !isErrReadOnlyFS(fmt.Errorf("wrapped: %w", syscall.EROFS)) {
		t.Error("isErrReadOnlyFS(wrapped EROFS) = false, want true")
	}
	if isErrReadOnlyFS(syscall.EACCES) {
		t.Error("isErrReadOnlyFS(syscall.EACCES) = true, want false")
	}
}
