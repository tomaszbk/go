package runtime

func checkgoarm() {
	// TODO(minux): FP checks like in os_linux_arm.go.

	// osinit not called yet, so numCPUStartup not set: must use
	// getCPUCount directly.
	if getCPUCount() > 1 && goarm < 7 {
		print("runtime: this system has multiple CPUs and must use\n")
		print("atomic synchronization instructions. Recompile using GOARM=7.\n")
		exit(1)
	}
}

//go:nosplit
func cputicks() int64 {
	// runtime·nanotime() is a poor approximation of CPU ticks that is enough for the profiler.
	return nanotime()
}
