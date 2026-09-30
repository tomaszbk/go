//go:build sparc64

package cpu

// The L1 line is 32 bytes; false sharing is governed by the 64-byte L2 line.
const cacheLineSize = 64

func initOptions() {}
