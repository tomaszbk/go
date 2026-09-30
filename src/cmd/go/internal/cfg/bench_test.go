package cfg

import (
	"cmd/internal/pathcache"
	"internal/testenv"
	"testing"
)

func BenchmarkLookPath(b *testing.B) {
	testenv.MustHaveExecPath(b, "go")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := pathcache.LookPath("go")
		if err != nil {
			b.Fatal(err)
		}
	}
}
