package c

import "./a"
import "./b"

var _ a.A = b.B() // ERROR "cannot use b\.B|incompatible type"
