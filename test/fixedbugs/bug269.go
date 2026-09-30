// run


// https://golang.org/issue/749

package main

func f() (ok bool) { return false }

func main() {
	var i interface{}
	i = f
	_ = i.(func()bool)
	_ = i.(func()(bool))
}
