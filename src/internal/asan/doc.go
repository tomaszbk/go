// Package asan contains helper functions for manually instrumenting
// code for the address sanitizer.
// The runtime package intentionally exports these functions only in the
// asan build; this package exports them unconditionally but without the
// "asan" build tag they are no-ops.
package asan
