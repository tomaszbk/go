// errorcheck


package main

func f() {
	g(f..3) // ERROR "unexpected literal \.3, expected name or \("
}
