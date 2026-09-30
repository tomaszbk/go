package unix

const (
	getrandomTrap       uintptr = 349
	copyFileRangeTrap   uintptr = 375
	pidfdSendSignalTrap uintptr = 424
	pidfdOpenTrap       uintptr = 434
	openat2Trap         uintptr = 437
)
