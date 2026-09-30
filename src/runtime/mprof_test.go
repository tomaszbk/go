package runtime_test

import (
	. "runtime"
	"testing"
)

func BenchmarkBlocksampled(b *testing.B) {
	for b.Loop() {
		Blocksampled(42, 1337)
	}
}
