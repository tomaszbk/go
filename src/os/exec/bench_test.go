package exec

import (
	"testing"
)

func BenchmarkExecHostname(b *testing.B) {
	b.ReportAllocs()
	path, err := LookPath("hostname")
	if err != nil {
		b.Fatalf("could not find hostname: %v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := Command(path).Run(); err != nil {
			b.Fatalf("hostname: %v", err)
		}
	}
}
