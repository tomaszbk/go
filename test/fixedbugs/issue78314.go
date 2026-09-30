// errorcheck


package main

func F() {
	for range (func(func(int, ...string) bool))(nil) { // ERROR "cannot be variadic"
	}
}

func main() {}
