// errorcheck

// Issue 6403: fix spurious 'const initializer is not a constant' error

package p

import "syscall"

const A int = syscall.X // ERROR "undefined: syscall.X|undefined identifier .*syscall.X"
const B int = voidpkg.X // ERROR "undefined: voidpkg|undefined name .*voidpkg"
