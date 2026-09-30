// errorcheck


package main

func main() {
L:
	for {
		for {
			break L2    // ERROR "L2"
			continue L2 // ERROR "L2"
		}
	}

L1:
	x := 1
	_ = x
	for {
		break L1    // ERROR "L1"
		continue L1 // ERROR "L1"
	}

	goto L
}
