// asmcheck


package codegen

func issue63332(c chan int) {
	x := 0
	// amd64:-`MOVQ`
	x += 2
	c <- x
}
