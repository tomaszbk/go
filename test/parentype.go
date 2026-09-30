// compile


// Test that types can be parenthesized.

package main

func f(interface{})
func g() {}
func main() {
	f(map[string]string{"a":"b","c":"d"})
	f([...]int{1,2,3})
	f(map[string]func(){"a":g,"c":g})
	f(make(chan(<-chan int)))
	f(make(chan<-(chan int)))
}
