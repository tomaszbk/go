// run


// Test that closures inside a generic function are not exported,
// even though not themselves generic.

package main

func Do[T any]() {
	_ = func() string {
		return ""
	}
}

func main() {
	Do[int]()
}
