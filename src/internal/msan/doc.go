// Package msan contains helper functions for manually instrumenting code
// for the memory sanitizer.
// This package exports the private msan routines in runtime unconditionally
// but without the "msan" build tag they are no-ops.
package msan
