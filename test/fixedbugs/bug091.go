// errorcheck

package main

func f1() {
exit:
	print("hi\n")
	goto exit
}

func f2() {
	const c = 1234
}

func f3() {
	i := c // ERROR "undef"
	_ = i
}

func main() {
	f3()
}
