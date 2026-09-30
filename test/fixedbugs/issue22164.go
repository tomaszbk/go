// errorcheck


// Test error recovery after missing closing parentheses in lists.

package p

func f() {
	x := f(g() // ERROR "unexpected newline"
	y := 1
}

func g() {
}

func h() {
	x := f(g() // ERROR "unexpected newline"
}

func i() {
	x := []int{1, 2, 3 // ERROR "unexpected newline"
	y := 0
}