package main

// size_t StrLen(_GoString_ s) {
// 	return _GoStringLen(s);
// }
import "C"

func main() {
	C.StrLen1() // ERROR HERE
}
