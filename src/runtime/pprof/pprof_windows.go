package pprof

import (
	"fmt"
	"internal/syscall/windows"
	"io"
	"syscall"
	"unsafe"
)

func addMaxRSS(w io.Writer) {
	var m windows.PROCESS_MEMORY_COUNTERS
	p, _ := syscall.GetCurrentProcess()
	err := windows.GetProcessMemoryInfo(p, &m, uint32(unsafe.Sizeof(m)))
	if err == nil {
		fmt.Fprintf(w, "# MaxRSS = %d\n", m.PeakWorkingSetSize)
	}
}
