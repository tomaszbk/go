// compile

package main

func main() {
L1:
L2:
	for i := 0; i < 10; i++ {
		print(i)
		break L2
	}

L3:
	;
L4:
	for i := 0; i < 10; i++ {
		print(i)
		break L4
	}
	goto L1
	goto L3
}

/*
bug137.go:9: break label is not defined: L2
bug137.go:15: break label is not defined: L4
*/
