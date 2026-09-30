// compile


// Gccgo mishandles composite literals of map with type bool.

package p

var M = map[bool]uint8{
	false: 0,
	true: 1,
}
