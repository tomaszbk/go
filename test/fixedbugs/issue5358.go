// errorcheck


// issue 5358: incorrect error message when using f(g()) form on ... args.

package main

func f(x int, y ...int) {}

func g() (int, []int)

func main() {
	f(g()) // ERROR "as int value in|incompatible type"
}
