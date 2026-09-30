package windows

type WSAData struct {
	Version      uint16
	HighVersion  uint16
	MaxSockets   uint16
	MaxUdpDg     uint16
	VendorInfo   *byte
	Description  [WSADESCRIPTION_LEN + 1]byte
	SystemStatus [WSASYS_STATUS_LEN + 1]byte
}

type Servent struct {
	Name    *byte
	Aliases **byte
	Proto   *byte
	Port    uint16
}

type JOBOBJECT_BASIC_LIMIT_INFORMATION struct {
	PerProcessUserTimeLimit int64
	PerJobUserTimeLimit     int64
	LimitFlags              uint32
	MinimumWorkingSetSize   uintptr
	MaximumWorkingSetSize   uintptr
	ActiveProcessLimit      uint32
	Affinity                uintptr
	PriorityClass           uint32
	SchedulingClass         uint32
}
