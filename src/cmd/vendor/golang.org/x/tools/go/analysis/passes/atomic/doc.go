// Package atomic defines an Analyzer that checks for common mistakes
// using the sync/atomic package.
//
// # Analyzer atomic
//
// atomic: check for common mistakes using the sync/atomic package
//
// The atomic checker looks for assignment statements of the form:
//
//	x = atomic.AddUint64(&x, 1)
//
// which are not atomic.
package atomic
