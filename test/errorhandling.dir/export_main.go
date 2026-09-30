package main

import (
	"errorhandling/lib"
	"fmt"
)

func main() {
	for _, fail := range []bool{false, true} {
		n, err := lib.Pass(fail)
		if fail && (n != 0 || err != lib.Failure) || !fail && (n != 8 || err != nil) {
			panic("exported propagation")
		}
		fmt.Println(n, err)
		n, err = lib.Local(fail)
		if fail && (n != -1 || err != lib.Failure) || !fail && (n != 7 || err != nil) {
			panic("exported handler")
		}
		fmt.Println(n, err)
		s, err := lib.Generic("value", fail)
		if fail && (s != "" || err != lib.Failure) || !fail && (s != "value" || err != nil) {
			panic("exported generic")
		}
		fmt.Println(s, err)
	}
}
