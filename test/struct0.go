// run


// Test zero length structs.
// Used to not be evaluated.
// Issue 2232.

package main

func recv(c chan interface{}) struct{} {
	return (<-c).(struct{})
}

var m = make(map[interface{}]int)

func recv1(c chan interface{}) {
	defer rec()
	m[(<-c).(struct{})] = 0
}

func rec() {
	recover()
}

func main() {
	c := make(chan interface{})
	go recv(c)
	c <- struct{}{}
	go recv1(c)
	c <- struct{}{}
}
