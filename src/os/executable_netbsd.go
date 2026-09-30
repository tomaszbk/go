package os

// From NetBSD's <sys/sysctl.h>
const (
	_CTL_KERN           = 1
	_KERN_PROC_ARGS     = 48
	_KERN_PROC_PATHNAME = 5
)

var executableMIB = [4]int32{_CTL_KERN, _KERN_PROC_ARGS, -1, _KERN_PROC_PATHNAME}
