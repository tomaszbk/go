package net

import (
	"internal/syscall/unix"
	"testing"
)

func TestMaxAckBacklog(t *testing.T) {
	n := 196602
	backlog := maxAckBacklog(n)
	expected := 1<<16 - 1
	if unix.KernelVersionGE(4, 1) {
		expected = n
	}
	if backlog != expected {
		t.Fatalf(`sk_max_ack_backlog mismatch, got %d, want %d`, backlog, expected)
	}
}
