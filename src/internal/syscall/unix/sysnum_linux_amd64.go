package unix

const (
	getrandomTrap       uintptr = 318
	copyFileRangeTrap   uintptr = 326
	pidfdSendSignalTrap uintptr = 424
	pidfdOpenTrap       uintptr = 434
	openat2Trap         uintptr = 437
)
