// run

package main

func main() {
L:
	for i := 0; i < 1; i++ {
	L1:
		for {
			break L
		}
		panic("BUG: not reached - break")
		if false {
			goto L1
		}
	}

L2:
	for i := 0; i < 1; i++ {
	L3:
		for {
			continue L2
		}
		panic("BUG: not reached - continue")
		if false {
			goto L3
		}
	}
}
