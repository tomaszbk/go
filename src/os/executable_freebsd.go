package os

// From FreeBSD's <sys/sysctl.h>
const (
	_CTL_KERN           = 1
	_KERN_PROC          = 14
	_KERN_PROC_PATHNAME = 12
)

var executableMIB = [4]int32{_CTL_KERN, _KERN_PROC, _KERN_PROC_PATHNAME, -1}
