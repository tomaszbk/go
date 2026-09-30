// errorcheck


package main

func main() {
	n.foo = 6 // ERROR "undefined: n in n.foo|undefined name .*n|undefined: n"
}
