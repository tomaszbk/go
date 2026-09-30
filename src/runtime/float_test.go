package runtime_test

import (
	"testing"
)

func TestIssue48807(t *testing.T) {
	for _, i := range []uint64{
		0x8234508000000001, // from issue48807
		1<<56 + 1<<32 + 1,
	} {
		got := float32(i)
		dontwant := float32(float64(i))
		if got == dontwant {
			// The test cases above should be uint64s such that
			// this equality doesn't hold. These examples trigger
			// the case where using an intermediate float64 doesn't work.
			t.Errorf("direct float32 conversion doesn't work: arg=%x got=%x dontwant=%x", i, got, dontwant)
		}
	}
}
