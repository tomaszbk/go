package unix

const (
	getrandomTrap       uintptr = 355
	copyFileRangeTrap   uintptr = 377
	pidfdSendSignalTrap uintptr = 424
	pidfdOpenTrap       uintptr = 434
	openat2Trap         uintptr = 437
)
