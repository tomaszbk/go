// errorcheck


package p

var _ chan [0x2FFFF]byte         // ERROR "channel element type too large"
var _ = make(chan [0x2FFFF]byte) // ERROR "channel element type too large"

var c1 chan [0x2FFFF]byte         // ERROR "channel element type too large"
var c2 = make(chan [0x2FFFF]byte) // ERROR "channel element type too large"
