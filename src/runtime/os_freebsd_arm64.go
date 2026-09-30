package runtime

//go:nosplit
func cputicks() int64 {
	// nanotime() is a poor approximation of CPU ticks that is enough for the profiler.
	return nanotime()
}
