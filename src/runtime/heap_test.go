package runtime_test

import (
	"testing"
	_ "unsafe"
)

//go:linkname heapObjectsCanMove runtime.heapObjectsCanMove
func heapObjectsCanMove() bool

func TestHeapObjectsCanMove(t *testing.T) {
	if heapObjectsCanMove() {
		// If this happens (or this test stops building),
		// it will break go4.org/unsafe/assume-no-moving-gc.
		t.Fatalf("heap objects can move!")
	}
}
