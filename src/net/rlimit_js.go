//go:build js

package net

// concurrentThreadsLimit returns the number of threads we permit to
// run concurrently doing DNS lookups.
func concurrentThreadsLimit() int {
	return 500
}
