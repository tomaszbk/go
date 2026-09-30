package sys

import (
	"testing"
)

func TestArchInFamily(t *testing.T) {
	if got, want := ArchPPC64LE.InFamily(AMD64), false; got != want {
		t.Errorf("Got ArchPPC64LE.InFamily(AMD64) = %v, want %v", got, want)
	}
	if got, want := ArchPPC64LE.InFamily(PPC64), true; got != want {
		t.Errorf("Got ArchPPC64LE.InFamily(PPC64) = %v, want %v", got, want)
	}
	if got, want := ArchPPC64LE.InFamily(AMD64, RISCV64), false; got != want {
		t.Errorf("Got ArchPPC64LE.InFamily(AMD64, RISCV64) = %v, want %v", got, want)
	}
	if got, want := ArchPPC64LE.InFamily(AMD64, PPC64), true; got != want {
		t.Errorf("Got ArchPPC64LE.InFamily(AMD64, PPC64) = %v, want %v", got, want)
	}
}
