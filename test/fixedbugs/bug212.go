// errorcheck


package main
type I int
type S struct { f map[I]int }
var v1 = S{ make(map[int]int) }		// ERROR "cannot|illegal|incompatible|wrong"
var v2 map[I]int = map[int]int{}	// ERROR "cannot|illegal|incompatible|wrong"
var v3 = S{ make(map[uint]int) }	// ERROR "cannot|illegal|incompatible|wrong"
