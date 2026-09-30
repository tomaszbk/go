// run


package main

import "fmt"

var x = uint32(0x01020304)
var y = [...]uint32{1,2,3,4,5}

func main() {
	fmt.Sprint(y[byte(x)])
}
