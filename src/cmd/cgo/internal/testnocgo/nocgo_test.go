package nocgo

import "testing"

func TestNop(t *testing.T) {
	i := NoCgo()
	if i != 42 {
		t.Errorf("got %d, want %d", i, 42)
	}
}
