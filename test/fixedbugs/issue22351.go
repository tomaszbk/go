// errorcheck

package main

import "unsafe"

const _ = uint64(unsafe.Offsetof(T{}.F)) // ERROR "undefined"
