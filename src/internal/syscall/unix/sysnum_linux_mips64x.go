//go:build mips64 || mips64le

package unix

const (
	getrandomTrap       uintptr = 5313
	copyFileRangeTrap   uintptr = 5320
	pidfdSendSignalTrap uintptr = 5424
	pidfdOpenTrap       uintptr = 5434
	openat2Trap         uintptr = 5437
)
