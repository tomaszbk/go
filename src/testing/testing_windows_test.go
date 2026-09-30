package testing_test

import (
	"testing"
	"time"
)

var sink time.Time
var sinkHPT testing.HighPrecisionTime

func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = time.Now()
	}
}

func BenchmarkHighPrecisionTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkHPT = testing.HighPrecisionTimeNow()
	}
}
