// compile


// these used to fail because the runtime
// functions that get called to implement them
// expected string, not T.

package main

type T string
func main() {
	var t T = "hello";
	println(t[0:4], t[4]);
	for _, _ = range t {
	}
	for _ = range t {
	}
	for range t {
	}
}
