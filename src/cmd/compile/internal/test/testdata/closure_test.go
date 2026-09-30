// closure.go tests closure operations.
package main

import "testing"

//go:noinline
func testCFunc_ssa() int {
	a := 0
	b := func() {
		switch {
		}
		a++
	}
	b()
	b()
	return a
}

func testCFunc(t *testing.T) {
	if want, got := 2, testCFunc_ssa(); got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

// TestClosure tests closure related behavior.
func TestClosure(t *testing.T) {
	testCFunc(t)
}
