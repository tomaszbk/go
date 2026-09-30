// errorcheck


package main

func main() {
	var i33 int64;
	if i33 == (1<<64) -1 {  // ERROR "overflow"
	}
}
