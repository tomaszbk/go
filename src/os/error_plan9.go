package os

import "syscall"

type syscallErrorType = syscall.ErrorString

var errENOSYS = syscall.NewError("function not implemented")
var errERANGE = syscall.NewError("out of range")
var errENOMEM = syscall.NewError("cannot allocate memory")
