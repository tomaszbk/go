package unix

import "syscall"

// Reference: https://man.netbsd.org/open.2
const noFollowErrno = syscall.EFTYPE
