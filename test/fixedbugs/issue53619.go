// run


package main

var c = b
var d = a

var a, b any = any(nil).(bool)

func main() {
	if c != false {
		panic(c)
	}
	if d != false {
		panic(d)
	}
}
