// errorcheck

package main

var x int

var a = []int{x: 1}    // ERROR "constant"
var b = [...]int{x: 1} // ERROR "constant"
var c = map[int]int{x: 1}
