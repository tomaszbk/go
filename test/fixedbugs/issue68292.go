// errorcheck


package main

func f[S any, T any](T) {}
func g() {
	f(0) // ERROR "in call to f, cannot infer S \(declared at issue68292.go:6:8\)"
}
