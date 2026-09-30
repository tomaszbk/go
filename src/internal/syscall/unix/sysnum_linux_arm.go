package unix

const (
	getrandomTrap       uintptr = 384
	copyFileRangeTrap   uintptr = 391
	pidfdSendSignalTrap uintptr = 424
	pidfdOpenTrap       uintptr = 434
	openat2Trap         uintptr = 437
)
