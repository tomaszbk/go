// run

// issue 4785: used to fail to compile

package main

func t(x, y interface{}) interface{} {
	return x.(float64) > y.(float64)
}

func main() {
	v := t(1.0, 2.0)
	if v != false {
		panic("bad comparison")
	}
}
