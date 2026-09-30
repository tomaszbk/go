// errorcheck


// Check error message for duplicated index in slice literal

package p

var s = []string{
	1: "dup",
	1: "dup", // ERROR "duplicate index in slice literal: 1|duplicate value for index 1|duplicate index 1"
}
