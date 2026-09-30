// Export guts for testing.

package runtime

import (
	"internal/runtime/syscall/windows"
	"unsafe"
)

var (
	OsYield                 = osyield
	TimeBeginPeriodRetValue = &timeBeginPeriodRetValue
)

func NumberOfProcessors() int32 {
	var info windows.SystemInfo
	stdcall(_GetSystemInfo, uintptr(unsafe.Pointer(&info)))
	return int32(info.NumberOfProcessors)
}

func GetCallerFp() uintptr {
	return getcallerfp()
}
