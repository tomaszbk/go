// run


// issue 5809: 6g and 8g attempted to constant propagate indexed LEA

package main

import "fmt"

func main() {
	const d16 = "0123456789ABCDEF"
	k := 0x1234
	var x [4]byte
	
	x[0] = d16[k>>12&0xf]
	x[1] = d16[k>>8&0xf]
	x[2] = d16[k>>4&0xf]
	x[3] = d16[k&0xf]
	
	if x != [4]byte{'1','2','3','4'} {
		fmt.Println(x)
		panic("x != [4]byte{'1','2','3','4'}")
	}
}
