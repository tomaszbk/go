// errorcheck

package main

func main() {
	for ; ; x := 1 { // ERROR "cannot declare in post statement"
		_ = x
		break
	}
}
