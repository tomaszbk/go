// run


// Issue 21887: println(^uint(0)) fails to compile

package main

import "strconv"

func main() {
	if strconv.IntSize == 32 {
		println(^uint(0))
	} else {
		println(^uint32(0))
	}

	if strconv.IntSize == 64 {
		println(^uint(0))
	} else {
		println(^uint64(0))
	}
}
