//go:build cgo

package cgobench_test

import (
	"internal/runtime/cgobench"
	"testing"
)

func BenchmarkCall(b *testing.B) {
	for b.Loop() {
		cgobench.Empty()
	}
}

func BenchmarkCallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cgobench.Empty()
		}
	})
}

func BenchmarkCgoCall(b *testing.B) {
	for b.Loop() {
		cgobench.EmptyC()
	}
}

func BenchmarkCgoCallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cgobench.EmptyC()
		}
	})
}

func BenchmarkCgoCallWithCallback(b *testing.B) {
	for b.Loop() {
		cgobench.CallbackC()
	}
}

func BenchmarkCgoCallParallelWithCallback(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cgobench.CallbackC()
		}
	})
}
