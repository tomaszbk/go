// errorcheck

package main

func f(a int, b ...int) {}

func main() {
	f(nil...) // ERROR "not enough arguments in call to f\n\thave \(nil\)\n\twant \(int, \[\]int\)|not enough arguments"
}
