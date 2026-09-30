package traceviewer

const (
	// Special P identifiers:
	FakeP    = 1000000 + iota
	TimerP   // depicts timer unblocks
	NetpollP // depicts network unblocks
	SyscallP // depicts returns from syscalls
	GCP      // depicts GC state
	ProfileP // depicts recording of CPU profile samples
)
