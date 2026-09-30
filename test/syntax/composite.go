// errorcheck


package main

var a = []int{
	3 // ERROR "need trailing comma before newline in composite literal|possibly missing comma or }"
}
