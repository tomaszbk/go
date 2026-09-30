//go:build mips || mipsle

package unix

const (
	getrandomTrap       uintptr = 4353
	copyFileRangeTrap   uintptr = 4360
	pidfdSendSignalTrap uintptr = 4424
	pidfdOpenTrap       uintptr = 4434
	openat2Trap         uintptr = 4437
)
