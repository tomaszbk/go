// Test that a live variable doesn't bring its type
// descriptor live.

package main

type T [10]string

var t T

func main() {
	println(t[8])
}
