package debug

import (
	"time"
)

// Implemented in package runtime.
func readGCStats(*[]time.Duration)
func freeOSMemory()
func setMaxStack(int) int
func setGCPercent(int32) int32
func setPanicOnFault(bool) bool
func setMaxThreads(int) int
func setMemoryLimit(int64) int64
